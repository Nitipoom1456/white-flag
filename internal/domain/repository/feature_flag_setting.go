package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/white-flag/internal/domain/entity"
	"gorm.io/gorm"
)

type FeatureFlagSettingRepository interface {
	FindByAppIDAndEnvironmentID(ctx context.Context, appID, environmentID uuid.UUID, offset, limit int) ([]entity.FeatureFlagSetting, int64, error)
}

type FeatureFlagSetting struct {
	DB *gorm.DB
}

func NewFeatureFlagSettingRepository(db *gorm.DB) FeatureFlagSettingRepository {
	return &FeatureFlagSetting{DB: db}
}

func (f *FeatureFlagSetting) FindByAppIDAndEnvironmentID(ctx context.Context, appID, environmentID uuid.UUID, offset, limit int) ([]entity.FeatureFlagSetting, int64, error) {
	base := f.DB.WithContext(ctx).
		Model(&entity.FeatureFlagSetting{}).
		Joins("FeatureFlag").
		Where(`"FeatureFlag".app_id = ?`, appID).
		Where("feature_flag_setting.environment_id = ?", environmentID)

	var settings []entity.FeatureFlagSetting
	if err := base.Session(&gorm.Session{}).
		Order(`"FeatureFlag".created_at DESC`).
		Offset(offset).Limit(limit).
		Find(&settings).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := base.Session(&gorm.Session{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return settings, total, nil
}
