package domain

import (
	"context"
	"productservice/internal/api/request"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
)

type ProductRepository interface {
	Create(db *infrastructure.Database, ctx context.Context, product *models.Product) (res *models.Product, err error)
	GetById(db *infrastructure.Database, ctx context.Context, id string) (res *models.Product, err error)
	List(db *infrastructure.Database, ctx context.Context, input request.ListProductRequest) ([]*models.Product, int64, error)
	Update(db *infrastructure.Database, ctx context.Context, id string, updates map[string]interface{}) (err error)
}

type CmsProductService interface {
	CreateProduct(ctx context.Context, req request.CreateProductRequest) (product *models.Product, err error)
	GetProductById(ctx context.Context, id string) (product *models.Product, err error)
	ListProduct(ctx context.Context, input request.ListProductRequest) (res []*models.Product, total int64, err error)
	DecreaseProductQuantity(ctx context.Context, id string, quantity int64) (err error)
	DecreaseProductQuantityMutex(ctx context.Context, id string, quantity int64) (err error)
}
