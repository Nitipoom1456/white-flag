package featureflag

import (
	"context"

	"github.com/white-flag/internal/domain/repository"
	"github.com/white-flag/internal/usecase"
)

type FeatureFlagUsecase interface {
	Create(ctx context.Context, params CreateFeatureFlagParams) *usecase.UsecaseError
}

type Usecase struct {
	FeatureFlagRepository repository.FeatureFlagRepository
}

func NewFeatureFlagUsecase(
	featureFlagRepository repository.FeatureFlagRepository,
) FeatureFlagUsecase {
	return &Usecase{
		FeatureFlagRepository: featureFlagRepository,
	}
}
