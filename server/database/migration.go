package database

import (
	"time"

	"bemunair2026/server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	level := logger.Info
	if cfg.AppEnv == "prod" || cfg.AppEnv == "production" {
		level = logger.Silent
	}

	db, err := gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(level),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(cfg.DBMaxOpen)
	sqlDB.SetMaxIdleConns(cfg.DBMaxIdle)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func Migrate() error {
	cfg := config.Load()
	db, err := Connect(cfg)
	if err != nil {
		return err
	}

	return NewMigrationManager(db).Run()
}
