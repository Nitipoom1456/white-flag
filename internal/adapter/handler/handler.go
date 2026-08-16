package handler

import (
	"github.com/white-flag/internal/infrastructure/config"
	"github.com/white-flag/internal/usecase/app"
	"github.com/white-flag/internal/usecase/appversion"
)

type Handler struct {
	Config            *config.Config
	AppUsecase        app.AppUsecase
	AppVersionUsecase appversion.AppVersionUsecase
}

func NewHandler(
	config *config.Config,
	appUsecase app.AppUsecase,
	appVersionUsecase appversion.AppVersionUsecase,
) *Handler {
	return &Handler{
		Config:            config,
		AppUsecase:        appUsecase,
		AppVersionUsecase: appVersionUsecase,
	}
}
