package domain

import (
	"context"
	"productservice/internal/api/request"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
)

type CmsListProductQueryOptions struct {
	request.PageOptions
	Price *int64 `form:"price"`
}

type CmsProductRepository interface {
	Create(db *infrastructure.Database, ctx context.Context, product *models.Product) (res *models.Product, err error)
	GetById(db *infrastructure.Database, ctx context.Context, id string) (res *models.Product, err error)
	GetByIdJoinCategory(db *infrastructure.Database, ctx context.Context, id string) (res *models.Product, err error)
	List(db *infrastructure.Database, ctx context.Context, input CmsListProductQueryOptions) ([]*models.Product, int64, error)
	Update(db *infrastructure.Database, ctx context.Context, id string, updates map[string]interface{}) (err error)
}
type ClientListProductQueryOptions struct {
	request.PageOptions
	Price *int64 `form:"price"`
}
type ClientProductRepository interface {
	List(db *infrastructure.Database, ctx context.Context, input ClientListProductQueryOptions) ([]*models.Product, int64, error)
}

type CmsProductService interface {
	CreateProduct(ctx context.Context, req request.CmsCreateProductRequest) (product *models.Product, err error)
	GetProductByID(ctx context.Context, id string) (product *models.Product, err error)
	ListProduct(ctx context.Context, input request.CmsListProductRequest) (res []*models.Product, total int64, err error)
	DecreaseProductQuantity(ctx context.Context, id string, quantity int64) (err error)
	DecreaseProductQuantityMutex(ctx context.Context, id string, quantity int64) (err error)
}

type ProductService interface {
	// GetProductByID(ctx context.Context, id string) (product *models.Product, err error)
	GetProductList(ctx context.Context, input request.ClientListProductRequest) (res []*models.Product, total int64, err error)
}
