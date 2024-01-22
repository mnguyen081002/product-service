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
	BulkDeleteByProductIdAndItemIndex(db *infrastructure.Database, ctx context.Context, productID string, char string, optionsID int) (err error)
	GetListByProductId(db *infrastructure.Database, ctx context.Context, productID string) (res []*models.ProductModel, err error)
	GetProductModelByID(db *infrastructure.Database, ctx context.Context, id string) (res *models.ProductModel, err error)
}

type ProductModelService interface {
	CreateProductModels(ctx context.Context, req request.CreateTierVariation) (productModel *models.ProductModel, err error)
	GetListByProductId(ctx context.Context, productID string) (res []*models.ProductModel, err error)
	GetProductModelByID(ctx context.Context, id string) (res *models.ProductModel, err error)
}
