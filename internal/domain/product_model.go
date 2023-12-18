package domain

import (
	"context"
	"productservice/internal/api/request"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
)

type ProductModelRepository interface {
	Create(db *infrastructure.Database, ctx context.Context, productModel *models.ProductModel) (res *models.ProductModel, err error)
	BulkCreate(db *infrastructure.Database, ctx context.Context, productModels []*models.ProductModel) (res []*models.ProductModel, err error)
	Update(db *infrastructure.Database, ctx context.Context, id string, updates map[string]interface{}) (err error)
	CountWithCondition(db *infrastructure.Database, ctx context.Context, condition map[string]interface{}) (count int64, err error)
	BulkDeleteByProductIdAndItemIndex(db *infrastructure.Database, ctx context.Context, productID string, firstChar string, optionsID int) (err error)
}

type ProductModelService interface {
	InitProductModel(ctx context.Context, req request.TierVariationCreate) (productModel *models.ProductModel, err error)
}
