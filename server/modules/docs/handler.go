package docs

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct{ dir string }

func NewHandler(dir string) *Handler { return &Handler{dir: dir} }

func (h *Handler) Index(c *gin.Context) {
	b, err := os.ReadFile(filepath.Join(h.dir, "index.json"))
	if err != nil {
		response.Error(c, http.StatusNotFound, response.NotFound, "Index docs tidak ditemukan")
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", b)
}

func (h *Handler) Show(c *gin.Context) {
	slug := strings.Trim(c.Param("slug"), "/")
	if strings.Contains(slug, "..") {
		response.Error(c, http.StatusNotFound, response.NotFound, "Docs tidak ditemukan")
		return
	}
	b, err := os.ReadFile(filepath.Join(h.dir, slug+".md"))
	if err != nil {
		response.Error(c, http.StatusNotFound, response.NotFound, "Docs tidak ditemukan")
		return
	}
	c.Data(http.StatusOK, "text/markdown; charset=utf-8", b)
}
