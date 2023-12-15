package gormlib

import (
	"context"
	"productservice/internal/api/request"
	"productservice/internal/api_errors"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
	"productservice/internal/utils"

	"github.com/pkg/errors"
	"go.uber.org/zap"
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

func (u productRepository) List(db *infrastructure.Database, ctx context.Context, input request.ListProductRequest) ([]*models.Product, *int64, error) {
	dbQuery := db.RDBMS.WithContext(ctx).Model(&models.Product{})
	var res []*models.Product
	if input.Search != "" {
		dbQuery = dbQuery.Where("name ILIKE ?", "%"+input.Search+"%")
	}

	if input.Price != nil {
		dbQuery = dbQuery.Where("price >= ?", input.Price)
	}

	total := new(int64)

	err := GormQueryPagination(dbQuery, input.PageOptions, &res).Error()
	if err != nil {
		return nil, nil, errors.Cause(err)
	}
	return res, total, nil
}

func (u productRepository) Update(db *infrastructure.Database, ctx context.Context, id string, updates map[string]interface{}) (err error) {
	if err := db.RDBMS.WithContext(ctx).Model(&models.Product{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return errors.Cause(err)
	}
	return nil
}
