package article

import (
	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/article/controller"
	"bemunair2026/server/modules/article/repository"
	"bemunair2026/server/modules/article/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup, repo repository.ArticleRepository, jwtSecret string) {
	articleService := service.NewArticleService(repo)
	articleController := controller.NewArticleController(articleService)

	public := api.Group("/articles")
	public.GET("", articleController.ListPublished)
	public.GET("/:slug", articleController.GetBySlug)

	admin := api.Group("/admin/articles", middlewares.Auth(jwtSecret), middlewares.AdminOnly())

	admin.GET("", articleController.ListAll)
	admin.POST("", articleController.Create)
	admin.GET("/:id", articleController.GetByID)
	admin.PUT("/:id", articleController.Update)
	admin.PUT("/:id/publish", articleController.SetPublished)
	admin.DELETE("/:id", articleController.Delete)
}
