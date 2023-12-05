package middlewares

import (
	config "productservice/config"
	"productservice/internal/infrastructure"

	"go.uber.org/zap"
)

type GinMiddleware struct {
	logger *zap.Logger
	config *config.Config
	db     infrastructure.Database
}

func NewMiddleware(config *config.Config, db infrastructure.Database, logger *zap.Logger) *GinMiddleware {
	return &GinMiddleware{
		logger: logger,
		config: config,
		db:     db,
	}
}
