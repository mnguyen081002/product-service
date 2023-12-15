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

	uuid "github.com/satori/go.uuid"
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

func (a *productModelService) InitProductModel(ctx context.Context, req request.TierVariationCreate) (productAttributes *models.ProductModels, err error) {

	for _, p := range req.ToArrayProductModel() {
		a.ufw.ProductModelRepository.Create(&a.db, ctx, &models.ProductModels{
			ProductID: uuid.FromStringOrNil(req.ProductID),
			Models:    p,
		})
	}

	return productAttributes, nil

}
