package gormlib

import (
	"context"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"

	"go.uber.org/zap"
)

type clientProductRepository struct {
	logger *zap.Logger
}

func NewClientProductRepository() domain.ClientProductRepository {
	return clientProductRepository{}
}

func (c clientProductRepository) List(db *infrastructure.Database, ctx context.Context, input domain.ClientListProductQueryOptions) ([]*models.Product, int64, error) {
	query := db.RDBMS.WithContext(ctx).Model(&models.Product{})
	if input.Price != nil {
		query = query.Where("price >= ?", input.Price)
	}

	var products []*models.Product
	var total int64
	err := GormQueryPagination(query, input.PageOptions, &products).Count(&total).Error()
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}
