package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	AppEnv     string
	Port       string
	TZ         string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBMaxOpen  int
	DBMaxIdle  int
	JWTSecret  string
	// SeedUserPassword is the single shared dev-login password for every row
	// seeded by database/seeders/001_users.sql. Real plaintext passwords never
	// live in that committed SQL file (or in git at all) — only this env var,
	// which stays in the gitignored .env. The seeder hashes it at run time.
	SeedUserPassword string
	WAURL            string
	WAAPIKey         string
	WAGroupJID1      string
	WAGroupJID2      string
	DocsDir          string
}

func Load() *Config {
	appEnv := strings.ToLower(strings.TrimSpace(env("APP_ENV", "dev")))
	isDev := appEnv == "dev" || appEnv == "development"
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		if !isDev {
			log.Fatalf("JWT_SECRET must be set when APP_ENV=%q (non-dev)", appEnv)
		}
		jwtSecret = "dev_secret_change_me"
	}
	return &Config{
		AppEnv:           env("APP_ENV", "dev"),
		Port:             env("PORT", "8080"),
		TZ:               env("TZ", "Asia/Jakarta"),
		DBHost:           env("DB_HOST", "127.0.0.1"),
		DBPort:           env("DB_PORT", "3306"),
		DBUser:           env("DB_USER", "bemunair"),
		DBPassword:       env("DB_PASSWORD", "bemunair_password"),
		DBName:           env("DB_NAME", "bemunair_db"),
		DBMaxOpen:        envInt("DB_MAX_OPEN", 20),
		DBMaxIdle:        envInt("DB_MAX_IDLE", 10),
		JWTSecret:        jwtSecret,
		SeedUserPassword: os.Getenv("SEED_USER_PASSWORD"),
		WAURL:            env("WA_ENGINE_URL", "http://localhost:3001"),
		WAAPIKey:         env("WA_ENGINE_API_KEY", ""),
		WAGroupJID1:      env("WA_GROUP_JID_1", ""),
		WAGroupJID2:      env("WA_GROUP_JID_2", ""),
		DocsDir:          env("DOCS_DIR", "../docs/api"),
	}
}

func (c *Config) DSN() string {
	cfg := mysql.NewConfig()
	cfg.User = c.DBUser
	cfg.Passwd = c.DBPassword
	cfg.Net = "tcp"
	cfg.Addr = c.DBHost + ":" + c.DBPort
	cfg.DBName = c.DBName
	cfg.ParseTime = true
	cfg.Loc = c.tzLocation()
	cfg.Params = map[string]string{"charset": "utf8mb4"}
	return cfg.FormatDSN()
}

func (c *Config) tzLocation() *time.Location {
	loc, err := time.LoadLocation(c.TZ)
	if err != nil {
		loc = time.UTC
	}
	return loc
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func envInt(key string, fallback int) int {
	value, err := strconv.Atoi(env(key, ""))
	if err != nil {
		return fallback
	}
	return value
}
