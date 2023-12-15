package gormlib

import (
	"context"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"

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

func (u tierVariationRepository) Update(db *infrastructure.Database, ctx context.Context, tierVariations map[string]interface{}, id string) (err error) {
	if err := db.RDBMS.WithContext(ctx).Model(&models.TierVariations{}).Where("id = ?", id).Updates(&tierVariations).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (u tierVariationRepository) GetById(db *infrastructure.Database, ctx context.Context, id string) (res *models.TierVariations, err error) {
	var tierVariations models.TierVariations
	if err := db.RDBMS.WithContext(ctx).Where("id = ?", id).First(&tierVariations).Error; err != nil {
		return nil, errors.Cause(err)
	}

	return &tierVariations, nil
}
