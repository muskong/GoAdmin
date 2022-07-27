package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoAdmin/internal/logic"
	"github.com/muskong/GoCore/helper"
	"github.com/spf13/cast"
)

func AdminRuleList(c *gin.Context) {
	var page logic.Page
	err := c.ShouldBind(&page)
	if err != nil {
		c.SecureJSON(http.StatusOK, "传入参数错误")
		return
	}

	err, data := logic.Rule.AdminRuleList(page)

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error())
		return
	}

	c.SecureJSON(http.StatusOK, data)
}

func AdminRuleCreate(c *gin.Context) {
	var user entity.AdminRule
	err := c.ShouldBind(&user)
	if err != nil {
		c.SecureJSON(http.StatusOK, "传入参数错误")
		return
	}

	err = logic.Rule.AdminRuleCreate(user)

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error())
		return
	}
	c.SecureJSON(http.StatusOK, "ok")
}

func AdminRuleAll(c *gin.Context) {
	var q struct {
		RuleId int `json:"adminRuleId"`
	}
	err := c.ShouldBindQuery(&q)
	if err != nil {
		c.SecureJSON(http.StatusOK, "传入参数错误")
		return
	}

	err, ruleData := logic.Rule.AdminRuleAll(q.RuleId)
	if err != nil {
		c.SecureJSON(helper.Error(err.Error()))
		return
	}

	c.SecureJSON(helper.Success(ruleData))
}

func AdminMenu(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		c.SecureJSON(http.StatusOK, "未登录")
	}

	user, err := logic.User.GetUser(cast.ToInt64(userId))
	if err != nil {
		c.SecureJSON(helper.Error(err.Error()))
		return
	}

	menu, err := logic.Rule.AdminRuleTree(user.Roles)
	if err != nil {
		c.SecureJSON(helper.Error(err.Error()))
		return
	}

	c.SecureJSON(helper.Success(map[string]any{
		"adminInfo": user,
		"menus":     menu,
		"siteConfig": map[string]string{
			"siteName": "test",
			"version":  "version",
		},
	}))
}
