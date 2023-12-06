package lib

import (
	"context"
	"fmt"
	config "productservice/config"
	"productservice/internal/api/middlewares"
	"productservice/internal/constants"
	"productservice/internal/infrastructure"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"go.uber.org/fx"
)

type Handler struct {
	*gin.RouterGroup
}

func NewServerGroup(instance *gin.Engine) *Handler {
	return &Handler{
		instance.Group("/v1/api"),
	}
}

func NewServer(lifecycle fx.Lifecycle, zap *zap.Logger, config *config.Config, db infrastructure.Database, middlewares *middlewares.GinMiddleware) *gin.Engine {
	switch config.Server.Env {
	case constants.Dev, constants.Local:
		gin.SetMode(gin.DebugMode)
	case constants.Prod:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	instance := gin.New()

	//instance.Use(gozap.RecoveryWithZap(zap, true))

	instance.Use(middlewares.JSONMiddleware)
	if config.Logger.Enabled {
		instance.Use(middlewares.Logger)
	}
	instance.Use(middlewares.ErrorHandler)

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			zap.Info("Starting HTTP server")

			SeedRoutes(instance, &db)
			go func() {
				addr := fmt.Sprint(config.Server.Host, ":", config.Server.Port)
				if err := instance.Run(addr); err != nil {
					zap.Fatal(fmt.Sprint("HTTP server failed to start %w", err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			zap.Info("Stopping HTTP server")
			return nil
		},
	})

	return instance
}

func SeedRoutes(engine *gin.Engine, db *infrastructure.Database) error {
	return nil
}
