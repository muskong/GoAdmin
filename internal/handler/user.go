package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoAdmin/internal/logic"
	"github.com/muskong/GoCore/respond"
)

func AdminUser(c *gin.Context) {
	var q struct {
		UserId int `json:"userId"`
	}
	err := c.ShouldBindUri(&q)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	data, err := logic.User.AdminUser(q.UserId)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}

func AdminUserList(c *gin.Context) {
	var page logic.Page
	err := c.ShouldBind(&page)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	err, data := logic.User.AdminUserList(page)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}

func AdminUserCreate(c *gin.Context) {
	var user entity.AdminUser
	err := c.ShouldBind(&user)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	err = logic.User.AdminUserCreate(user)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data("ok"))
}

func AdminUserUpdate(c *gin.Context) {
	var user entity.AdminUser
	err := c.ShouldBind(&user)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	err = logic.User.AdminUserUpdate(user)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data("ok"))
}

func AdminUserDelete(c *gin.Context) {
	var q struct {
		UserId int `json:"userId"`
	}
	err := c.ShouldBindUri(&q)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	err = logic.User.AdminUserDelete(q.UserId)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data("ok"))
}
