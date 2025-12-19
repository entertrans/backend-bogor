package router

import (
	"github.com/gin-gonic/gin"

	"github.com/entertrans/go-base-project.git/internal/handler"
	"github.com/entertrans/go-base-project.git/internal/middleware"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, authHandler handler.AuthHandler, jwtSecret string) {
	// Public auth routes
	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
		// Untuk menambahkan Refresh, ForgotPassword, dll
		// Anda perlu menambahkan method tersebut ke interface AuthHandler terlebih dahulu
	}

	// Protected routes with authentication
	protectedGroup := rg.Group("/")
	protectedGroup.Use(middleware.AuthMiddleware(jwtSecret))
	{
		protectedGroup.GET("/profile", authHandler.Profile)
		// Untuk menambahkan UpdateProfile, Logout, dll
		// Anda perlu menambahkan method tersebut ke interface AuthHandler terlebih dahulu
	}
}