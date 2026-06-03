package letter_template

import (
	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/letter_template/controller"
	"bemunair2026/server/modules/letter_template/repository"
	"bemunair2026/server/modules/letter_template/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup, repo repository.LetterTemplateRepository, jwtSecret string) {
	templateService := service.NewLetterTemplateService(repo)
	templateController := controller.NewLetterTemplateController(templateService)

	templates := api.Group("/letter-templates", middlewares.Auth(jwtSecret), middlewares.AdminOnly())
	templates.POST("", templateController.Create)
	templates.GET("", templateController.List)
	templates.GET("/:id", templateController.Get)
	templates.PUT("/:id", templateController.Update)
	templates.DELETE("/:id", templateController.Delete)
}
