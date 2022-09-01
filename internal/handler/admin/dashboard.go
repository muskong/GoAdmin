package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/logic"
	"github.com/muskong/GoCore/respond"
	"github.com/spf13/cast"
)

func Dashboard(c *gin.Context) {
	c.SecureJSON(respond.Data(map[string]any{
		"remark": "开源等于互助；开源需要大家一起来支持，支持的方式有很多种，比如使用、推荐、写教程、保护生态、贡献代码、回答问题、分享经验、打赏赞助等；欢迎您加入我们！",
	}))
}

func Sites(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok || userId == "" {
		ctx.SecureJSON(respond.Message("未登录"))
		return
	}

	user, err := logic.AdminUser.AdminUser(cast.ToInt(userId))
	if err != nil {
		ctx.SecureJSON(respond.Message(err.Error()))
		return
	}

	menu, err := logic.Rule.AdminRuleGroup(user.Roles)
	if err != nil {
		ctx.SecureJSON(respond.Message(err.Error()))
		return
	}

	ctx.SecureJSON(respond.Data(map[string]any{
		"AdminInfo": user,
		"Menus":     menu,
		"SiteConfig": map[string]string{
			"SiteName":     "站点名称",
			"RecordNumber": "测试ICP备8888888号-1",
			"Version":      "v1.0.0",
		},
	}))
}
func Index(ctx *gin.Context) {
	ctx.SecureJSON(respond.Data("ok"))
}
