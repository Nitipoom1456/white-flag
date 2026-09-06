package app

import (
	"context"

	"github.com/white-flag/internal/infrastructure/logger"
	"github.com/white-flag/internal/usecase"
)

type App struct {
	ID        string
	Name      string
	CreatedAt string
}

func (u *Usecase) List(ctx context.Context, page, pageSize int) ([]App, int, *usecase.UsecaseError) {
	offset := (page - 1) * pageSize
	limit := pageSize
	apps, total, err := u.AppRepository.List(ctx, offset, limit)
	if err != nil {
		logger.Error(ctx, "Failed to get list apps", err)
		return nil, 0, &usecase.UsecaseError{
			Message: usecase.MSG_SERVER_ERROR,
			Code:    usecase.CODE_SERVER_ERROR,
			Err:     err,
		}
	}

	result := make([]App, 0, len(apps))
	for _, app := range apps {
		result = append(result, App{
			ID:        app.ID.String(),
			Name:      app.Name,
			CreatedAt: app.CreatedAt.String(),
		})
	}

	return result, int(total), nil
}
