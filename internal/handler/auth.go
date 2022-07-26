package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/logic"
)

func AuthLogin(c *gin.Context) {
	var userForm logic.LoginData
	err := c.ShouldBind(&userForm)
	if err != nil {
		c.SecureJSON(http.StatusOK, "传入参数错误")
		return
	}

	data, err := logic.LoginVerify(userForm)

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error)
		return
	}

	c.SecureJSON(http.StatusOK, data)
}
