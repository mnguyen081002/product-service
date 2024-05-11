package service

import (
	"context"
	"productservice/config"
	"productservice/internal/api/request"
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

func (p *productService) GetProductList(ctx context.Context, input request.ClientListProductRequest) (res []*models.Product, total int64, err error) {
	return p.ufw.ClientProductRepository.List(&p.db, ctx, domain.ClientListProductQueryOptions{
		PageOptions: input.PageOptions,
		Price:       input.Price,
	})
}
