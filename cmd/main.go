package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AssylzhanZharzhanov/go-shortener/internal/application"
	"github.com/AssylzhanZharzhanov/go-shortener/internal/helpers"
)

// @title Swagger Mercury
// @version 1.0
// @description Documentation for Mercury APIs.
// @termsOfService http://swagger.io/terms/

// @host localhost:8000
// @schemes http
// @BasePath /
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	var (
		ctx = context.Background()
	)

	cfg, err := helpers.LoadConfig()
	if err != nil {
		log.Fatal("can not load app config:", err)
	}

	app := application.NewApp(cfg)
	go func() {
		if err := app.Run(); err != nil {
			log.Fatalf("Error in starting application: %s", err.Error())
		}
	}()

	//graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := app.Shutdown(ctx); err != nil {
		log.Printf("error occured on application shutting down: %s", err.Error())
	}
}
