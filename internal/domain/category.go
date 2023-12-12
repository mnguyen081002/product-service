package domain

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"productservice/internal/api/request"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
)

type CategoryRepository interface {
	Create(db *infrastructure.Database, ctx context.Context, cat *models.Category) (res *models.Category, err error)
	Update(db *infrastructure.Database, ctx context.Context, update map[string]interface{}, id uuid.UUID) (res *models.Category, err error)
	UpdateNoSub(db *infrastructure.Database, ctx context.Context, noSub bool, id uuid.UUID) (err error)
	CountSubCategoryOfCategory(db *infrastructure.Database, ctx context.Context, id uuid.UUID) (total int64, err error)
	FindByID(db *infrastructure.Database, ctx context.Context, id uuid.UUID) (cat *models.Category, err error)
	ListCategories(db *infrastructure.Database, ctx context.Context, options request.ListCategoryRequest) (categories []*models.Category, total int64, err error)
	GetProductCategories(db *infrastructure.Database, ctx context.Context, productID uuid.UUID) (categories []*models.Category, err error)
}

type CmsCategoryService interface {
	CreateCategory(ctx context.Context, req request.CreateCategoryRequest) (category *models.Category, err error)
	UpdateCategory(ctx context.Context, input request.UpdateCategoryRequest, id uuid.UUID) (category *models.Category, err error)
	ListCategories(ctx context.Context, options request.ListCategoryRequest) (categories []*models.Category, total int64, err error)
}

type CategoryService interface {
	GetProductCategories(ctx context.Context, productID uuid.UUID) (categories []*models.Category, err error)
}
