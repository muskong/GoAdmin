package internal

import (
	"github.com/muskong/GoAdmin/internal/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoCore/middlewares"
	"github.com/muskong/GoPkg/jwt"
)

func GinRouter() *gin.Engine {
	tokenName := jwt.Jwt.GetTokenName()
	notAuth := map[string]bool{
		"/admin/user/login": true,
		"/admin/sites":      true,
	}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"lang", "atoken", "utoken", "Origin", "Content-Length", "Content-Type"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))
	// router.Use(cors.Default())
	router.Use(middlewares.GinUserMiddleware(tokenName, notAuth))

	router.GET("/", handler.Index)
	adm := router.Group("/admin")

	adm.GET("/sites", handler.Sites)
	adm.GET("/dashboard", handler.Dashboard)

	adminUser := adm.Group("/user")
	{
		adminUser.POST("/login", handler.AuthLogin)

		adminUser.GET(":userId", handler.AdminUser)
		adminUser.GET("/list", handler.AdminUserList)
		adminUser.POST("/create", handler.AdminUserCreate)
		adminUser.POST("/update", handler.AdminUserUpdate)
		adminUser.DELETE("/delete/:userId", handler.AdminUserDelete)
	}

	adminRole := adm.Group("/role")
	{
		// admin/auth.group/index
		adminRole.GET("/group", handler.Role.AdminRoleGroup)

		adminRole.GET("/rules", handler.Role.AdminRoleRuleList)
		adminRole.GET("/all", handler.Role.AdminRoleAll)
		adminRole.POST("/saveRule", handler.Role.AdminRoleSaveRule)

		adminRole.GET(":roleId", handler.Role.AdminRole)
		adminRole.GET("/list", handler.Role.AdminRoleList)
		adminRole.POST("/create", handler.Role.AdminRoleCreate)
		adminRole.POST("/update", handler.Role.AdminRoleUpdate)
		adminRole.DELETE("/delete/:roleId", handler.Role.AdminRoleDelete)
	}

	adminRule := adm.Group("/rule")
	{
		adminRule.GET("/menu", handler.Rule.AdminMenu)
		adminRule.GET("/all", handler.Rule.AdminRuleAll)

		adminRule.GET(":ruleId", handler.Rule.AdminRule)
		adminRule.GET("/list", handler.Rule.AdminRuleList)
		adminRule.POST("/create", handler.Rule.AdminRuleCreate)
		adminRule.POST("/update", handler.Rule.AdminRuleUpdate)
		adminRule.DELETE("/delete/:ruleId", handler.Rule.AdminRuleDelete)
	}

	return router
}
