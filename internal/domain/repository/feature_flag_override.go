package repository

import (
	"gorm.io/gorm"
)

type FeatureFlagOverrideRepository interface {
}

type FeatureFlagOverride struct {
	DB *gorm.DB
}

func NewFeatureFlagOverrideRepository(db *gorm.DB) FeatureFlagOverrideRepository {
	return &FeatureFlagOverride{DB: db}
}
