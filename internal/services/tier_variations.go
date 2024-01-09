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

func (a *tierVariationService) CreateTierVariation(ctx context.Context, req request.CreateTierVariation) (tierVar *models.TierVariations, err error) {

	arrModels := request.ToArrayProductModel(req.Variations, req.ProductID)

	err = a.dbTransaction.WithTransaction(func(tx *infrastructure.Database) error {

		tierVar, err = a.ufw.TierVariationRepository.Create(tx, ctx, &models.TierVariations{
			ProductID: uuid.FromStringOrNil(req.ProductID),
			Options:   request.ToOptionsTierVar(req.Variations),
		})

		if err != nil {
			return err
		}

		_, err = a.ufw.ProductModelRepository.BulkCreate(tx, ctx, arrModels)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return tierVar, err
	}

	return tierVar, err
}

// UpdateProductAttributes
func (a *tierVariationService) UpdateTierVariationOptions(ctx context.Context, req request.UpdateTierVariation, id string) (err error) {
	t, err := a.ufw.TierVariationRepository.GetByProductID(&a.db, ctx, id)

	if err != nil {
		return err
	}

	arrTierVal := []request.Variation{}

	for i, v := range t.Options {
		if v.ID == req.ID {
			for _, opt := range req.Options {
				t.Options[i].Options = append(t.Options[i].Options, opt)
			}
		}
		arrTierVal = append(arrTierVal, request.Variation{
			Name:    v.Name,
			Options: v.Options,
		})
	}

	err = a.dbTransaction.WithTransaction(func(tx *infrastructure.Database) error {

		err = a.ufw.TierVariationRepository.UpdateByProductId(tx, ctx, map[string]interface{}{
			"options": t.Options,
		}, id)

		if err != nil {
			return err
		}

		lenProductModelsUpdate := models.GetLengthOptions(t.Options, req.ID) - 1

		_, err = a.ufw.ProductModelRepository.BulkCreate(tx, ctx, request.ToArrayProductModelUpdate(arrTierVal, t.ProductID.String(), lenProductModelsUpdate))

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// DeleteTierVariationOptions
func (a *tierVariationService) DeleteTierVariationOptions(ctx context.Context, req request.DeleteTierVariation) (err error) {

	id := req.ProductID

	t, err := a.ufw.TierVariationRepository.GetByProductID(&a.db, ctx, id)

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

	err = a.dbTransaction.WithTransaction(func(tx *infrastructure.Database) error {

		err = a.ufw.TierVariationRepository.UpdateByProductId(tx, ctx, map[string]interface{}{
			"options": t.Options,
		}, id)

		if err != nil {
			return err
		}

		err = a.ufw.ProductModelRepository.BulkDeleteByProductIdAndItemIndex(tx, ctx, req.ProductID, fmt.Sprintf("%d", posEleDel), req.ID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// GetByProductID
func (a *tierVariationService) GetByProductID(ctx context.Context, id string) (res *models.TierVariations, err error) {
	return a.ufw.TierVariationRepository.GetByProductID(&a.db, ctx, id)
}
