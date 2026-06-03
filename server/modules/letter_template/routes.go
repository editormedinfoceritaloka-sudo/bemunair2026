package letter_template

import (
	"bemunair2026/server/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup, repo *Repository, jwtSecret string) {
	handler := NewHandler(repo)

	templates := api.Group("/letter-templates", middlewares.Auth(jwtSecret), middlewares.AdminOnly())
	templates.POST("", handler.Create)
	templates.GET("", handler.List)
	templates.GET("/:id", handler.Get)
	templates.PUT("/:id", handler.Update)
	templates.DELETE("/:id", handler.Delete)
}
