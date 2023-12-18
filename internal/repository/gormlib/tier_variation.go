package gormlib

import (
	"context"
	"productservice/internal/api_errors"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
	"productservice/internal/utils"

	"github.com/pkg/errors"
)

type tierVariationRepository struct {
}

func NewTierVariationRepository() domain.TierVariationRepository {
	return tierVariationRepository{}
}

func (u tierVariationRepository) Create(db *infrastructure.Database, ctx context.Context, tierVariations *models.TierVariations) (res *models.TierVariations, err error) {
	if err := db.RDBMS.WithContext(ctx).Create(&tierVariations).Error; err != nil {
		return nil, errors.Cause(err)
	}

	return tierVariations, nil
}

func (u tierVariationRepository) UpdateByProductId(db *infrastructure.Database, ctx context.Context, updates map[string]interface{}, id string) (err error) {
	tx := db.RDBMS.WithContext(ctx).Model(&models.TierVariations{}).Where("product_id = ?", id).Updates(&updates)

	if tx.Error != nil {
		return errors.WithStack(tx.Error)
	}

	if tx.RowsAffected == 0 {
		return errors.New(api_errors.ErrTierVariationNotFound)
	}

	return nil
}

func (u tierVariationRepository) GetById(db *infrastructure.Database, ctx context.Context, id string) (res *models.TierVariations, err error) {

	var tierVariations models.TierVariations
	if err := db.RDBMS.WithContext(ctx).Where("id = ?", id).First(&tierVariations).Error; err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New(api_errors.ErrTierVariationNotFound)
		}
		return nil, errors.Cause(err)
	}

	return &tierVariations, nil
}

// BulkDeleteWithCondition
func (u tierVariationRepository) BulkDeleteWithCondition(db *infrastructure.Database, ctx context.Context, cond map[string]interface{}) (err error) {
	if err := db.RDBMS.WithContext(ctx).Where(cond).Delete(&models.TierVariations{}).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// GetByProductID
func (u tierVariationRepository) GetByProductID(db *infrastructure.Database, ctx context.Context, productId string) (res *models.TierVariations, err error) {

	var tierVariations models.TierVariations
	if err := db.RDBMS.WithContext(ctx).Where("product_id = ?", productId).First(&tierVariations).Error; err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New(api_errors.ErrTierVariationNotFound)
		}
		return nil, errors.Cause(err)
	}

	return &tierVariations, nil
}
