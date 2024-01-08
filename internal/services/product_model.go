package service

import (
	"context"
	"productservice/config"
	"productservice/internal/api/request"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
	"productservice/internal/repository"
	"sync"

	"go.uber.org/zap"
)

type productModelService struct {
	db            infrastructure.Database
	dbTransaction infrastructure.DatabaseTransaction
	prroductModel domain.ProductModelRepository
	ufw           *repository.UnitOfWork
	config        *config.Config
	logger        *zap.Logger
	mu            sync.Mutex
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

func (a *productModelService) CreateProductModels(ctx context.Context, req request.CreateTierVariation) (productAttributes *models.ProductModel, err error) {
	_, err = a.ufw.ProductModelRepository.BulkCreate(&a.db, ctx, request.ToArrayProductModel(req.Variations, req.ProductID))
	if err != nil {
		return nil, err
	}

	return productAttributes, nil
}

func (a *productModelService) GetListByProductId(ctx context.Context, productID string) (res []*models.ProductModel, err error) {
	res, err = a.ufw.ProductModelRepository.GetListByProductId(&a.db, ctx, productID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
