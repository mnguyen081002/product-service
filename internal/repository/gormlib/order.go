package gormlib

import (
	"context"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type orderRepository struct {
	logger *zap.Logger
}

func (u orderRepository) GetOrdersByUserId(db *infrastructure.Database, ctx context.Context, userID string) ([]*models.Order, error) {
	var orders []*models.Order
	if err := db.RDBMS.WithContext(ctx).Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, errors.Cause(err)
	}

	return orders, nil
}

func (u orderRepository) CreateOrder(db *infrastructure.Database, ctx context.Context, order *models.Order) (*models.Order, error) {
	if err := db.RDBMS.WithContext(ctx).Create(&order).Error; err != nil {
		return nil, errors.Cause(err)
	}

	return order, nil
}

func NewOrderRepository(logger *zap.Logger) domain.OrderRepository {
	return orderRepository{logger: logger}
}
