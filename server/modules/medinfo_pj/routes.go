package medinfo_pj

import (
	"bemunair2026/server/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup, repo *Repository, jwtSecret string) {
	handler := NewHandler(repo)

	queue := api.Group("/medinfo-pj/queue", middlewares.Auth(jwtSecret), middlewares.AdminOnly())
	queue.GET("", handler.List)
	queue.POST("", handler.Create)
	queue.PUT("/reorder", handler.Reorder)
	queue.DELETE("/:id", handler.Delete)
}
