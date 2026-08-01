package cabinet

import (
	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/cabinet/controller"
	"bemunair2026/server/modules/cabinet/repository"
	"bemunair2026/server/modules/cabinet/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup, dbRepository repository.Repository, jwtSecret string) {
	c := controller.New(service.New(dbRepository))

	public := api.Group("/cabinet")
	public.GET("", c.ActiveCabinet)
	public.GET("/:slug", c.CabinetBySlug)
	public.GET("/units/:slug", c.PublicUnit)
	public.GET("/units/:slug/programs", c.PublicPrograms)
	public.GET("/units/:slug/programs/:programSlug", c.PublicProgram)

	authenticated := api.Group("", middlewares.Auth(jwtSecret), middlewares.AuthenticatedAdmin())
	adminCabinets := authenticated.Group("/admin/cabinet-terms")
	adminCabinets.GET("", c.ListCabinets)
	adminCabinets.POST("", middlewares.MedinfoOnly(), c.CreateCabinet)
	adminCabinets.PUT("/:id", middlewares.MedinfoOnly(), c.UpdateCabinet)

	adminOrganizations := authenticated.Group("/admin/organizations")
	adminOrganizations.GET("", c.ListUnits)
	adminOrganizations.POST("", c.CreateUnit)
	adminOrganizations.PUT("/:id", c.UpdateUnit)
	adminOrganizations.GET("/:id/members", c.ListMembers)
	adminOrganizations.POST("/:id/members", c.CreateMember)
	adminOrganizations.PUT("/members/:id", c.UpdateMember)

	adminPrograms := authenticated.Group("/admin/work-programs")
	adminPrograms.GET("", c.ListPrograms)
	adminPrograms.GET("/:id", c.AdminProgram)
	adminPrograms.POST("", c.CreateProgram)
	adminPrograms.PUT("/:id", c.UpdateProgram)
	adminPrograms.PUT("/:id/publish", c.PublishProgram)
	adminPrograms.POST("/:id/milestones", c.CreateMilestone)
	adminPrograms.POST("/:id/documentations", c.CreateDocumentation)
	adminPrograms.PUT("/:id/documentations/reorder", c.ReorderDocumentation)

	authenticated.POST("/admin/media-assets", c.CreateMedia)
}
