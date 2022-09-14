package public

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoCore/respond"
)

type _order struct{}

var OrderHandler = &_order{}

func (*_order) Notify(ctx *gin.Context) {
	var q Jwt
	err := ctx.ShouldBindUri(&q)
	if err != nil {
		ctx.AbortWithStatusJSON(respond.Message("传入参数错误"))
		return
	}

	ctx.SecureJSON(respond.Data("best"))
}
