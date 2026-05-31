package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	AppEnv      string
	Port        string
	TZ          string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBMaxOpen   int
	DBMaxIdle   int
	JWTSecret   string
	WAURL       string
	WAAPIKey    string
	WAGroupJID1 string
	WAGroupJID2 string
	DocsDir     string
}

func Load() *Config {
	return &Config{
		AppEnv:      env("APP_ENV", "dev"),
		Port:        env("PORT", "8080"),
		TZ:          env("TZ", "Asia/Jakarta"),
		DBHost:      env("DB_HOST", "127.0.0.1"),
		DBPort:      env("DB_PORT", "3306"),
		DBUser:      env("DB_USER", "bemunair"),
		DBPassword:  env("DB_PASSWORD", "bemunair_password"),
		DBName:      env("DB_NAME", "bemunair_db"),
		DBMaxOpen:   envInt("DB_MAX_OPEN", 20),
		DBMaxIdle:   envInt("DB_MAX_IDLE", 10),
		JWTSecret:   env("JWT_SECRET", "dev_secret_change_me"),
		WAURL:       env("WA_ENGINE_URL", "http://localhost:3001"),
		WAAPIKey:    env("WA_ENGINE_API_KEY", ""),
		WAGroupJID1: env("WA_GROUP_JID_1", ""),
		WAGroupJID2: env("WA_GROUP_JID_2", ""),
		DocsDir:     env("DOCS_DIR", "../docs/api"),
	}
}

func (c *Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FJakarta",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
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
