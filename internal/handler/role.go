package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoAdmin/internal/logic"
)

func AdminRoleList(c *gin.Context) {
	var page logic.Page
	err := c.ShouldBind(&page)
	if err != nil {
		c.SecureJSON(http.StatusOK, "传入参数错误")
		return
	}

	err, data := logic.AdminRoleList(page)

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error())
		return
	}

	c.SecureJSON(http.StatusOK, data)
}

func AdminRoleCreate(c *gin.Context) {
	var user entity.AdminRole
	err := c.ShouldBind(&user)
	if err != nil {
		c.SecureJSON(http.StatusOK, "传入参数错误")
		return
	}

	err = logic.AdminRoleCreate(user)

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error())
		return
	}
}

func AdminRoleRuleList(c *gin.Context) {
	var q struct {
		RoleId int `json:"roleId"`
	}
	err := c.ShouldBindQuery(&q)
	if q.RoleId <= 0 {
		c.SecureJSON(http.StatusOK, "传入参数错误")
		return
	}

	err, data := logic.AdminRoleRuleList(q.RoleId)

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error())
		return
	}

	c.SecureJSON(http.StatusOK, data)
}

func AdminRoleAll(c *gin.Context) {
	err, data := logic.AdminRoleAll()

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error())
		return
	}

	c.SecureJSON(http.StatusOK, data)
}

func AdminRoleSaveRule(c *gin.Context) {
	var roleRule logic.RoleRuleObject
	err := c.ShouldBind(&roleRule)
	if err != nil {
		c.SecureJSON(http.StatusOK, "传入参数错误")
		return
	}

	err = logic.AdminRoleSaveRule(roleRule)

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error())
		return
	}
	c.SecureJSON(http.StatusOK, "设置成功")
}
