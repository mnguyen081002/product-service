package service

import (
	"context"
	"encoding/json"
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

type cmsProductAttributesService struct {
	db                          infrastructure.Database
	dbTransaction               infrastructure.DatabaseTransaction
	cmsProductAttributesService domain.CmsProductAttributesService
	ufw                         *repository.UnitOfWork
	config                      *config.Config
	logger                      *zap.Logger
	mu                          sync.Mutex
}

func NewCmsProductAttributesService(
	db infrastructure.Database,
	dbTransaction infrastructure.DatabaseTransaction,
	ufw *repository.UnitOfWork,
	config *config.Config,
	logger *zap.Logger,
) domain.CmsProductAttributesService {
	return &cmsProductAttributesService{
		db:            db,
		dbTransaction: dbTransaction,
		ufw:           ufw,
		config:        config,
		logger:        logger,
	}
}

func (a *cmsProductAttributesService) CreateProductAttributes(ctx context.Context, req request.ProductAttributesCreate) (productAttributes *models.ProductAttributes, err error) {
	return a.ufw.ProductAttributesRepository.Create(&a.db, ctx, &models.ProductAttributes{
		ProductID:  uuid.FromStringOrNil(req.ProductID),
		Atrributes: request.ConvertAttributeModel(req.Attributes),
	})
}

// UpdateProductAttributes
func (a *cmsProductAttributesService) UpdateProductAttributes(ctx context.Context, id string, req request.ProductAttributesUpdate) (err error) {

	//check product attributes exist
	_, err = a.ufw.ProductAttributesRepository.GetById(&a.db, ctx, id)

	if err != nil {
		return err
	}

	r := request.ConvertAttributeModel(req.Attributes)
	ms, _ := json.Marshal(r)

	updates := map[string]interface{}{
		"attributes": ms,
	}
	return a.ufw.ProductAttributesRepository.Update(&a.db, ctx, id, updates)
}

// GetProductAttributesById
func (a *cmsProductAttributesService) GetProductAttributesById(ctx context.Context, id string) (productAttributes *models.ProductAttributes, err error) {
	return a.ufw.ProductAttributesRepository.GetById(&a.db, ctx, id)
}
