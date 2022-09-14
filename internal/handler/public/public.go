package public

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoCore/respond"
)

func Index(ctx *gin.Context) {
	ctx.SecureJSON(respond.Data("best"))
}
