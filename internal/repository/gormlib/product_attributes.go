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

type productAttributesRepository struct {
}

func NewProductAttributesRepository() domain.ProductAttributesRepository {
	return productAttributesRepository{}
}

func (u productAttributesRepository) Create(db *infrastructure.Database, ctx context.Context, productAttributes *models.ProductAttributes) (res *models.ProductAttributes, err error) {
	if err := db.RDBMS.WithContext(ctx).Create(&productAttributes).Error; err != nil {
		return nil, errors.Cause(err)
	}

	return productAttributes, nil
}

// Update
func (u productAttributesRepository) Update(db *infrastructure.Database, ctx context.Context, id string, updates map[string]interface{}) (err error) {
	if err := db.RDBMS.WithContext(ctx).Model(&models.ProductAttributes{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// GetById
func (u productAttributesRepository) GetById(db *infrastructure.Database, ctx context.Context, id string) (res *models.ProductAttributes, err error) {
	var productAttributes models.ProductAttributes
	if err := db.RDBMS.WithContext(ctx).Where("id = ?", id).First(&productAttributes).Error; err != nil {

		if utils.ErrNoRows(err) {
			return nil, errors.New(api_errors.ErrProductAttributesNotFound)
		}

		return nil, errors.Cause(err)
	}

	return &productAttributes, nil
}
