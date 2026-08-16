package appversion

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/white-flag/internal/domain/entity"
	"github.com/white-flag/internal/infrastructure/logger"
	"github.com/white-flag/internal/usecase"
	"gorm.io/gorm"
)

func (u *Usecase) Create(ctx context.Context, appID, version string) *usecase.UsecaseError {
	err := u.AppVersionRepository.Create(ctx, entity.AppVersion{
		AppID:   uuid.MustParse(appID),
		Version: version,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			logger.Error(ctx, "App version already exists", err, logger.F("version", version))
			return &usecase.UsecaseError{
				Message: usecase.MSG_APP_DUPLICATED_KEY,
				Code:    usecase.CODE_DUPLICATED_KEY,
				Err:     err,
			}
		}
		logger.Error(ctx, "Failed to create app version", err, logger.F("version", version))
		return &usecase.UsecaseError{
			Message: usecase.MSG_SERVER_ERROR,
			Code:    usecase.CODE_SERVER_ERROR,
			Err:     err,
		}
	}
	return nil
}
