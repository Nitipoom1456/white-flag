package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/white-flag/internal/domain/entity"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type FeatureFlagRepository interface {
	Create(ctx context.Context, flag CreateFeatureFlagParams) error
}

type FeatureFlag struct {
	DB *gorm.DB
}

func NewFeatureFlagRepository(db *gorm.DB) FeatureFlagRepository {
	return &FeatureFlag{DB: db}
}

type CreateFeatureFlagParams struct {
	AppID         uuid.UUID
	Name          string
	Description   *string
	Active        bool
	Tags          []string
	MinAppVersion *uuid.UUID
	MaxAppVersion *uuid.UUID
}

func (f *FeatureFlag) Create(ctx context.Context, flag CreateFeatureFlagParams) error {
	return f.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		featureFlag := entity.FeatureFlag{
			AppID:       flag.AppID,
			Name:        flag.Name,
			Description: flag.Description,
			Tags:        datatypes.JSONSlice[string](flag.Tags),
		}

		if err := tx.Create(&featureFlag).Error; err != nil {
			return err
		}

		var envIDs []uuid.UUID
		if err := tx.Model(&entity.Environment{}).
			Where("app_id = ?", flag.AppID).
			Pluck("id", &envIDs).Error; err != nil {
			return err
		}
		if len(envIDs) == 0 {
			return gorm.ErrRecordNotFound
		}

		settings := make([]entity.FeatureFlagSetting, 0, len(envIDs))
		for _, envID := range envIDs {
			s := entity.FeatureFlagSetting{
				FeatureFlagID:   featureFlag.ID,
				EnvironmentID:   envID,
				Active:          flag.Active,
				MinAppVersionID: flag.MinAppVersion,
				MaxAppVersionID: flag.MaxAppVersion,
			}
			settings = append(settings, s)
		}

		return tx.Create(&settings).Error
	})
}
