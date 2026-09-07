package featureflagsetting

import (
	"context"

	"github.com/google/uuid"
	"github.com/white-flag/internal/infrastructure/logger"
	"github.com/white-flag/internal/usecase"
)

type FeatureFlagDashboard struct {
	ID              string
	FeatureFlagID   string
	EnvironmentID   string
	Name            string
	Description     *string
	Active          bool
	Tags            []string
	MinAppVersionID *string
	MaxAppVersionID *string
	CreatedAt       string
	UpdatedAt       string
}

func (u *Usecase) GetFeatureFlagSettingsByAppIDAndEnvironmentID(ctx context.Context, appID, environmentID string, page, pageSize int) ([]FeatureFlagDashboard, int, *usecase.UsecaseError) {
	appUUID := uuid.MustParse(appID)
	envUUID := uuid.MustParse(environmentID)
	offset := (page - 1) * pageSize
	featureFlagSettings, total, err := u.FeatureFlagSettingRepository.FindByAppIDAndEnvironmentID(ctx, appUUID, envUUID, offset, pageSize)
	if err != nil {
		logger.Error(ctx, "Failed to get feature flag settings", err,
			logger.F("app_id", appID),
			logger.F("environment_id", environmentID),
		)
		return nil, 0, &usecase.UsecaseError{
			Err:     err,
			Code:    usecase.CODE_SERVER_ERROR,
			Message: usecase.MSG_SERVER_ERROR,
		}
	}

	result := make([]FeatureFlagDashboard, len(featureFlagSettings))
	for i, setting := range featureFlagSettings {
		var maxAppVersionID, minAppVersionID *string
		if setting.MaxAppVersionID != nil {
			temps := setting.MaxAppVersionID.String()
			maxAppVersionID = &temps
		}
		if setting.MinAppVersionID != nil {
			temp := setting.MinAppVersionID.String()
			minAppVersionID = &temp
		}
		result[i] = FeatureFlagDashboard{
			ID:              setting.ID.String(),
			FeatureFlagID:   setting.FeatureFlagID.String(),
			EnvironmentID:   setting.EnvironmentID.String(),
			Name:            setting.FeatureFlag.Name,
			Description:     setting.FeatureFlag.Description,
			Active:          setting.Active,
			Tags:            setting.FeatureFlag.Tags,
			MinAppVersionID: minAppVersionID,
			MaxAppVersionID: maxAppVersionID,
			CreatedAt:       setting.CreatedAt.String(),
			UpdatedAt:       setting.UpdatedAt.String(),
		}
	}
	return result, int(total), nil
}
