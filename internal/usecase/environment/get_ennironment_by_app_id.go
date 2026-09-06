package environment

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/white-flag/internal/infrastructure/logger"
	"github.com/white-flag/internal/usecase"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Environment struct {
	ID    string
	Name  string
	AppID string
}

func (u *Usecase) GetByAppID(ctx context.Context, appID string) ([]Environment, *usecase.UsecaseError) {
	appUUID := uuid.MustParse(appID)
	environments, err := u.EnvironmentRepository.FindByAppID(ctx, appUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error(ctx, "Environment not found", err, zap.String("app_id", appID))
			return []Environment{}, nil
		}
		logger.Error(ctx, "Failed to get environments by app ID", err, zap.String("app_id", appID))
		return nil, &usecase.UsecaseError{
			Err:     err,
			Code:    usecase.CODE_SERVER_ERROR,
			Message: usecase.MSG_SERVER_ERROR,
		}
	}

	result := make([]Environment, len(environments))
	for i, v := range environments {
		result[i] = Environment{
			ID:    v.ID.String(),
			Name:  v.Name,
			AppID: v.AppID.String(),
		}
	}
	return result, nil
}
