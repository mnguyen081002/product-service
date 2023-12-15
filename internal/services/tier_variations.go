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

type tierVariationService struct {
	db                          infrastructure.Database
	dbTransaction               infrastructure.DatabaseTransaction
	cmsProductAttributesService domain.CmsProductAttributesService
	ufw                         *repository.UnitOfWork
	config                      *config.Config
	logger                      *zap.Logger
	mu                          sync.Mutex
}

func NewTierVariationService(
	db infrastructure.Database,
	dbTransaction infrastructure.DatabaseTransaction,
	ufw *repository.UnitOfWork,
	config *config.Config,
	logger *zap.Logger,
) domain.TierVariationService {
	return &tierVariationService{
		db:            db,
		dbTransaction: dbTransaction,
		ufw:           ufw,
		config:        config,
		logger:        logger,
	}
}

func (a *tierVariationService) CreateTierVariation(ctx context.Context, req request.TierVariationCreate) (tierVar *models.TierVariations, err error) {

	for _, p := range req.ToArrayProductModel() {
		a.ufw.ProductModelRepository.Create(&a.db, ctx, &models.ProductModels{
			ProductID: uuid.FromStringOrNil(req.ProductID),
			Models:    p,
		})
	}

	return a.ufw.TierVariationRepository.Create(&a.db, ctx, &models.TierVariations{
		ProductID: uuid.FromStringOrNil(req.ProductID),
		Options:   req.ToModelTierVar().Options,
	})
}

// UpdateProductAttributes
func (a *tierVariationService) UpdateTierVariationOptions(ctx context.Context, req request.TierVariationUpdate, id string) (err error) {

	//check product attributes exist
	_, err = a.ufw.TierVariationRepository.GetById(&a.db, ctx, id)

	if err != nil {
		return err
	}

	r := req.ToOptionsTierVar()
	ms, _ := json.Marshal(r)

	updates := map[string]interface{}{
		"options": ms,
	}
	return a.ufw.TierVariationRepository.Update(&a.db, ctx, updates, id)
}
