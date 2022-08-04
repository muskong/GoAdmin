package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/logic"
	"github.com/muskong/GoCore/respond"
)

func AuthLogin(c *gin.Context) {
	var userForm logic.LoginData
	err := c.ShouldBind(&userForm)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	data, err := logic.LoginVerify(userForm)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}
func Sites(ctx *gin.Context) {
	ctx.SecureJSON(respond.Data(map[string]any{
		"siteName":     "站点名称",
		"recordNumber": "测试ICP备8888888号-1",
		"version":      "v1.0.0",
	}))
}
func Index(ctx *gin.Context) {
	ctx.SecureJSON(respond.Data("ok"))
}
