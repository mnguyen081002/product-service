package gormlib

import (
	"context"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"

	"github.com/pkg/errors"
)

type productModelRepository struct {
}

func NewProductModelRepository() domain.ProductModelRepository {
	return productModelRepository{}
}

func (u productModelRepository) Create(db *infrastructure.Database, ctx context.Context, productModels *models.ProductModels) (res *models.ProductModels, err error) {
	if err := db.RDBMS.WithContext(ctx).Create(&productModels).Error; err != nil {
		return nil, errors.Cause(err)
	}

	return productModels, nil
}

// Update
func (u productModelRepository) Update(db *infrastructure.Database, ctx context.Context, id string, updates map[string]interface{}) (err error) {
	if err := db.RDBMS.WithContext(ctx).Model(&models.ProductAttributes{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
