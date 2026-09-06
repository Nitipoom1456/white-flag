package environment

import (
	"context"

	"github.com/google/uuid"
	"github.com/white-flag/internal/domain/entity"
	"github.com/white-flag/internal/infrastructure/logger"
	"github.com/white-flag/internal/usecase"
	"go.uber.org/zap"
)

func (u *Usecase) Create(ctx context.Context, appID, name string) *usecase.UsecaseError {
	appUUID := uuid.MustParse(appID)
	err := u.EnvironmentRepository.Create(ctx, entity.Environment{
		Name:  name,
		AppID: appUUID,
	})
	if err != nil {
		logger.Error(ctx, "Failed to create environment", err, zap.String("app_id", appID))
		return &usecase.UsecaseError{
			Err:     err,
			Code:    usecase.CODE_SERVER_ERROR,
			Message: usecase.MSG_SERVER_ERROR,
		}
	}
	return nil
}
