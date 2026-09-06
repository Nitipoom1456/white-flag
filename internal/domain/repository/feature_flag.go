package repository

import (
	"gorm.io/gorm"
)

type FeatureFlagRepository interface {
}

type FeatureFlag struct {
	DB *gorm.DB
}

func NewFeatureFlagRepository(db *gorm.DB) FeatureFlagRepository {
	return &FeatureFlag{DB: db}
}
