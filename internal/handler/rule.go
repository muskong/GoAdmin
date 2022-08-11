package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoAdmin/internal/logic"
	"github.com/muskong/GoCore/respond"
)

type ruleGrop struct{}

var Rule = &ruleGrop{}

func (*ruleGrop) AdminRule(c *gin.Context) {
	var q struct {
		RuleId string `json:"ruleId"`
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
		RuleId string `json:"ruleId"`
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

func (*ruleGrop) AdminRuleTree(c *gin.Context) {
	data, err := logic.Rule.AdminRuleTree()

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}
