package docs

import (
	"bemunair2026/server/modules/docs/controller"
	"bemunair2026/server/modules/docs/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup, docsDir string) {
	docsService := service.NewDocsService(docsDir)
	docsController := controller.NewDocsController(docsService)

	docs := api.Group("/docs")
	docs.GET("", docsController.Index)
	docs.GET("/:slug", docsController.Show)
}
