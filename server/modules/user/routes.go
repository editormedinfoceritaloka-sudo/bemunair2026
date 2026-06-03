package user

import (
	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/user/controller"
	"bemunair2026/server/modules/user/repository"
	"bemunair2026/server/modules/user/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup, repo repository.UserRepository, jwtSecret string) {
	userService := service.NewUserService(repo)
	userController := controller.NewUserController(userService)

	users := api.Group("/users", middlewares.Auth(jwtSecret), middlewares.AdminOnly())
	users.GET("", userController.List)
	users.GET("/:id", userController.Get)
	users.POST("", userController.Create)
	users.PUT("/:id", userController.Update)
	users.DELETE("/:id", userController.Delete)
}
