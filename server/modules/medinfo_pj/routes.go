package medinfo_pj

import (
	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/medinfo_pj/controller"
	"bemunair2026/server/modules/medinfo_pj/repository"
	"bemunair2026/server/modules/medinfo_pj/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup, repo repository.MedinfoPJRepository, jwtSecret string) {
	queueService := service.NewMedinfoPJService(repo)
	queueController := controller.NewMedinfoPJController(queueService)

	queue := api.Group("/medinfo-pj/queue", middlewares.Auth(jwtSecret), middlewares.AdminOnly())
	queue.GET("", queueController.List)
	queue.POST("", queueController.Create)
	queue.PUT("/reorder", queueController.Reorder)
	queue.DELETE("/:id", queueController.Delete)
}
