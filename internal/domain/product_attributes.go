package domain

import (
	"context"
	"productservice/internal/api/request"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
)

type ProductAttributesRepository interface {
	Create(db *infrastructure.Database, ctx context.Context, productAttributes *models.ProductAttributes) (res *models.ProductAttributes, err error)
	Update(db *infrastructure.Database, ctx context.Context, id string, updates map[string]interface{}) (err error)
	GetById(db *infrastructure.Database, ctx context.Context, id string) (res *models.ProductAttributes, err error)
	// List(db *infrastructure.Database, ctx context.Context, input request.ListProductRequest) ([]*models.ProductAttributes, *int64, error)
}

type CmsProductAttributesService interface {
	CreateProductAttributes(ctx context.Context, req request.ProductAttributesCreate) (product *models.ProductAttributes, err error)
	UpdateProductAttributes(ctx context.Context, id string, req request.ProductAttributesUpdate) (err error)
	GetProductAttributesById(ctx context.Context, id string) (product *models.ProductAttributes, err error)
}
