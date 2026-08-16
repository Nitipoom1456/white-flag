package server

import (
	"github.com/gin-gonic/gin"
	"github.com/white-flag/internal/adapter"
	"github.com/white-flag/internal/adapter/handler"
	"github.com/white-flag/internal/adapter/middleware"
	"github.com/white-flag/internal/infrastructure/config"
)

type Server struct {
	Config *config.Config
	Router *gin.Engine
}

func NewServer(
	config *config.Config,
	handler *handler.Handler,
) *Server {
	engine := gin.New()
	engine.Use(middleware.Roundtrip(), gin.Recovery())

	router := adapter.NewRouter(engine)
	router.SetupRoutes(handler)
	return &Server{
		Config: config,
		Router: engine,
	}
}

func (s *Server) Start() error {
	return s.Router.Run()
}
