package featureflagsetting

import (
	"context"

	"github.com/white-flag/internal/domain/repository"
	"github.com/white-flag/internal/usecase"
)

type FeatureFlagSettingUsecase interface {
	GetFeatureFlagSettingsByAppIDAndEnvironmentID(ctx context.Context, appID, environmentID string, page, pageSize int) ([]FeatureFlagDashboard, int, *usecase.UsecaseError)
}

type Usecase struct {
	FeatureFlagSettingRepository repository.FeatureFlagSettingRepository
}

func NewFeatureFlagSettingUsecase(
	featureFlagSettingRepository repository.FeatureFlagSettingRepository,
) FeatureFlagSettingUsecase {
	return &Usecase{
		FeatureFlagSettingRepository: featureFlagSettingRepository,
	}
}
