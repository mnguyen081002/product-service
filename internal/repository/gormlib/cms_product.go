package gormlib

import (
	"context"
	"net/http"
	"productservice/internal/api_errors"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
	"productservice/internal/utils"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type cmsProductRepository struct {
	logger *zap.Logger
}

func NewCmsProductRepository() domain.CmsProductRepository {
	return cmsProductRepository{}
}

func (u cmsProductRepository) Create(db *infrastructure.Database, ctx context.Context, product *models.Product) (res *models.Product, err error) {
	if err := db.RDBMS.WithContext(ctx).Create(&product).Error; err != nil {
		return nil, errors.Cause(err)
	}

	return product, nil
}

func (u cmsProductRepository) GetById(db *infrastructure.Database, ctx context.Context, id string) (res *models.Product, err error) {
	if err := db.RDBMS.WithContext(ctx).First(&res, "id = ?", id).Error; err != nil {
		if utils.ErrNoRows(err) {
			return nil, api_errors.NewExternalError(api_errors.MapErrCodeMsg[api_errors.ErrProductNotFound],
				http.StatusNotFound, api_errors.ErrProductNotFound)
		}
		return nil, errors.Cause(err)
	}
	return
}

func (u cmsProductRepository) List(db *infrastructure.Database, ctx context.Context, input domain.CmsListProductQueryOptions) (products []*models.Product, total int64, err error) {
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

func (u cmsProductRepository) Update(db *infrastructure.Database, ctx context.Context, id string, updates map[string]interface{}) (err error) {
	if err := db.RDBMS.WithContext(ctx).Model(&models.Product{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return errors.Cause(err)
	}
	return nil
}

func (u cmsProductRepository) GetByIdJoinCategory(db *infrastructure.Database, ctx context.Context, id string) (res *models.Product, err error) {
	if err := db.RDBMS.WithContext(ctx).Preload("Category").First(&res, "id = ?", id).Error; err != nil {
		if utils.ErrNoRows(err) {
			return nil, api_errors.NewExternalError(api_errors.MapErrCodeMsg[api_errors.ErrProductNotFound],
				http.StatusNotFound, api_errors.ErrProductNotFound)
		}
		return nil, errors.Cause(err)
	}
	return
}
