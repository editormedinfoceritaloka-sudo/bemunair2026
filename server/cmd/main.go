package main

import (
	"log"
	"os"

	"bemunair2026/server/config"
	"bemunair2026/server/database"
	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/article"
	articleRepository "bemunair2026/server/modules/article/repository"
	"bemunair2026/server/modules/auth"
	content "bemunair2026/server/modules/content_submission"
	contentRepository "bemunair2026/server/modules/content_submission/repository"
	"bemunair2026/server/modules/cron"
	"bemunair2026/server/modules/docs"
	letter "bemunair2026/server/modules/letter_submission"
	letterRepository "bemunair2026/server/modules/letter_submission/repository"
	"bemunair2026/server/modules/letter_template"
	templateRepository "bemunair2026/server/modules/letter_template/repository"
	"bemunair2026/server/modules/medinfo_pj"
	medinfoRepository "bemunair2026/server/modules/medinfo_pj/repository"
	"bemunair2026/server/modules/submission_support"
	"bemunair2026/server/modules/user"
	userRepository "bemunair2026/server/modules/user/repository"
	"bemunair2026/server/pkg"
	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

func runDatabaseCommand(command string) (bool, error) {
	switch command {
	case "migrate":
		return true, database.Migrate()
	case "seed":
		return true, database.Seed()
	case "setup":
		if err := database.Migrate(); err != nil {
			return true, err
		}
		return true, database.Seed()
	default:
		return false, nil
	}
}

func main() {
	if len(os.Args) > 1 {
		handled, err := runDatabaseCommand(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		if handled {
			return
		}
	}

	cfg := config.Load()
	_ = os.Setenv("TZ", cfg.TZ)
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	router := gin.New()
	router.Use(middlewares.RequestLogger(), middlewares.Recovery(), middlewares.CORS())

	userRepo := userRepository.NewUserRepository(db)
	waClient := pkg.NewWAClient(cfg.WAURL, cfg.WAAPIKey)
	contentRepo := contentRepository.NewContentSubmissionRepository(db)
	letterRepo := letterRepository.NewLetterSubmissionRepository(db)
	queueRepo := medinfoRepository.NewMedinfoPJRepository(db)
	templateRepo := templateRepository.NewLetterTemplateRepository(db)
	articleRepo := articleRepository.NewArticleRepository(db)

	router.GET("/ping", func(c *gin.Context) { response.OK(c, "pong", nil) })

	v1 := router.Group("/api/v1")
	{
		auth.RegisterRoutes(v1, userRepo, cfg.JWTSecret)
		docs.RegisterRoutes(v1, cfg.DocsDir)
		user.RegisterRoutes(v1, userRepo, cfg.JWTSecret)
		content.RegisterRoutes(v1, contentRepo, userRepo, waClient, cfg.JWTSecret)
		letter.RegisterRoutes(v1, letterRepo, userRepo, waClient, cfg.JWTSecret)
		medinfo_pj.RegisterRoutes(v1, queueRepo, cfg.JWTSecret)
		letter_template.RegisterRoutes(v1, templateRepo, cfg.JWTSecret)
		article.RegisterRoutes(v1, articleRepo, cfg.JWTSecret)
		submission_support.RegisterRoutes(v1, db, cfg.JWTSecret)
	}

	cron.StartDailyCron(db, waClient, cfg)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
