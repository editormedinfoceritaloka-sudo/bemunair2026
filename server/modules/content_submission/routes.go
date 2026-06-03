package content_submission

import (
	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/content_submission/controller"
	contentRepository "bemunair2026/server/modules/content_submission/repository"
	"bemunair2026/server/modules/content_submission/service"
	userRepository "bemunair2026/server/modules/user/repository"
	"bemunair2026/server/pkg"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	api *gin.RouterGroup,
	repo contentRepository.ContentSubmissionRepository,
	users userRepository.UserRepository,
	wa pkg.WASender,
	jwtSecret string,
) {
	contentService := service.NewContentSubmissionService(repo, users, wa)
	contentController := controller.NewContentSubmissionController(contentService)

	submissions := api.Group("/content-submissions", middlewares.Auth(jwtSecret))
	submissions.POST("", contentController.Create)
	submissions.GET("", contentController.List)
	submissions.GET("/:id", contentController.Get)

	admin := submissions.Group("", middlewares.AdminOnly())
	admin.PUT("/:id/status", contentController.UpdateStatus)
	admin.DELETE("/:id", contentController.Delete)
}
