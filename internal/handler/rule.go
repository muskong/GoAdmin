package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoAdmin/internal/logic"
)

func AdminRuleList(c *gin.Context) {
	var page logic.Page
	err := c.ShouldBind(&page)
	if err != nil {
		c.SecureJSON(http.StatusOK, "传入参数错误")
		return
	}

	err, data := logic.AdminRuleList(page)

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

	err = logic.AdminRuleCreate(user)

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error())
		return
	}
	c.SecureJSON(http.StatusOK, "ok")
}

func AdminRuleAll(c *gin.Context) {
	var q struct {
		RuleId int64 `json:"adminRuleId"`
	}
	err := c.ShouldBindQuery(&q)
	if err != nil {
		c.SecureJSON(http.StatusOK, "传入参数错误")
		return
	}

	err, ruleData := logic.AdminRuleAll(q.RuleId)
	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error())
		return
	}

	c.SecureJSON(http.StatusOK, ruleData)
}
