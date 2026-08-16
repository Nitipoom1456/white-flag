package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/white-flag/internal/bootstrap"
	"github.com/white-flag/internal/infrastructure/logger"
)

func main() {
	application, err := bootstrap.NewApplication()
	if err != nil {
		logger.Error(context.Background(), "Failed to create application", err)
	}

	err = application.Run()
	if err != nil {
		logger.Error(context.Background(), "Failed to run application", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info(context.Background(), "Shutting down .....")
}
