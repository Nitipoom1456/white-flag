package app

import (
	"context"
	"errors"

	"github.com/white-flag/internal/domain/entity"
	"github.com/white-flag/internal/infrastructure/logger"
	"github.com/white-flag/internal/usecase"
	"gorm.io/gorm"
)

func (u *Usecase) Create(ctx context.Context, appName string) *usecase.UsecaseError {
	err := u.AppRepository.Create(ctx, entity.App{
		Name: appName,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			logger.Error(ctx, "App name already exists", err, logger.F("appName", appName))
			return &usecase.UsecaseError{
				Message: usecase.MSG_APP_DUPLICATED_KEY,
				Code:    usecase.CODE_DUPLICATED_KEY,
				Err:     err,
			}
		}
		logger.Error(ctx, "Failed to create app", err, logger.F("appName", appName))
		return &usecase.UsecaseError{
			Message: usecase.MSG_SERVER_ERROR,
			Code:    usecase.CODE_SERVER_ERROR,
			Err:     err,
		}
	}
	return nil
}
