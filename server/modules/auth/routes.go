package auth

import (
	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/auth/controller"
	"bemunair2026/server/modules/auth/service"
	"bemunair2026/server/modules/user/repository"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup, users repository.UserRepository, jwtSecret string) {
	authService := service.NewAuthService(users, jwtSecret)
	authController := controller.NewAuthController(authService)

	auth := api.Group("/auth")
	auth.POST("/register", authController.Register)
	auth.POST("/login", authController.Login)
	auth.GET("/me", middlewares.Auth(jwtSecret), authController.Me)
}
