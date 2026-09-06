package environment

import (
	"context"

	"github.com/google/uuid"
	"github.com/white-flag/internal/infrastructure/logger"
	"github.com/white-flag/internal/usecase"
	"go.uber.org/zap"
)

func (u *Usecase) Delete(ctx context.Context, id string) *usecase.UsecaseError {
	idUUID := uuid.MustParse(id)
	err := u.EnvironmentRepository.Delete(ctx, idUUID)
	if err != nil {
		logger.Error(ctx, "Failed to delete environment", err, zap.String("id", id))
		return &usecase.UsecaseError{
			Err:     err,
			Code:    usecase.CODE_SERVER_ERROR,
			Message: usecase.MSG_SERVER_ERROR,
		}
	}
	return nil
}
