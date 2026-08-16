package bootstrap

import (
	"github.com/white-flag/internal/adapter/handler"
	"github.com/white-flag/internal/domain/repository"
	"github.com/white-flag/internal/infrastructure/config"
	"github.com/white-flag/internal/infrastructure/database"
	"github.com/white-flag/internal/infrastructure/logger"
	"github.com/white-flag/internal/infrastructure/server"
	"github.com/white-flag/internal/usecase/app"
	"github.com/white-flag/internal/usecase/appversion"
)

type Application struct {
	Config *config.Config
	Logger *logger.Logger
	DB     *database.Database
	Server *server.Server
}

func NewApplication() (*Application, error) {
	config, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	logger.SetDefault(log)
	db := database.NewDatabase(config, log)

	appRepository := repository.NewAppRepository(db.DB)
	appVersionRepository := repository.NewAppVersionRepository(db.DB)

	appUseCase := app.NewAppUsecase(appRepository)
	appVersionUseCase := appversion.NewAppVersionUsecase(appVersionRepository)

	Handler := handler.NewHandler(
		config,
		appUseCase,
		appVersionUseCase,
	)

	server := server.NewServer(config, Handler)
	return &Application{
		Config: config,
		Logger: log,
		DB:     db,
		Server: server,
	}, nil
}

func (a *Application) Run() error {
	return a.Server.Start()
}
