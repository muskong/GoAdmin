package internal

import (
	"github.com/muskong/GoAdmin/internal/handler/admin"
	"github.com/muskong/GoAdmin/internal/handler/public"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoCore/middlewares"
	"github.com/muskong/GoPkg/jwt"
)

type Flags struct {
	Mode string
	Port int
}

func GinRouter(opts *Flags) *gin.Engine {
	tokenName := jwt.Jwt.GetTokenName()
	notAuth := map[string]bool{
		"/admin/admin/login": true,
		// "/admin/sites":      true,
	}

	gin.SetMode(opts.Mode)
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"lang", "atoken", "utoken", "Origin", "Content-Length", "Content-Type"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))
	// router.Use(cors.Default())
	// router.Use(middlewares.GinUserMiddleware(tokenName, notAuth))

	router.GET("/", admin.Index)
	adm := router.Group("/admin")
	{
		adm.Use(middlewares.GinAdminMiddleware(tokenName, notAuth))

		adm.GET("/sites", admin.Sites)
		adm.GET("/dashboard", admin.Dashboard)
		adm.GET("/logs", admin.Log.AdminLogList)

		adminUser := adm.Group("/admin")
		{
			adminUser.POST("/login", admin.Auth.Login)

			adminUser.GET(":nanoid", admin.User.AdminUser)
			adminUser.GET("/list", admin.User.AdminUserList)
			adminUser.POST("/create", admin.User.AdminUserCreate)
			adminUser.POST("/update", admin.User.AdminUserUpdate)
			adminUser.DELETE("/delete/:nanoid", admin.User.AdminUserDelete)
		}

		adminRole := adm.Group("/role")
		{
			// admin/auth.group/index
			adminRole.GET("/group", admin.Role.AdminRoleGroup)
			adminRole.GET("/tree", admin.Role.AdminRoleTree)

			adminRole.GET("/rules", admin.Role.AdminRoleRuleList)
			adminRole.GET("/all", admin.Role.AdminRoleAll)
			adminRole.POST("/saveRule", admin.Role.AdminRoleSaveRule)

			adminRole.GET(":nanoid", admin.Role.AdminRole)
			adminRole.GET("/list", admin.Role.AdminRoleList)
			adminRole.POST("/create", admin.Role.AdminRoleCreate)
			adminRole.POST("/update", admin.Role.AdminRoleUpdate)
			adminRole.DELETE("/delete/:nanoid", admin.Role.AdminRoleDelete)
		}

		adminRule := adm.Group("/rule")
		{
			adminRule.GET("/group", admin.Rule.AdminRuleGroup)
			adminRule.GET("/tree", admin.Rule.AdminRuleTree)
			adminRule.GET("/all", admin.Rule.AdminRuleAll)

			adminRule.GET(":nanoid", admin.Rule.AdminRule)
			adminRule.GET("/list", admin.Rule.AdminRuleList)
			adminRule.POST("/create", admin.Rule.AdminRuleCreate)
			adminRule.POST("/update", admin.Rule.AdminRuleUpdate)
			adminRule.DELETE("/delete/:nanoid", admin.Rule.AdminRuleDelete)
		}

		productRule := adm.Group("/product")
		{
			productRule.GET("/all", admin.ProductHandler.ProductSelect)
			productRule.GET("/list", admin.ProductHandler.Products)
			productRule.POST("/create", admin.ProductHandler.ProductCreate)
			productRule.POST("/update", admin.ProductHandler.ProductUpdate)
			productRule.DELETE("/delete/:id", admin.ProductHandler.ProductDelete)
		}
		productAmountRule := adm.Group("/productAmount")
		{
			productAmountRule.GET("/list", admin.ProductAmountHandler.ProductAmounts)
			productAmountRule.POST("/create", admin.ProductAmountHandler.ProductAmountCreate)
			productAmountRule.POST("/update", admin.ProductAmountHandler.ProductAmountUpdate)
			productAmountRule.DELETE("/delete/:id", admin.ProductAmountHandler.ProductAmountDelete)
		}

		productCardRule := adm.Group("/productCard")
		{
			productCardRule.GET("/list", admin.ProductCardHandler.ProductCards)
			productCardRule.POST("/create", admin.ProductCardHandler.ProductCardCreate)
			productCardRule.POST("/update", admin.ProductCardHandler.ProductCardUpdate)
			productCardRule.DELETE("/delete/:id", admin.ProductCardHandler.ProductCardDelete)
		}

		productChannelRule := adm.Group("/productChannel")
		{
			productChannelRule.GET("/list", admin.ProductChannelHandler.ProductChannels)
			productChannelRule.POST("/create", admin.ProductChannelHandler.ProductChannelCreate)
			productChannelRule.POST("/update", admin.ProductChannelHandler.ProductChannelUpdate)
			productChannelRule.DELETE("/delete/:id", admin.ProductChannelHandler.ProductChannelDelete)
		}

		productServiceRule := adm.Group("/productService")
		{
			productServiceRule.GET("/pluginList", admin.ProductServiceHandler.ProductServicePluginList)
			productServiceRule.GET("/plugin/:fileName", admin.ProductServiceHandler.ProductServicePlugin)
			productServiceRule.POST("/install", admin.ProductServiceHandler.ProductServiceInstall)
			productServiceRule.GET("/list", admin.ProductServiceHandler.ProductServices)
			productServiceRule.POST("/update", admin.ProductServiceHandler.ProductServiceUpdate)
			productServiceRule.DELETE("/delete/:id", admin.ProductServiceHandler.ProductServiceDelete)
		}

		userRule := adm.Group("/user")
		{
			userRule.GET("/list", admin.UserHandler.Users)
			userRule.POST("/create", admin.UserHandler.UserCreate)
			userRule.POST("/update", admin.UserHandler.UserUpdate)
			userRule.DELETE("/delete/:id", admin.UserHandler.UserDelete)
		}
		userAccountRule := adm.Group("/userAccount")
		{
			userAccountRule.GET("/list", admin.UserAccountHandler.UserAccounts)
			userAccountRule.POST("/create", admin.UserAccountHandler.UserAccountCreate)
		}
		userBankRule := adm.Group("/userBank")
		{
			userBankRule.GET("/list", admin.UserBankHandler.UserBanks)
			userBankRule.POST("/create", admin.UserBankHandler.UserBankCreate)
			userBankRule.POST("/update", admin.UserBankHandler.UserBankUpdate)
			userBankRule.DELETE("/delete/:id", admin.UserBankHandler.UserBankDelete)
		}
		userGroupRule := adm.Group("/userGroup")
		{
			userGroupRule.GET("/list", admin.UserGroupHandler.UserGroups)
			userGroupRule.POST("/create", admin.UserGroupHandler.UserGroupCreate)
			userGroupRule.POST("/update", admin.UserGroupHandler.UserGroupUpdate)
			userGroupRule.DELETE("/delete/:id", admin.UserGroupHandler.UserGroupDelete)
		}
		userVerifiedRule := adm.Group("/userVerified")
		{
			userVerifiedRule.GET("/list", admin.UserVerifiedHandler.UserVerifieds)
			userVerifiedRule.POST("/create", admin.UserVerifiedHandler.UserVerifiedCreate)
			userVerifiedRule.POST("/update", admin.UserVerifiedHandler.UserVerifiedUpdate)
			userVerifiedRule.DELETE("/delete/:id", admin.UserVerifiedHandler.UserVerifiedDelete)
		}

		orderRule := adm.Group("/order")
		{
			orderRule.GET("/list", admin.OrderHandler.Orders)
			orderRule.POST("/create", admin.OrderHandler.OrderCreate)
			orderRule.POST("/update", admin.OrderHandler.OrderUpdate)
			orderRule.DELETE("/delete/:id", admin.OrderHandler.OrderDelete)
		}
	}
	api := router.Group("/api")
	{
		api.Use(middlewares.GinApiMiddleware(tokenName))
		api.GET("/", admin.Index)
	}
	usr := router.Group("/user")
	{
		usr.Use(middlewares.GinUserMiddleware(tokenName, notAuth))
		usr.GET("/", admin.Index)
	}
	pub := router.Group("/public")
	{
		pub.GET("/", public.Index)
		pub.Any("notify/:token", public.OrderHandler.Notify)
	}

	return router
}
