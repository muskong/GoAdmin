package internal

import (
	"github.com/muskong/GoAdmin/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/muskong/GoCore/middlewares"
	"github.com/muskong/GoPkg/jwt"
)

func GinRouter() *gin.Engine {
	tokenName := jwt.Jwt.GetTokenName()
	notAuth := map[string]bool{
		"/user/login": true,
	}

	router := gin.Default()
	router.Use(middlewares.GinCORS())
	router.Use(middlewares.GinUserMiddleware(tokenName, notAuth))

	adminUser := router.Group("/user")
	{
		adminUser.POST("/login", handler.AuthLogin)
		adminUser.GET("/list", handler.AdminUserList)
		adminUser.POST("/create", handler.AdminUserCreate)
	}

	adminRole := router.Group("/role")
	{
		adminRole.GET("/rules", handler.AdminRoleRuleList)
		adminRole.GET("/all", handler.AdminRoleAll)
		adminRole.GET("/list", handler.AdminRoleList)
		adminRole.POST("/create", handler.AdminRoleCreate)
		adminRole.POST("/saveRule", handler.AdminRoleSaveRule)
	}

	adminRule := router.Group("/rule")
	{
		adminRule.GET("/all", handler.AdminRuleAll)
		adminRule.GET("/list", handler.AdminRuleList)
		adminRule.POST("/create", handler.AdminRuleCreate)
	}

	return router
}
