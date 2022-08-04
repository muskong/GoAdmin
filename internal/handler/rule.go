package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoAdmin/internal/logic"
	"github.com/muskong/GoCore/respond"
	"github.com/spf13/cast"
)

type ruleGrop struct{}

var Rule = &ruleGrop{}

func (*ruleGrop) AdminRule(c *gin.Context) {
	var q struct {
		RuleId int `json:"ruleId"`
	}
	err := c.ShouldBindUri(&q)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	data, err := logic.Rule.AdminRule(q.RuleId)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}

func (*ruleGrop) AdminRuleList(c *gin.Context) {
	var page logic.Page
	err := c.ShouldBind(&page)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	err, data := logic.Rule.AdminRuleList(page)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}

func (*ruleGrop) AdminRuleCreate(c *gin.Context) {
	var user entity.AdminRule
	err := c.ShouldBind(&user)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	err = logic.Rule.AdminRuleCreate(user)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}

func (*ruleGrop) AdminRuleUpdate(c *gin.Context) {
	var rule entity.AdminRule
	err := c.ShouldBind(&rule)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	err = logic.Rule.AdminRuleUpdate(rule)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}

func (*ruleGrop) AdminRuleDelete(c *gin.Context) {
	var user entity.AdminRule
	err := c.ShouldBind(&user)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	err = logic.Rule.AdminRuleCreate(user)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}

func (*ruleGrop) AdminRuleAll(c *gin.Context) {
	var q struct {
		RuleId int `json:"ruleId"`
	}
	err := c.ShouldBindQuery(&q)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	err, ruleData := logic.Rule.AdminRuleAll(q.RuleId)
	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(ruleData))
}

func (*ruleGrop) AdminMenu(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		c.SecureJSON(respond.Message("未登录"))
	}

	user, err := logic.User.AdminUser(cast.ToInt(userId))
	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	menu, err := logic.Rule.AdminRuleTree(user.Roles)
	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(map[string]any{
		"adminInfo": user,
		"menus":     menu,
		"siteConfig": map[string]string{
			"siteName": "test",
			"version":  "version",
		},
	}))
}
