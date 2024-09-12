package api

import (
	"auth/api/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Auth service
// @version 1.0
// @description Auth service
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

	admin := r.Group("/admin")
	{
		admin.DELETE("/delete-user", h.DeleteUser)
		admin.GET("/list-users", h.ListAllUsers)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/register-user", h.RegisterUser)
		auth.POST("/login", h.Login)
		auth.POST("/register-courier", h.RegisterCourier)
		auth.POST("/verify-email", h.VerifyEmail)
	}

	pro := r.Group("/profile")
	{
		pro.GET("get-profile", h.GetProfile)
		pro.PUT("update-profile", h.UpdateProfile)
		pro.DELETE("delete-profile", h.DeleteProfile)
	}

	add := r.Group("/additional")
	{
		add.PUT("/refresh-token", h.RefreshToken)
		add.PUT("/change-password", h.ChangePassword)
		add.POST("/forget-password", h.ForgetPassword)
		add.PUT("/reset-password", h.ResetPassword)
		add.PUT("/change-email", h.ChangeEmail)
		add.POST("/verify-new-email", h.VerifyNewEmail)
	}

	swaggerURL := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, swaggerURL))

	return r
}
