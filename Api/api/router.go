package api

import (
	"api/api/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Voting service
// @version 1.0
// @description Delivery service
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	cart := r.Group("/cart")
	{
		cart.POST("/create-cart", h.CreateCart)
		cart.GET("/get-cart", h.GetCart)
		cart.GET("/get-cart-by-user", h.GetCartByUser)
		cart.PUT("/update-cart", h.UpdateCart)
		cart.DELETE("/delete-cart", h.DeleteCart)
		cart.POST("/create-cart-item", h.CreateCartItem)
		cart.GET("/get-cart-item", h.GetCartItem)
		cart.DELETE("/delete-cart-item", h.DeleteCartItem)
		cart.PUT("/update-cart-item-quantity", h.UpdateCartItemQuantity)
	}

	office := r.Group("/office")
	{
		office.POST("/create-office", h.CreateOffice)
		office.GET("/get-office", h.GetOffice)
		office.GET("/get-all-ofice", h.GetAllOffice)
		office.PUT("/update-office", h.UpdateOffice)
		office.DELETE("/delete-office", h.DeleteOffice)
	}

	order := r.Group("/order")
	{
		order.POST("/create-order", h.CreateOrder)
		order.GET("/get-order", h.GetOrder)
		order.GET("/get-all-orders", h.GetAllOrders)
		order.GET("/get-order-by-client", h.GetOrderByClient)
		order.PUT("/update-order", h.UpdateOrder)
		order.DELETE("/delete-order", h.DeleteOrder)
	}
	product := r.Group("/product")
	{
		product.POST("/create-product", h.CreateProduct)
		product.GET("/get-product", h.GetProduct)
		product.GET("/get-all-products", h.GetAllProduct)
		product.PUT("/update-product", h.UpdateProduct)
		product.DELETE("/delete-product", h.DeleteProduct)
	}

	swaggerURL := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, swaggerURL))

	return r
}
