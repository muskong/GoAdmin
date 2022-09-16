package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/muskong/GoAdmin/internal"
	"github.com/muskong/GoAdmin/internal/logic"

	"github.com/muskong/GoCore/config"
	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/jwt"
	"github.com/muskong/GoPkg/zaplog"
)

func main() {
	config.Init()

	zaplog.InitLogger(&zaplog.ZapConfig{
		Env:        config.Log.GetString("env"),
		AppName:    config.App.GetString("app.name"),
		AppSubName: config.App.GetString("app.sub.name"),
		Level:      config.Log.GetString("level"),
		Logfile:    config.Log.GetString("path"),
		MaxSize:    config.Log.GetInt("max_size"),
		MaxBackups: config.Log.GetInt("max_backups"),
		MaxAge:     config.Log.GetInt("max_age"),
		Compress:   config.Log.GetBool("compress"),
	})

	gorm.Init(&gorm.GormConfig{
		Driver:          config.Db.GetString("driver"),
		Dsn:             config.Db.GetString("dsn"),
		MaxOpenConns:    config.Db.GetInt("max_open"),
		MaxIdleConns:    config.Db.GetInt("max_idle"),
		MaxConnLifetime: 0,
		Debug:           config.Db.GetBool("debug_sql"),
	})
	jwt.JwtInit(&jwt.JwtConfig{
		TokenName: config.App.GetString("token.name"),
		Server:    config.App.GetInt64("uuid.server"),
		Exp:       config.App.GetInt64("token.expired"),
		Iss:       config.App.GetString("app.name"),
		Aud:       config.App.GetString("app.sub.name"),
		Pub:       config.Rsa.GetString("rsa.public"),
		Pri:       config.Rsa.GetString("rsa.private"),
	})

	logic.InitToken()
	logic.InitCard()

	// 服务器
	listenAddr := config.App.GetString("app.server")
	//err = http.ListenAndServe(listenAddr, header)
	server := &http.Server{
		Addr:         listenAddr,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler:      internal.GinRouter(),
	}

	go func() {
		zaplog.Sugar.Info("Server Start")
		err := server.ListenAndServe()
		// 服务连接
		if err != nil && err != http.ErrServerClosed {
			zaplog.Sugar.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		zaplog.Sugar.Fatal("Server Shutdown:", err)
	}
	zaplog.Sugar.Info("Server exiting")
}
