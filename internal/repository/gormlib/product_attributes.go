package gormlib

import (
	"context"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"

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
