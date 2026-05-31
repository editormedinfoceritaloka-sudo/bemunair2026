package main

import (
	"log"
	"os"

	"bemunair2026/server/config"
	"bemunair2026/server/database"
	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/auth"
	content "bemunair2026/server/modules/content_submission"
	"bemunair2026/server/modules/cron"
	"bemunair2026/server/modules/docs"
	letter "bemunair2026/server/modules/letter_submission"
	"bemunair2026/server/modules/letter_template"
	"bemunair2026/server/modules/medinfo_pj"
	"bemunair2026/server/modules/user"
	"bemunair2026/server/pkg"
	"bemunair2026/server/pkg/response"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	_ = os.Setenv("TZ", cfg.TZ)
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	router := gin.New()
	router.Use(middlewares.RequestLogger(), middlewares.Recovery(), middlewares.CORS())

	userRepo := user.NewRepository(db)
	waClient := pkg.NewWAClient(cfg.WAURL, cfg.WAAPIKey)
	authService := auth.NewService(userRepo, cfg.JWTSecret)
	authHandler := auth.NewHandler(authService, userRepo)
	userHandler := user.NewHandler(userRepo)
	contentRepo := content.NewRepository(db)
	contentHandler := content.NewHandler(contentRepo, content.NewService(contentRepo, waClient), userRepo)
	letterRepo := letter.NewRepository(db)
	letterHandler := letter.NewHandler(letterRepo, letter.NewService(letterRepo, waClient), userRepo)
	queueHandler := medinfo_pj.NewHandler(medinfo_pj.NewRepository(db))
	templateHandler := letter_template.NewHandler(letter_template.NewRepository(db))
	docsHandler := docs.NewHandler(cfg.DocsDir)

	router.GET("/ping", func(c *gin.Context) { response.OK(c, "pong", nil) })
	router.POST("/api/auth/register", authHandler.Register)
	router.POST("/api/auth/login", authHandler.Login)
	router.GET("/api/docs", docsHandler.Index)
	router.GET("/api/docs/:slug", docsHandler.Show)

	api := router.Group("/api", middlewares.Auth(cfg.JWTSecret))
	api.GET("/auth/me", authHandler.Me)
	api.POST("/content-submissions", contentHandler.Create)
	api.GET("/content-submissions", contentHandler.List)
	api.GET("/content-submissions/:id", contentHandler.Get)
	api.POST("/letter-submissions", letterHandler.Create)
	api.GET("/letter-submissions", letterHandler.List)
	api.GET("/letter-submissions/:id", letterHandler.Get)

	admin := api.Group("", middlewares.AdminOnly())
	admin.GET("/users", userHandler.List)
	admin.GET("/users/:id", userHandler.Get)
	admin.POST("/users", userHandler.Create)
	admin.PUT("/users/:id", userHandler.Update)
	admin.DELETE("/users/:id", userHandler.Delete)
	admin.PUT("/content-submissions/:id/status", contentHandler.UpdateStatus)
	admin.DELETE("/content-submissions/:id", contentHandler.Delete)
	admin.PUT("/letter-submissions/:id/status", letterHandler.UpdateStatus)
	admin.DELETE("/letter-submissions/:id", letterHandler.Delete)
	admin.GET("/medinfo-pj/queue", queueHandler.List)
	admin.POST("/medinfo-pj/queue", queueHandler.Create)
	admin.PUT("/medinfo-pj/queue/reorder", queueHandler.Reorder)
	admin.DELETE("/medinfo-pj/queue/:id", queueHandler.Delete)
	admin.POST("/letter-templates", templateHandler.Create)
	admin.GET("/letter-templates", templateHandler.List)
	admin.GET("/letter-templates/:id", templateHandler.Get)
	admin.PUT("/letter-templates/:id", templateHandler.Update)
	admin.DELETE("/letter-templates/:id", templateHandler.Delete)

	cron.StartDailyCron(db, waClient, cfg)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
