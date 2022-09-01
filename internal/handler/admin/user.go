package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoAdmin/internal/logic"
	"github.com/muskong/GoCore/respond"
)

type _user struct{}

var UserHandler = &_user{}

func (*_user) Users(c *gin.Context) {
	var page logic.Page
	err := c.ShouldBind(&page)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	data, err := logic.User.List(page)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}

func (*_user) UserCreate(c *gin.Context) {
	var data entity.User
	err := c.ShouldBind(&data)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	logic.User.Context(c)
	err = logic.User.Create(data)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}

func (*_user) UserUpdate(c *gin.Context) {
	var data entity.User
	err := c.ShouldBind(&data)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	logic.User.Context(c)
	err = logic.User.Update(data)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}

func (*_user) UserDelete(c *gin.Context) {
	var q ProductRequest
	err := c.ShouldBindUri(&q)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	logic.User.Context(c)
	err = logic.User.Delete(q.Id)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}
