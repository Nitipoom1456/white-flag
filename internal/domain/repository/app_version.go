package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/white-flag/internal/domain/entity"
	"gorm.io/gorm"
)

type AppVersionRepository interface {
	Create(ctx context.Context, appVersion entity.AppVersion) error
	FindByAppID(ctx context.Context, appID uuid.UUID, offset, limit int) ([]entity.AppVersion, int64, error)
	Delete(ctx context.Context, id uuid.UUID) error
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

func (a *AppVersion) FindByAppID(ctx context.Context, appID uuid.UUID, offset, limit int) ([]entity.AppVersion, int64, error) {
	var appVersions []entity.AppVersion
	if err := a.DB.WithContext(ctx).Where("app_id = ?", appID).Offset(offset).Limit(limit).Find(&appVersions).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := a.DB.WithContext(ctx).Model(&entity.AppVersion{}).Where("app_id = ?", appID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return appVersions, total, nil
}

func (a *AppVersion) Delete(ctx context.Context, id uuid.UUID) error {
	return a.DB.WithContext(ctx).Delete(&entity.AppVersion{ID: id}).Error
}
