package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoAdmin/internal/logic"
	"github.com/muskong/GoPkg/work"
)

func AdminUserList(c *gin.Context) {
	var page logic.Page
	err := work.Context.GetJson(&page)
	if err != nil {
		c.SecureJSON(http.StatusOK, "传入参数错误")
		return
	}

	err, data := logic.AdminUserList(page)

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error())
		return
	}

	c.SecureJSON(http.StatusOK, data)
}

func AdminUserCreate(c *gin.Context) {
	var user entity.AdminUserObject
	err := work.Context.GetJson(&user)
	if err != nil {
		c.SecureJSON(http.StatusOK, "传入参数错误")
		return
	}

	err = logic.AdminUserCreate(user)

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error())
		return
	}

	c.SecureJSON(http.StatusOK, "ok")
}
