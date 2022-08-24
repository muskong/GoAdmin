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
		"/admin/admin/login": true,
		// "/admin/sites":      true,
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
	adm.GET("/logs", handler.Log.AdminLogList)

	adminUser := adm.Group("/admin")
	{
		adminUser.POST("/login", handler.Auth.Login)

		adminUser.GET(":nanoid", handler.User.AdminUser)
		adminUser.GET("/list", handler.User.AdminUserList)
		adminUser.POST("/create", handler.User.AdminUserCreate)
		adminUser.POST("/update", handler.User.AdminUserUpdate)
		adminUser.DELETE("/delete/:nanoid", handler.User.AdminUserDelete)
	}

	adminRole := adm.Group("/role")
	{
		// admin/auth.group/index
		adminRole.GET("/group", handler.Role.AdminRoleGroup)
		adminRole.GET("/tree", handler.Role.AdminRoleTree)

		adminRole.GET("/rules", handler.Role.AdminRoleRuleList)
		adminRole.GET("/all", handler.Role.AdminRoleAll)
		adminRole.POST("/saveRule", handler.Role.AdminRoleSaveRule)

		adminRole.GET(":nanoid", handler.Role.AdminRole)
		adminRole.GET("/list", handler.Role.AdminRoleList)
		adminRole.POST("/create", handler.Role.AdminRoleCreate)
		adminRole.POST("/update", handler.Role.AdminRoleUpdate)
		adminRole.DELETE("/delete/:nanoid", handler.Role.AdminRoleDelete)
	}

	adminRule := adm.Group("/rule")
	{
		adminRule.GET("/group", handler.Rule.AdminRuleGroup)
		adminRule.GET("/tree", handler.Rule.AdminRuleTree)
		adminRule.GET("/all", handler.Rule.AdminRuleAll)

		adminRule.GET(":nanoid", handler.Rule.AdminRule)
		adminRule.GET("/list", handler.Rule.AdminRuleList)
		adminRule.POST("/create", handler.Rule.AdminRuleCreate)
		adminRule.POST("/update", handler.Rule.AdminRuleUpdate)
		adminRule.DELETE("/delete/:nanoid", handler.Rule.AdminRuleDelete)
	}

	productRule := adm.Group("/product")
	{
		productRule.GET("/all", handler.ProductHandler.ProductSelect)
		productRule.GET("/list", handler.ProductHandler.Products)
		productRule.POST("/create", handler.ProductHandler.ProductCreate)
		productRule.POST("/update", handler.ProductHandler.ProductUpdate)
		productRule.DELETE("/delete/:id", handler.ProductHandler.ProductDelete)
	}
	productAmountRule := adm.Group("/productAmount")
	{
		productAmountRule.GET("/list", handler.ProductAmountHandler.ProductAmounts)
		productAmountRule.POST("/create", handler.ProductAmountHandler.ProductAmountCreate)
		productAmountRule.POST("/update", handler.ProductAmountHandler.ProductAmountUpdate)
		productAmountRule.DELETE("/delete/:id", handler.ProductAmountHandler.ProductAmountDelete)
	}

	productCardRule := adm.Group("/productCard")
	{
		productCardRule.GET("/list", handler.ProductCardHandler.ProductCards)
		productCardRule.POST("/create", handler.ProductCardHandler.ProductCardCreate)
		productCardRule.POST("/update", handler.ProductCardHandler.ProductCardUpdate)
		productCardRule.DELETE("/delete/:id", handler.ProductCardHandler.ProductCardDelete)
	}

	productChannelRule := adm.Group("/productChannel")
	{
		productChannelRule.GET("/list", handler.ProductChannelHandler.ProductChannels)
		productChannelRule.POST("/create", handler.ProductChannelHandler.ProductChannelCreate)
		productChannelRule.POST("/update", handler.ProductChannelHandler.ProductChannelUpdate)
		productChannelRule.DELETE("/delete/:id", handler.ProductChannelHandler.ProductChannelDelete)
	}

	productServiceRule := adm.Group("/productService")
	{
		productServiceRule.GET("/install", handler.ProductServiceHandler.ProductServiceInstall)
		productServiceRule.GET("/list", handler.ProductServiceHandler.ProductServices)
		productServiceRule.POST("/create", handler.ProductServiceHandler.ProductServiceCreate)
		productServiceRule.POST("/update", handler.ProductServiceHandler.ProductServiceUpdate)
		productServiceRule.DELETE("/delete/:id", handler.ProductServiceHandler.ProductServiceDelete)
	}

	return router
}
