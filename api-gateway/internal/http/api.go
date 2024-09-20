package http

import (
	"flashSale_gateway/internal/http/handlers"

	_ "flashSale_gateway/docs"
	m "flashSale_gateway/internal/http/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Flash Sale API Documentation
// @version 1.0
// @description API for Instant Delivery resources
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handlers.Handler) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// enforcer, err := casbin.NewEnforcer("./internal/http/casbin/model.conf", "./internal/http/casbin/policy.csv")
	// if err != nil {
	// 	panic(err)
	// }
	// router.Use(m.NewAuth(enforcer))

	router.POST("/register", h.RegisterUser).Use(m.Middleware())
	router.POST("/login", h.LoginUser).Use(m.Middleware())
	router.POST("/forgot-password", h.ForgotPassword)
	router.POST("/reset-password", h.ResetPassword)
	router.GET("/users", h.GetAllUsers).Use(m.JWTMiddleware())

	user := router.Group("/v1/user").Use(m.JWTMiddleware())
	{
		user.GET("/profiles", h.GetProfile)
		user.PUT("/profiles", h.EditProfile)
		user.PUT("/passwords", h.ChangePassword)
		user.GET("/setting", h.GetSetting)
		user.PUT("/setting", h.EditSetting)
		user.DELETE("/", h.DeleteUser)
	}

	flashSaleProduct := router.Group("/v1/flashSaleProduct")
	{
		flashSaleProduct.POST("/create", h.CreateFlashSaleProduct)
		flashSaleProduct.GET("/:id", h.GetFlashSaleProduct)
		flashSaleProduct.GET("/list", h.ListFlashSaleProducts)
		flashSaleProduct.PUT("/update/:id", h.UpdateFlashSaleProduct)
		flashSaleProduct.DELETE("/delete/:id", h.DeleteFlashSaleProduct)
	}
	flashSale := router.Group("/v1/flashSale")
	{
		flashSale.POST("/create", h.CreateFlashSale)
		flashSale.GET("/:id", h.GetFlashSale)
		flashSale.GET("/list", h.ListFlashSales)
		flashSale.PUT("/update/:id", h.UpdateFlashSale)
		flashSale.DELETE("/delete/:id", h.DeleteFlashSale)

		flashSale.GET("/:id/location", h.GetStoreLocation)
		flashSale.POST("/products", h.AddProductToFlashSale)
		flashSale.DELETE("/products", h.RemoveProductFromFlashSale)
		flashSale.POST("/:id/cancel", h.CancelFlashSale)
	}
	order := router.Group("/v1/order")
	{
		order.POST("/create", h.CreateOrder)
		order.GET("/:id", h.GetOrder)
		order.GET("/list", h.ListOrders)
		order.PUT("/update/:id", h.UpdateOrder)
		order.DELETE("/delete/:id", h.DeleteOrder)

		order.GET("/history", h.GetOrderHistory)
		order.POST("/:id/cancel", h.CancelOrder)
	}
	product := router.Group("/v1/product")
	{
		product.POST("/create", h.CreateProduct)
		product.GET("/:id", h.GetProduct)
		product.GET("/list", h.ListProducts)
		product.PUT("/update/:id", h.UpdateProduct)
		product.DELETE("/delete/:id", h.DeleteProduct)
	}
	notifications := router.Group("/v1/notification")
	{
		notifications.POST("/create", h.CreateNotification)
		notifications.GET("/:id", h.GetNotification)
		notifications.PUT("/update/:id", h.UpdateNotification)
		notifications.DELETE("/delete/:id", h.DeleteNotification)
		notifications.GET("/list", h.ListNotifications)
	}
	review := router.Group("/v1/reviews")
	{
		review.POST("", h.CreateReview)
		review.GET("/productId/rating", h.GetProductRating)
	}
	social := router.Group("/v1/deals")
	{
		social.POST("/share", h.ShareDeal)
		social.GET("/:flashSaleId/sharing", h.GetSharingStats)
		
	}


	return router
}
