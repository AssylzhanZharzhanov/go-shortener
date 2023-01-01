package application

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/AssylzhanZharzhanov/go-shortener/internal/helpers"
	shortenerEndpoint "github.com/AssylzhanZharzhanov/go-shortener/internal/shortener/delivery/http"
	"github.com/AssylzhanZharzhanov/go-shortener/internal/shortener/repository"
	"github.com/AssylzhanZharzhanov/go-shortener/internal/shortener/service"
	"github.com/AssylzhanZharzhanov/go-shortener/pkg/database/postgres"
	"github.com/AssylzhanZharzhanov/go-shortener/pkg/database/redis"

	"github.com/gin-gonic/gin"
)

// App - application struct
type App struct {
	cfg helpers.AppConfig

	server *http.Server
}

// NewApp - creates a new application
func NewApp(cfg helpers.AppConfig) *App {
	return &App{
		cfg: cfg,
	}
}

// Run - starts application
func (a *App) Run() error {

	//database connections
	postgresDB, err := postgres.NewConnection(a.cfg.DSN)
	if err != nil {
		log.Fatalf("mongodb connection error: %s", err.Error())
	}

	redisClient := redis.NewRedisClient(a.cfg)

	//repository layer
	shortenerRepository := repository.NewShortenerRepository(postgresDB)
	shortenerRedisRepository := repository.NewRedisRepository(redisClient)

	//service layer
	shortenerService := service.NewShortenerService(shortenerRepository, shortenerRedisRepository)

	//handler layer
	router := gin.New()
	api := router.Group("/api")
	shortenerEndpoint.RegisterEndpoints(api, shortenerService)

	a.server = &http.Server{
		Addr:           ":" + a.cfg.Port,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return a.server.ListenAndServe()
}

// Shutdown - stops application
func (a *App) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}
