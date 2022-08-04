package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoCore/respond"
)

func Dashboard(c *gin.Context) {
	c.SecureJSON(respond.Data(map[string]any{
		"remark": "开源等于互助；开源需要大家一起来支持，支持的方式有很多种，比如使用、推荐、写教程、保护生态、贡献代码、回答问题、分享经验、打赏赞助等；欢迎您加入我们！",
	}))
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
