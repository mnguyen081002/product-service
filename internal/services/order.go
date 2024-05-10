package service

import (
	"context"
	"productservice/internal/api/request"
	"productservice/internal/api_errors"
	"productservice/internal/constants"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
	"productservice/internal/utils"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type orderService struct {
	orderRepo        domain.OrderRepository
	productModelRepo domain.ProductModelRepository
	orderItemRepo    domain.OrderItemRepository
	dbTransaction    infrastructure.DatabaseTransaction
	db               infrastructure.Database
}

func NewOrderService(orderRepo domain.OrderRepository) domain.OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (o *orderService) CancelOrder(ctx context.Context, orderID string) (*models.Order, error) {
	panic("unimplemented")
}

func (o *orderService) CompleteOrder(ctx context.Context, orderID string) (*models.Order, error) {
	panic("unimplemented")
}

func (o *orderService) CreateOrder(ctx context.Context, input request.CreateOrderRequest) (*models.Order, error) {
	user_id, err := utils.GetUserUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	err = o.dbTransaction.WithTransaction(func(tx *infrastructure.Database) error {
		productModelIds := []uuid.UUID{}
		mapProductModel := map[uuid.UUID]int64{}
		for _, item := range input.Items {
			productModelIds = append(productModelIds, item.ProductModelID)
			mapProductModel[item.ProductModelID] = item.Quantity
		}
		productModels, err := o.productModelRepo.GetListProductModelByIds(tx, ctx, productModelIds)
		if err != nil {
			return err
		}

		missProductModels := []uuid.UUID{}
		for _, item := range productModels {
			if mapProductModel[item.ID] == 0 {
				missProductModels = append(missProductModels, item.ID)
			}
		}

		if len(missProductModels) > 0 {
			return errors.New(api_errors.ErrProductNotFound)
		}

		var totalPrice int64
		for _, item := range productModels {
			totalPrice += item.Price * mapProductModel[item.ID]
		}
		order := &models.Order{
			UserID: user_id.String(),
			Status: constants.OrderStatusPending,
		}

		order.TotalPrice = totalPrice

		r, err := o.orderRepo.CreateOrder(tx, ctx, order)
		if err != nil {
			return err
		}

		orderItems := []*models.OrderItem{}
		for _, item := range productModels {
			orderItem := &models.OrderItem{
				OrderID:        r.ID,
				ProductModelID: item.ID,
				Quantity:       mapProductModel[item.ID],
				Price:          item.Price,
				Status:         constants.OrderItemStatusPending,
				UserID:         user_id.String(),
			}
			orderItems = append(order.Items, orderItem)
		}

		_, err = o.orderItemRepo.BulkCreateOrderItems(tx, ctx, orderItems)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (o *orderService) DeliveryOrder(ctx context.Context, orderID string) (*models.Order, error) {
	panic("unimplemented")
}

func (o *orderService) GetOrder(ctx context.Context, orderID string) (*models.Order, error) {
	panic("unimplemented")
}

func (o *orderService) GetOrdersOfCurrentUser(ctx context.Context) ([]*models.Order, error) {
	userId, err := utils.GetUserUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	orders, err := o.orderRepo.GetOrdersByUserId(&o.db, ctx, userId.String())
	if err != nil {
		return nil, err
	}

	return orders, nil
}
