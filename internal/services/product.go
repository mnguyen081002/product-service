package service

import (
	"context"
	"productservice/config"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
	"productservice/internal/repository"

	"go.uber.org/zap"
)

type productService struct {
	db            infrastructure.Database
	dbTransaction infrastructure.DatabaseTransaction
	ufw           *repository.UnitOfWork
	config        *config.Config
	logger        *zap.Logger
}

func NewProductService(
	db infrastructure.Database,
	dbTransaction infrastructure.DatabaseTransaction,
	ufw *repository.UnitOfWork,
	config *config.Config,
	logger *zap.Logger,
) domain.ProductService {
	return &productService{
		db:            db,
		dbTransaction: dbTransaction,
		ufw:           ufw,
		config:        config,
		logger:        logger,
	}
}

func (p *productService) GetProductByID(ctx context.Context, id string) (product *models.Product, err error) {
	return p.ufw.ProductRepository.GetByIdJoinCategory(&p.db, ctx, id)
}
