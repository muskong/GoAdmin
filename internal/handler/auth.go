package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/logic"
	"github.com/muskong/GoCore/respond"
)

type authController struct{}

var Auth = &authController{}

func (*authController) Login(c *gin.Context) {
	var userForm logic.LoginData
	err := c.ShouldBind(&userForm)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}

	data, err := logic.LoginVerify(userForm)

	if err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(data))
}
