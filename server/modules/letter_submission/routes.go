package letter_submission

import (
	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/letter_submission/controller"
	letterRepository "bemunair2026/server/modules/letter_submission/repository"
	"bemunair2026/server/modules/letter_submission/service"
	userRepository "bemunair2026/server/modules/user/repository"
	"bemunair2026/server/pkg"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	api *gin.RouterGroup,
	repo letterRepository.LetterSubmissionRepository,
	users userRepository.UserRepository,
	wa pkg.WASender,
	jwtSecret string,
) {
	letterService := service.NewLetterSubmissionService(repo, users, wa)
	letterController := controller.NewLetterSubmissionController(letterService)

	submissions := api.Group("/letter-submissions", middlewares.Auth(jwtSecret))
	submissions.POST("", letterController.Create)
	submissions.GET("", letterController.List)
	submissions.GET("/:id", letterController.Get)

	admin := submissions.Group("", middlewares.AdminOnly())
	admin.PUT("/:id/status", letterController.UpdateStatus)
	admin.DELETE("/:id", letterController.Delete)
}
