package domain

import (
	"context"
	"productservice/internal/api/request"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
)

type OrderRepository interface {
	Create(db *infrastructure.Database, ctx context.Context, order *models.Order) (res *models.Order, err error)
	Update(db *infrastructure.Database, ctx context.Context, order *models.Order, orderID string) (res *models.Order, err error)
	GetList(db *infrastructure.Database, ctx context.Context, req request.GetListRequest, userID string) (res []*models.Order, err error)
	Delete(db *infrastructure.Database, ctx context.Context, orderID string) (err error)
}

type OrderService interface {
	CreateOrder(ctx context.Context, req request.CreateOrderRequest, userID string) (order *models.Order, err error)
	UpdateOrder(ctx context.Context, req request.UpdateOrderRequest, orderID string) (order *models.Order, err error)
	GetList(ctx context.Context, req request.GetListRequest, userID string) (res []*models.Order, err error)
	Delete(ctx context.Context, orderID string) (err error)
}
