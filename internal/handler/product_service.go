package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoAdmin/internal/logic"
	"github.com/muskong/GoCore/respond"
	"github.com/muskong/GoPkg/zaplog"
)

type _productService struct{}

var ProductServiceHandler = &_productService{}

func (*_productService) ProductServices(c *gin.Context) {
	var page logic.Page
	err := c.ShouldBind(&page)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	data, err := logic.ProductService.List(page)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}

func (*_productService) ProductServiceCreate(c *gin.Context) {
	var data entity.ProductService
	err := c.ShouldBind(&data)
	if err != nil {
		zaplog.Sugar.Info(err)
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	logic.ProductService.Context(c)
	err = logic.ProductService.Create(data)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}

func (*_productService) ProductServiceUpdate(c *gin.Context) {
	var data entity.ProductService
	err := c.ShouldBind(&data)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	logic.ProductService.Context(c)
	err = logic.ProductService.Update(data)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}

func (*_productService) ProductServiceDelete(c *gin.Context) {
	var q ProductRequest
	err := c.ShouldBind(&q)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	logic.ProductService.Context(c)
	err = logic.ProductService.Delete(q.Id)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data("ok"))
}

func (*_productService) ProductServiceInstall(c *gin.Context) {
	data, err := logic.ProductService.Install()

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}
	c.SecureJSON(respond.Data(data))
}
