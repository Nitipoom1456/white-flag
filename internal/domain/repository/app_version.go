package repository

import (
	"context"

	"github.com/white-flag/internal/domain/entity"
	"gorm.io/gorm"
)

type AppVersionRepository interface {
	Create(ctx context.Context, appVersion entity.AppVersion) error
}

type AppVersion struct {
	DB *gorm.DB
}

func NewAppVersionRepository(db *gorm.DB) AppVersionRepository {
	return &AppVersion{DB: db}
}

func (a *AppVersion) Create(ctx context.Context, appVersion entity.AppVersion) error {
	return a.DB.WithContext(ctx).Create(&appVersion).Error
}
