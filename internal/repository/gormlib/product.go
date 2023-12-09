package gormlib

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"productservice/internal/api/request"
	"productservice/internal/api_errors"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
	"productservice/internal/utils"
)

type productRepository struct {
	logger *zap.Logger
}

func NewProductRepository() domain.ProductRepository {
	return productRepository{}
}

func (u productRepository) Create(db *infrastructure.Database, ctx context.Context, product *models.Product) (res *models.Product, err error) {
	if err := db.RDBMS.WithContext(ctx).Create(&product).Error; err != nil {
		return nil, errors.Cause(err)
	}

	return product, nil
}

func (u productRepository) GetById(db *infrastructure.Database, ctx context.Context, id string) (res *models.Product, err error) {
	if err := db.RDBMS.WithContext(ctx).First(&res, "id = ?", id).Error; err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New(api_errors.ErrProductNotFound)
		}
		return nil, errors.Cause(err)
	}
	return
}

func (u productRepository) List(db *infrastructure.Database, ctx context.Context, input request.ListProductRequest) (products []*models.Product, total int64, err error) {
	dbQuery := db.RDBMS.WithContext(ctx).Model(&models.Product{})
	if input.Search != nil {
		dbQuery = dbQuery.Where("name ILIKE ?", "%"+*input.Search+"%")
	}

	if input.Price != nil {
		dbQuery = dbQuery.Where("price >= ?", input.Price)
	}

	err = GormQueryPagination(dbQuery, input.PageOptions, &products).Error()
	if err != nil {
		return nil, 0, errors.Cause(err)
	}
	return products, total, nil
}

func (u productRepository) Update(db *infrastructure.Database, ctx context.Context, id string, updates map[string]interface{}) (err error) {
	if err := db.RDBMS.WithContext(ctx).Model(&models.Product{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return errors.Cause(err)
	}
	return nil
}
