package domain

import (
	"context"
	"productservice/internal/api/request"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
)

type OrderService interface {
	CreateOrder(ctx context.Context, input request.CreateOrderRequest) (*models.Order, error)
	GetOrder(ctx context.Context, orderID string) (*models.Order, error)
	GetOrdersOfCurrentUser(ctx context.Context) ([]*models.Order, error)
	CancelOrder(ctx context.Context, orderID string) (*models.Order, error)
	CompleteOrder(ctx context.Context, orderID string) (*models.Order, error)
	DeliveryOrder(ctx context.Context, orderID string) (*models.Order, error)
}

type OrderRepository interface {
	CreateOrder(db *infrastructure.Database, ctx context.Context, order *models.Order) (*models.Order, error)
	GetOrdersByUserId(db *infrastructure.Database, ctx context.Context, userID string) ([]*models.Order, error)
}
