package controller

import (
	"net/http"

	"bemunair2026/server/modules/docs/service"
	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type DocsController interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
}

type docsController struct {
	service service.DocsService
}

var _ DocsController = (*docsController)(nil)

func NewDocsController(service service.DocsService) DocsController {
	return &docsController{service: service}
}

func (c *docsController) Index(ctx *gin.Context) {
	body, err := c.service.Index()
	if err != nil {
		res := response.BuildResponseFailed("Index docs tidak ditemukan", response.NotFound, nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}
	ctx.Data(http.StatusOK, "application/json; charset=utf-8", body)
}

func (c *docsController) Show(ctx *gin.Context) {
	body, err := c.service.Show(ctx.Param("slug"))
	if err != nil {
		res := response.BuildResponseFailed("Docs tidak ditemukan", response.NotFound, nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}
	ctx.Data(http.StatusOK, "text/markdown; charset=utf-8", body)
}
