package service

import (
	"context"
	"fmt"
	"productservice/config"
	"productservice/internal/api/request"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
	"productservice/internal/repository"
	"productservice/internal/utils"
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

	arrModels := request.ToArrayProductModel(req.Variations, req.ProductID)

	a.ufw.ProductModelRepository.BulkCreate(&a.db, ctx, arrModels)

	return a.ufw.TierVariationRepository.Create(&a.db, ctx, &models.TierVariations{
		ProductID: uuid.FromStringOrNil(req.ProductID),
		Options:   request.ToOptionsTierVar(req.Variations),
	})
}

// UpdateProductAttributes
func (a *tierVariationService) UpdateTierVariationOptions(ctx context.Context, req request.TierVariationUpdate, id string) (err error) {

	t, err := a.ufw.TierVariationRepository.GetById(&a.db, ctx, id)

	if err != nil {
		return err
	}

	arrTierVal := []request.Variation{}

	for i, v := range t.Options {
		if v.ID == req.ID {
			arrTierVal = append(arrTierVal, request.Variation{
				Name:    v.Name,
				Options: req.Options,
			})
			for _, opt := range req.Options {
				t.Options[i].Options = append(t.Options[i].Options, opt)
			}
		} else {
			arrTierVal = append(arrTierVal, request.Variation{
				Name:    v.Name,
				Options: v.Options,
			})
		}

	}

	err = a.ufw.TierVariationRepository.Update(&a.db, ctx, map[string]interface{}{
		"options": t.Options,
	}, id)

	if err != nil {
		return err
	}

	fmt.Println(arrTierVal)

	_, err = a.ufw.ProductModelRepository.BulkCreate(&a.db, ctx, request.ToArrayProductModelUpdate(arrTierVal, t.ProductID.String(), len(t.Options)-1))

	if err != nil {
		return err
	}

	return nil
}

// DeleteTierVariationOptions
func (a *tierVariationService) DeleteTierVariationOptions(ctx context.Context, req request.TierVariationDelete, id string) (err error) {

	t, err := a.ufw.TierVariationRepository.GetById(&a.db, ctx, id)

	if err != nil {
		return err
	}

	posEleDel := 0

	for i, v := range t.Options {
		if v.ID == req.ID {
			for _, val := range req.Options {
				posEleDel = utils.FindIndex(v.Options, val)
				t.Options[i].Options = utils.RemoveIndex(t.Options[i].Options, posEleDel)
			}

		}

	}

	err = a.ufw.TierVariationRepository.Update(&a.db, ctx, map[string]interface{}{
		"options": t.Options,
	}, id)

	if err != nil {
		return err
	}

	err = a.ufw.ProductModelRepository.BulkDeleteByProductIdAndItemIndex(&a.db, ctx, req.ProductID, fmt.Sprintf("%d", posEleDel), req.ID)
	if err != nil {
		return err
	}

	return nil
}
