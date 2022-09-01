package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoAdmin/internal/logic"
	"github.com/muskong/GoCore/respond"
)

type _userVerified struct{}

var UserVerifiedHandler = &_userVerified{}

func (*_userVerified) UserVerifieds(c *gin.Context) {
	var page logic.Page
	err := c.ShouldBind(&page)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	data, err := logic.UserVerified.List(page)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}

func (*_userVerified) UserVerifiedCreate(c *gin.Context) {
	var data entity.UserVerified
	err := c.ShouldBind(&data)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	logic.UserVerified.Context(c)
	err = logic.UserVerified.Create(data)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}

func (*_userVerified) UserVerifiedUpdate(c *gin.Context) {
	var data entity.UserVerified
	err := c.ShouldBind(&data)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	logic.UserVerified.Context(c)
	err = logic.UserVerified.Update(data)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}

func (*_userVerified) UserVerifiedDelete(c *gin.Context) {
	var q ProductRequest
	err := c.ShouldBindUri(&q)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	logic.UserVerified.Context(c)
	err = logic.UserVerified.Delete(q.Id)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}
