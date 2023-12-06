package bootstrap

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"productservice/config"
	controller "productservice/internal/api/controllers"
	"productservice/internal/api/middlewares"
	"productservice/internal/api/route"
	"productservice/internal/infrastructure"
	"productservice/internal/lib"
	"productservice/internal/messaging"
	"productservice/internal/repository"
	service "productservice/internal/services"
	"productservice/internal/utils"
)

func inject() fx.Option {
	return fx.Options(
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		//fx.NopLogger,
		fx.Provide(
			config.NewConfig("/config/config.yaml"),
			utils.NewTimeoutContext,
		),
		route.Module,
		lib.Module,
		repository.Module,
		service.Module,
		controller.Module,
		middlewares.Module,
		infrastructure.Module,
		messaging.Module,
	)
}

func Run() {
	fx.New(inject()).Run()
}
