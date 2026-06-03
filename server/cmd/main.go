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
	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		if err := database.Migrate(); err != nil {
			log.Fatal(err)
		}
		return
	}

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
	contentRepo := content.NewRepository(db)
	letterRepo := letter.NewRepository(db)
	queueRepo := medinfo_pj.NewRepository(db)
	templateRepo := letter_template.NewRepository(db)

	router.GET("/ping", func(c *gin.Context) { response.OK(c, "pong", nil) })

	api := router.Group("/api")
	{
		auth.RegisterRoutes(api, userRepo, cfg.JWTSecret)
		docs.RegisterRoutes(api, cfg.DocsDir)
		user.RegisterRoutes(api, userRepo, cfg.JWTSecret)
		content.RegisterRoutes(api, contentRepo, userRepo, waClient, cfg.JWTSecret)
		letter.RegisterRoutes(api, letterRepo, userRepo, waClient, cfg.JWTSecret)
		medinfo_pj.RegisterRoutes(api, queueRepo, cfg.JWTSecret)
		letter_template.RegisterRoutes(api, templateRepo, cfg.JWTSecret)
	}

	cron.StartDailyCron(db, waClient, cfg)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
