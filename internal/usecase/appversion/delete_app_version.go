package appversion

import (
	"context"

	"github.com/google/uuid"
	"github.com/white-flag/internal/infrastructure/logger"
	"github.com/white-flag/internal/usecase"
)

func (u *Usecase) Delete(ctx context.Context, id string) *usecase.UsecaseError {
	idUUID := uuid.MustParse(id)
	err := u.AppVersionRepository.Delete(ctx, idUUID)
	if err != nil {
		logger.Error(ctx, "Failed to delete app version", err)
		return &usecase.UsecaseError{
			Err:     err,
			Message: usecase.MSG_SERVER_ERROR,
			Code:    usecase.CODE_SERVER_ERROR,
		}
	}
	return nil
}
