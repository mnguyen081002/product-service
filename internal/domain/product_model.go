package domain

import (
	"context"
	"productservice/internal/api/request"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
)

type ProductModelRepository interface {
	Create(db *infrastructure.Database, ctx context.Context, productModel *models.ProductModels) (res *models.ProductModels, err error)
	Update(db *infrastructure.Database, ctx context.Context, id string, updates map[string]interface{}) (err error)
}

type ProductModelService interface {
	InitProductModel(ctx context.Context, req request.TierVariationCreate) (productModel *models.ProductModels, err error)
}
