package app

import (
	"context"

	"github.com/white-flag/internal/domain/repository"
	"github.com/white-flag/internal/usecase"
)

type AppUsecase interface {
	Create(ctx context.Context, appName string) *usecase.UsecaseError
	List(ctx context.Context, page, pageSize int) ([]App, int, *usecase.UsecaseError)
}

type Usecase struct {
	AppRepository repository.AppRepository
}

func NewAppUsecase(appRepository repository.AppRepository) AppUsecase {
	return &Usecase{AppRepository: appRepository}
}
