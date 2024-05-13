package service

import (
	"productservice/config"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/repository"

	"go.uber.org/zap"
)

type productModelService struct {
	db            infrastructure.Database
	dbTransaction infrastructure.DatabaseTransaction
	ufw           *repository.UnitOfWork
	config        *config.Config
	logger        *zap.Logger
}

func NewProductModel(
	db infrastructure.Database,
	dbTransaction infrastructure.DatabaseTransaction,
	ufw *repository.UnitOfWork,
	config *config.Config,
	logger *zap.Logger,
) domain.ProductModelService {
	return &productModelService{
		db:            db,
		dbTransaction: dbTransaction,
		ufw:           ufw,
		config:        config,
		logger:        logger,
	}
}
