package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoAdmin/internal/logic"
	"github.com/muskong/GoCore/respond"
)

type userController struct{}

var User = &userController{}

func (*userController) AdminUser(c *gin.Context) {
	var q UserRequest
	err := c.ShouldBindUri(&q)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	data, err := logic.AdminUser.AdminUser(q.UserId)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}

func (*userController) AdminUserList(c *gin.Context) {
	var page logic.Page
	err := c.ShouldBind(&page)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	data, err := logic.AdminUser.AdminUserList(page)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}

func (*userController) AdminUserCreate(c *gin.Context) {
	var user entity.AdminUser
	err := c.ShouldBind(&user)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}
	logic.AdminUser.Context(c)
	err = logic.AdminUser.AdminUserCreate(user)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data("ok"))
}

func (*userController) AdminUserUpdate(c *gin.Context) {
	var user entity.AdminUser
	err := c.ShouldBind(&user)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	logic.AdminUser.Context(c)
	err = logic.AdminUser.AdminUserUpdate(user)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data("ok"))
}

func (*userController) AdminUserDelete(c *gin.Context) {
	var q UserRequest
	err := c.ShouldBindUri(&q)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	logic.AdminUser.Context(c)
	err = logic.AdminUser.AdminUserDelete(q.UserId)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data("ok"))
}
