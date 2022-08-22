package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoAdmin/internal/logic"
	"github.com/muskong/GoCore/respond"
)

type _product struct{}

var ProductHandler = &_product{}

func (*_product) ProductSelect(c *gin.Context) {
	data := logic.Product.Selects()

	c.SecureJSON(respond.Data(data))
}

func (*_product) Products(c *gin.Context) {
	var page logic.Page
	err := c.ShouldBind(&page)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	data, err := logic.Product.List(page)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}

func (*_product) ProductCreate(c *gin.Context) {
	var data entity.Product
	err := c.ShouldBind(&data)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	logic.Product.Context(c)
	err = logic.Product.Create(data)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}

func (*_product) ProductUpdate(c *gin.Context) {
	var data entity.Product
	err := c.ShouldBind(&data)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	logic.Product.Context(c)
	err = logic.Product.Update(data)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}

func (*_product) ProductDelete(c *gin.Context) {
	var q ProductRequest
	err := c.ShouldBind(&q)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	logic.Product.Context(c)
	err = logic.Product.Delete(q.Id)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}
