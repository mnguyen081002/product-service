package service

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"productservice/config"
	"productservice/internal/api/request"
	"productservice/internal/api_errors"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
	"productservice/internal/repository"
	"sync"
	"time"
)

type cmsProductService struct {
	db                infrastructure.Database
	dbTransaction     infrastructure.DatabaseTransaction
	cmsProductService domain.CmsProductService
	ufw               *repository.UnitOfWork
	config            *config.Config
	logger            *zap.Logger
	mu                sync.Mutex
}

func NewCmsProductService(
	db infrastructure.Database,
	dbTransaction infrastructure.DatabaseTransaction,
	ufw *repository.UnitOfWork,
	config *config.Config,
	logger *zap.Logger,
) domain.CmsProductService {
	return &cmsProductService{
		db:            db,
		dbTransaction: dbTransaction,
		ufw:           ufw,
		config:        config,
		logger:        logger,
	}
}

func (a *cmsProductService) CreateProduct(ctx context.Context, req request.CreateProductRequest) (product *models.Product, err error) {
	return a.ufw.ProductRepository.Create(&a.db, ctx, &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
		Images:      req.Images,
	})
}

func (a *cmsProductService) GetProductById(ctx context.Context, id string) (product *models.Product, err error) {
	return a.ufw.ProductRepository.GetById(&a.db, ctx, id)
}

func (a *cmsProductService) ListProduct(ctx context.Context, input request.ListProductRequest) (res []*models.Product, total *int64, err error) {
	list, total, err := a.ufw.ProductRepository.List(&a.db, ctx, input)
	if err != nil {
		return nil, nil, err
	}
	return list, total, nil
}

func (a *cmsProductService) DecreaseProductQuantity(ctx context.Context, id string, quantity int64) (err error) {
	if quantity < 0 {
		return errors.New(api_errors.ErrQuantityMustHigherThanZero)
	}
	product, err := a.ufw.ProductRepository.GetById(&a.db, ctx, id)
	if err != nil {
		return err
	}

	if product.Quantity < quantity {
		return errors.New(api_errors.ErrQuantityNotEnough)
	}

	return a.ufw.ProductRepository.Update(&a.db, ctx, id, map[string]interface{}{
		"quantity": product.Quantity - quantity,
	})
}

func (a *cmsProductService) DecreaseProductQuantityMutex(ctx context.Context, id string, quantity int64) (err error) {
	if quantity < 0 {
		return errors.New(api_errors.ErrQuantityMustHigherThanZero)
	}
	defer a.mu.Unlock()
	a.mu.Lock()
	product, err := a.ufw.ProductRepository.GetById(&a.db, ctx, id)
	if err != nil {
		return err
	}
	time.Sleep(100 * time.Millisecond) // for test concurrency

	if product.Quantity < quantity {
		return errors.New(api_errors.ErrQuantityNotEnough)
	}

	err = a.ufw.ProductRepository.Update(&a.db, ctx, id, map[string]interface{}{
		"quantity": product.Quantity - quantity,
	})
	if err != nil {
		return err
	}
	return nil
}
