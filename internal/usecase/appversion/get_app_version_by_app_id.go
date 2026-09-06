package appversion

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/white-flag/internal/infrastructure/logger"
	"github.com/white-flag/internal/usecase"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppVersion struct {
	ID      string
	Version string
	AppID   string
}

func (u *Usecase) GetByAppID(ctx context.Context, appID string, page, pageSize int) ([]AppVersion, int, *usecase.UsecaseError) {
	offset := (page - 1) * pageSize
	limit := pageSize
	appUUID := uuid.MustParse(appID)
	appVersions, total, err := u.AppVersionRepository.FindByAppID(ctx, appUUID, offset, limit)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error(ctx, "App version not found", err, zap.String("app_id", appID))
			return []AppVersion{}, 0, nil
		}
		logger.Error(ctx, "Failed to get app versions by app ID", err, zap.String("app_id", appID))
		return nil, 0, &usecase.UsecaseError{
			Message: usecase.MSG_APP_VERSION_DUPLICATED_KEY,
			Code:    usecase.CODE_SERVER_ERROR,
			Err:     err,
		}
	}

	result := make([]AppVersion, len(appVersions))
	for i, v := range appVersions {
		result[i] = AppVersion{
			ID:      v.ID.String(),
			Version: v.Version,
			AppID:   v.AppID.String(),
		}
	}

	return result, int(total), nil
}
