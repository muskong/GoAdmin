package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoAdmin/internal/logic"
	"github.com/muskong/GoCore/respond"
)

// 00000000  AdminRoleGroup ---
type roleGrop struct{}

var Role = &roleGrop{}

func (*roleGrop) AdminRole(c *gin.Context) {
	var q struct {
		RoleId int `json:"roleId"`
	}
	err := c.ShouldBindQuery(&q)
	if q.RoleId <= 0 {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	data, err := logic.Role.AdminRole(q.RoleId)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}
func (*roleGrop) AdminRoleList(c *gin.Context) {
	var page logic.Page
	err := c.ShouldBind(&page)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	err, data := logic.Role.AdminRoleList(page)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}

func (*roleGrop) AdminRoleCreate(c *gin.Context) {
	var user entity.AdminRole
	err := c.ShouldBind(&user)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	err = logic.Role.AdminRoleCreate(user)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
}

func (*roleGrop) AdminRoleUpdate(c *gin.Context) {
	var user entity.AdminRole
	err := c.ShouldBind(&user)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	err = logic.Role.AdminRoleCreate(user)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
}

func (*roleGrop) AdminRoleDelete(c *gin.Context) {
	var q struct {
		RoleId int `json:"roleId"`
	}
	err := c.ShouldBindQuery(&q)
	if q.RoleId <= 0 {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	err = logic.Role.AdminRoleDelete(q.RoleId)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}

func (*roleGrop) AdminRoleRuleList(c *gin.Context) {
	var q struct {
		RoleId int `json:"roleId"`
	}
	err := c.ShouldBindQuery(&q)
	if q.RoleId <= 0 {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	err, data := logic.Role.AdminRoleRuleList(q.RoleId)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}

func (*roleGrop) AdminRoleAll(c *gin.Context) {
	err, data := logic.Role.AdminRoleAll()

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}

func (*roleGrop) AdminRoleSaveRule(c *gin.Context) {
	var roleRule logic.RoleRuleObject
	err := c.ShouldBind(&roleRule)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	err = logic.Role.AdminRoleSaveRule(roleRule)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("设置成功"))
}

func (*roleGrop) AdminRoleGroup(c *gin.Context) {

	data, err := logic.Role.AdminRoleGroupList()

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}
