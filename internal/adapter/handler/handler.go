package handler

import (
	"github.com/white-flag/internal/infrastructure/config"
	"github.com/white-flag/internal/usecase/app"
	"github.com/white-flag/internal/usecase/appversion"
	"github.com/white-flag/internal/usecase/environment"
	featureflag "github.com/white-flag/internal/usecase/feature_flag"
	featureflagsetting "github.com/white-flag/internal/usecase/feature_flag_setting"
)

type Handler struct {
	Config                    *config.Config
	AppUsecase                app.AppUsecase
	AppVersionUsecase         appversion.AppVersionUsecase
	EnvironmentUsecase        environment.EnvironmentUsecase
	FeatureFlagUsecase        featureflag.FeatureFlagUsecase
	FeatureFlagSettingUsecase featureflagsetting.FeatureFlagSettingUsecase
}

func NewHandler(
	config *config.Config,
	appUsecase app.AppUsecase,
	appVersionUsecase appversion.AppVersionUsecase,
	environmentUsecase environment.EnvironmentUsecase,
	featureFlagUsecase featureflag.FeatureFlagUsecase,
	featureFlagSettingUsecase featureflagsetting.FeatureFlagSettingUsecase,
) *Handler {
	return &Handler{
		Config:                    config,
		AppUsecase:                appUsecase,
		AppVersionUsecase:         appVersionUsecase,
		EnvironmentUsecase:        environmentUsecase,
		FeatureFlagUsecase:        featureFlagUsecase,
		FeatureFlagSettingUsecase: featureFlagSettingUsecase,
	}
}
