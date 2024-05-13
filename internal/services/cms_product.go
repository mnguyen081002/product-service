package service

import (
	"context"
	"net/http"
	"productservice/config"
	"productservice/internal/api/request"
	"productservice/internal/api_errors"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
	"productservice/internal/repository"
	"sync"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
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

func (a *cmsProductService) CreateProduct(ctx context.Context, req request.CmsCreateProductRequest) (product *models.Product, err error) {
	catId, err := uuid.FromString(req.CategoryID)
	if err != nil {
		return nil, err
	}
	var p *models.Product
	err = a.dbTransaction.WithTransaction(func(tx *infrastructure.Database) error {
		p, err = a.ufw.CmsProductRepository.Create(tx, ctx, &models.Product{
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
			Quantity:    req.Quantity,
			Images:      req.Images,
			CategoryID:  &catId,
		})
		if err != nil {
			return err
		}
		_, err = a.ufw.CmsTierVariationRepository.Create(tx, ctx, &models.TierVariations{
			ProductID: p.ID,
			Options:   request.ToOptionsTierVar(req.TierVariants),
		})
		if err != nil {
			return err
		}
		_, err = a.ufw.CmsProductModelRepository.BulkCreate(tx, ctx, request.ToArrayProductModel(req.TierVariants, p.ID.String()))
		if err != nil {
			return err
		}
		_, err = a.ufw.CmsProductAttributesRepository.Create(tx, ctx, &models.ProductAttributes{
			ProductID:  p.ID,
			Atrributes: request.ConvertAttributeModel(req.Attributes),
		})
		return err
	})
	return p, err
}

func (a *cmsProductService) GetProductByID(ctx context.Context, id string) (product *models.Product, err error) {
	return a.ufw.CmsProductRepository.GetById(&a.db, ctx, id)
}

func (a *cmsProductService) ListProduct(ctx context.Context, input request.CmsListProductRequest) (res []*models.Product, total int64, err error) {
	list, total, err := a.ufw.CmsProductRepository.List(&a.db, ctx, domain.CmsListProductQueryOptions{
		PageOptions: input.PageOptions,
		Price:       input.Price,
	})
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (a *cmsProductService) DecreaseProductQuantity(ctx context.Context, id string, quantity int64) (err error) {
	if quantity < 0 {
		return api_errors.NewExternalError(api_errors.MapErrCodeMsg[api_errors.ErrQuantityMustHigherThanZero],
			http.StatusBadRequest,
			api_errors.ErrQuantityMustHigherThanZero)
	}
	product, err := a.ufw.CmsProductRepository.GetById(&a.db, ctx, id)
	if err != nil {
		return err
	}

	if product.Quantity < quantity {
		return api_errors.NewExternalError(api_errors.MapErrCodeMsg[api_errors.ErrQuantityNotEnough],
			http.StatusBadRequest,
			api_errors.ErrQuantityNotEnough)
	}

	return a.ufw.CmsProductRepository.Update(&a.db, ctx, id, map[string]interface{}{
		"quantity": product.Quantity - quantity,
	})
}

func (a *cmsProductService) DecreaseProductQuantityMutex(ctx context.Context, id string, quantity int64) (err error) {
	if quantity < 0 {
		return errors.New(api_errors.ErrQuantityMustHigherThanZero)
	}
	defer a.mu.Unlock()
	a.mu.Lock()
	product, err := a.ufw.CmsProductRepository.GetById(&a.db, ctx, id)
	if err != nil {
		return err
	}
	time.Sleep(100 * time.Millisecond) // for test concurrency

	if product.Quantity < quantity {
		return errors.New(api_errors.ErrQuantityNotEnough)
	}

	err = a.ufw.CmsProductRepository.Update(&a.db, ctx, id, map[string]interface{}{
		"quantity": product.Quantity - quantity,
	})
	if err != nil {
		return err
	}
	return nil
}
