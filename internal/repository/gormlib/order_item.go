package gormlib

import (
	"context"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type orderItemRepository struct {
	logger *zap.Logger
}

func (o orderItemRepository) BulkCreateOrderItems(db *infrastructure.Database, ctx context.Context, orderItems []*models.OrderItem) ([]*models.OrderItem, error) {
	err := db.RDBMS.WithContext(ctx).Create(&orderItems).Error
	if err != nil {
		return nil, errors.Cause(err)
	}
	return orderItems, nil
}

func (o orderItemRepository) GetOrderItem(id uuid.UUID) (*models.OrderItem, error) {
	return nil, nil
}

func NewOrderItemRepository(logger *zap.Logger) domain.OrderItemRepository {
	return orderItemRepository{
		logger: logger,
	}
}
