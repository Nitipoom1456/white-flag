package appversion

import (
	"context"

	"github.com/white-flag/internal/domain/repository"
	"github.com/white-flag/internal/usecase"
)

type AppVersionUsecase interface {
	Create(ctx context.Context, appID, version string) *usecase.UsecaseError
}

type Usecase struct {
	AppVersionRepository repository.AppVersionRepository
}

func NewAppVersionUsecase(appVersionRepository repository.AppVersionRepository) AppVersionUsecase {
	return &Usecase{AppVersionRepository: appVersionRepository}
}
