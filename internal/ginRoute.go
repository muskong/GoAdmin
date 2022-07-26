package internal

import (
	"github.com/muskong/GoAdmin/internal/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoCore/middlewares"
	"github.com/muskong/GoPkg/jwt"
)

func GinRouter() *gin.Engine {
	tokenName := jwt.Jwt.GetTokenName()
	notAuth := map[string]bool{
		"/admin/user/login": true,
		"/admin/sites":      true,
	}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"lang", "atoken", "utoken", "Origin", "Content-Length", "Content-Type"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))
	// router.Use(cors.Default())
	router.Use(middlewares.GinUserMiddleware(tokenName, notAuth))

	router.GET("/", handler.Index)
	adm := router.Group("/admin")

	adm.GET("/sites", handler.Sites)

	adminUser := adm.Group("/user")
	{
		adminUser.POST("/login", handler.AuthLogin)
		adminUser.GET("/list", handler.AdminUserList)
		adminUser.POST("/create", handler.AdminUserCreate)
	}

	adminRole := adm.Group("/role")
	{
		adminRole.GET("/rules", handler.AdminRoleRuleList)
		adminRole.GET("/all", handler.AdminRoleAll)
		adminRole.GET("/list", handler.AdminRoleList)
		adminRole.POST("/create", handler.AdminRoleCreate)
		adminRole.POST("/saveRule", handler.AdminRoleSaveRule)
	}

	adminRule := adm.Group("/rule")
	{
		adminRule.GET("/all", handler.AdminRuleAll)
		adminRule.GET("/list", handler.AdminRuleList)
		adminRule.POST("/create", handler.AdminRuleCreate)
	}

	return router
}
