package environment

import (
	"context"

	"github.com/white-flag/internal/domain/repository"
	"github.com/white-flag/internal/usecase"
)

type EnvironmentUsecase interface {
	Create(ctx context.Context, appID, name string) *usecase.UsecaseError
	GetByAppID(ctx context.Context, appID string) ([]Environment, *usecase.UsecaseError)
	Delete(ctx context.Context, id string) *usecase.UsecaseError
}

type Usecase struct {
	EnvironmentRepository repository.EnvironmentRepository
}

func NewEnvironmentUsecase(environmentRepository repository.EnvironmentRepository) EnvironmentUsecase {
	return &Usecase{EnvironmentRepository: environmentRepository}
}
