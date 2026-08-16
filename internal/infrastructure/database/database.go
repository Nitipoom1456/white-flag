package database

import (
	"github.com/white-flag/internal/infrastructure/config"
	"github.com/white-flag/internal/infrastructure/logger"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(config *config.Config, logger *logger.Logger) *Database {
	db, err := gorm.Open(postgres.Open(config.DBURL), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		logger.Zap.Fatal("failed to connect to database", zap.Error(err))
	}
	return &Database{DB: db}
}
