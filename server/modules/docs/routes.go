package docs

import "github.com/gin-gonic/gin"

func RegisterRoutes(api *gin.RouterGroup, docsDir string) {
	handler := NewHandler(docsDir)

	docs := api.Group("/docs")
	docs.GET("", handler.Index)
	docs.GET("/:slug", handler.Show)
}
