package repository

import (
	"gorm.io/gorm"
)

type FeatureFlagSettingRepository interface {
}

type FeatureFlagSetting struct {
	DB *gorm.DB
}

func NewFeatureFlagSettingRepository(db *gorm.DB) FeatureFlagSettingRepository {
	return &FeatureFlagSetting{DB: db}
}
