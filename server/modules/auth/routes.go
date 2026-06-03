package auth

import (
	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/auth/handler"
	"bemunair2026/server/modules/auth/repository"
	"bemunair2026/server/modules/auth/service"
	"bemunair2026/server/modules/auth/validation"
	"bemunair2026/server/modules/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup, users *user.Repository, jwtSecret string) {
	authRepository := repository.NewRepository(users)
	authService := service.NewService(authRepository, jwtSecret)
	authValidation := validation.NewAuthValidation()
	authHandler := handler.NewHandler(authService, authValidation)

	auth := api.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
	auth.GET("/me", middlewares.Auth(jwtSecret), authHandler.Me)
}
