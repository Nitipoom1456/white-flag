package appversion

import (
	"context"

	"github.com/white-flag/internal/domain/repository"
	"github.com/white-flag/internal/usecase"
)

type AppVersionUsecase interface {
	Create(ctx context.Context, appID, version string) *usecase.UsecaseError
	GetByAppID(ctx context.Context, appID string, page, pageSize int) ([]AppVersion, int, *usecase.UsecaseError)
}

type Usecase struct {
	AppVersionRepository repository.AppVersionRepository
}

func NewAppVersionUsecase(appVersionRepository repository.AppVersionRepository) AppVersionUsecase {
	return &Usecase{AppVersionRepository: appVersionRepository}
}
