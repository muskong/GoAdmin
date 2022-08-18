package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoAdmin/internal/logic"
	"github.com/muskong/GoCore/respond"
	"github.com/muskong/GoPkg/zaplog"
)

// 00000000  AdminRoleGroup ---
type roleGrop struct{}

var Role = &roleGrop{}

func (*roleGrop) AdminRole(c *gin.Context) {
	var q struct {
		RoleId string `json:"roleId"`
	}
	err := c.ShouldBindQuery(&q)
	if q.RoleId == "" {
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

	data, err := logic.Role.AdminRoleList(page)

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
		zaplog.Sugar.Error(err)
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	logic.Log.Context(c)
	err = logic.Role.AdminRoleCreate(user)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}

func (*roleGrop) AdminRoleUpdate(c *gin.Context) {
	var user entity.AdminRole
	err := c.ShouldBind(&user)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	logic.Role.Context(c)
	err = logic.Role.AdminRoleCreate(user)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}

func (*roleGrop) AdminRoleDelete(c *gin.Context) {
	var q Id
	err := c.ShouldBindUri(&q)
	if err != nil {
		zaplog.Sugar.Error(err)
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	logic.Role.Context(c)
	err = logic.Role.AdminRoleDelete(q.Nanoid)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}

func (*roleGrop) AdminRoleRuleList(c *gin.Context) {
	var q struct {
		RoleId string `json:"roleId"`
	}
	err := c.ShouldBindQuery(&q)
	if q.RoleId == "" {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	data, err := logic.Role.AdminRoleRuleList(q.RoleId)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}

func (*roleGrop) AdminRoleAll(c *gin.Context) {
	data, err := logic.Role.AdminRoleAll()

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

	logic.Role.Context(c)
	err = logic.Role.AdminRoleSaveRule(roleRule)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("设置成功"))
}

func (*roleGrop) AdminRoleTree(c *gin.Context) {

	data, err := logic.Role.AdminRoleTree()

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}

func (*roleGrop) AdminRoleGroup(c *gin.Context) {

	data, err := logic.Role.AdminRoleGroupList()

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}
