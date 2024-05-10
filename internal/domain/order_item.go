package domain

import (
	"context"
	"productservice/internal/infrastructure"
	"productservice/internal/models"

	uuid "github.com/satori/go.uuid"
)

type OrderItemService interface {
	GetOrderItem(id uuid.UUID) (*models.OrderItem, error)
	CreateOrderItem(orderItem *models.OrderItem) (*models.OrderItem, error)
	CancelOrderItem(id uuid.UUID) (*models.OrderItem, error)
	RefundOrderItem(id uuid.UUID) (*models.OrderItem, error)
	DeliveryOrderItem(id uuid.UUID) (*models.OrderItem, error)
}

type OrderItemRepository interface {
	GetOrderItem(id uuid.UUID) (*models.OrderItem, error)
	BulkCreateOrderItems(db *infrastructure.Database, ctx context.Context, orderItems []*models.OrderItem) ([]*models.OrderItem, error)
}
