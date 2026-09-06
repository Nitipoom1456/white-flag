package featureflag

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/white-flag/internal/domain/repository"
	"github.com/white-flag/internal/infrastructure/logger"
	"github.com/white-flag/internal/usecase"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CreateFeatureFlagParams struct {
	AppID         string
	Name          string
	Description   *string
	Active        bool
	Tags          []string
	MinAppVersion *string
	MaxAppVersion *string
}

func (u *Usecase) Create(ctx context.Context, params CreateFeatureFlagParams) *usecase.UsecaseError {
	appUUID := uuid.MustParse(params.AppID)
	var minVersionUUID, maxVersionUUID *uuid.UUID
	if params.MinAppVersion != nil {
		v := uuid.MustParse(*params.MinAppVersion)
		minVersionUUID = &v
	}
	if params.MaxAppVersion != nil {
		v := uuid.MustParse(*params.MaxAppVersion)
		maxVersionUUID = &v
	}
	err := u.FeatureFlagRepository.Create(ctx, repository.CreateFeatureFlagParams{
		AppID:         appUUID,
		Name:          params.Name,
		Description:   params.Description,
		Active:        params.Active,
		Tags:          params.Tags,
		MinAppVersion: minVersionUUID,
		MaxAppVersion: maxVersionUUID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error(ctx, "Environment not found", err, zap.String("app_id", params.AppID))
			return &usecase.UsecaseError{
				Err:     err,
				Code:    "environment_not_found",
				Message: "Environment not found",
			}
		}
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			logger.Error(ctx, "Duplicate key", err, zap.String("app_id", params.AppID))
			return &usecase.UsecaseError{
				Err:     err,
				Code:    "duplicate_key",
				Message: "Duplicate key",
			}
		}
		logger.Error(ctx, "Failed to create feature flag", err, zap.String("app_id", params.AppID))
		return &usecase.UsecaseError{
			Err:     err,
			Code:    "failed_to_create_feature_flag",
			Message: "Failed to create feature flag",
		}
	}
	return nil
}
