package content_submission

import (
	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/user/repository"
	"bemunair2026/server/pkg"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup, repo *Repository, users repository.UserRepository, wa pkg.WASender, jwtSecret string) {
	handler := NewHandler(repo, NewService(repo, wa), users)

	submissions := api.Group("/content-submissions", middlewares.Auth(jwtSecret))
	submissions.POST("", handler.Create)
	submissions.GET("", handler.List)
	submissions.GET("/:id", handler.Get)

	admin := submissions.Group("", middlewares.AdminOnly())
	admin.PUT("/:id/status", handler.UpdateStatus)
	admin.DELETE("/:id", handler.Delete)
}
