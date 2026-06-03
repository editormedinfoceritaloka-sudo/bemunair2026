package user

import (
	"bemunair2026/server/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup, repo *Repository, jwtSecret string) {
	handler := NewHandler(repo)

	users := api.Group("/users", middlewares.Auth(jwtSecret), middlewares.AdminOnly())
	users.GET("", handler.List)
	users.GET("/:id", handler.Get)
	users.POST("", handler.Create)
	users.PUT("/:id", handler.Update)
	users.DELETE("/:id", handler.Delete)
}
