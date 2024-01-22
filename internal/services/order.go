package service

import (
	"context"
	"productservice/config"
	"productservice/internal/api/request"
	"productservice/internal/constants"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
	"productservice/internal/repository"
	"sync"

	"go.uber.org/zap"
)

type orderService struct {
	db                  infrastructure.Database
	dbTransaction       infrastructure.DatabaseTransaction
	orderService        domain.OrderService
	productService      domain.ProductService
	productModelService domain.ProductModelService
	ufw                 *repository.UnitOfWork
	config              *config.Config
	logger              *zap.Logger
	mu                  sync.Mutex
}

func NewCmsOrderService(
	db infrastructure.Database,
	dbTransaction infrastructure.DatabaseTransaction,
	ufw *repository.UnitOfWork,
	config *config.Config,
	logger *zap.Logger,
	proproductService domain.ProductService,
	productModelService domain.ProductModelService,
) domain.OrderService {
	return &orderService{
		db:                  db,
		dbTransaction:       dbTransaction,
		ufw:                 ufw,
		config:              config,
		logger:              logger,
		productService:      proproductService,
		productModelService: productModelService,
	}
}

func (a *orderService) CreateOrder(ctx context.Context, req request.CreateOrderRequest, userID string) (order *models.Order, err error) {
	product, err := a.productService.GetProductByID(ctx, req.ProductID)

	if err != nil {
		return nil, err
	}

	productModel, err := a.productModelService.GetProductModelByID(ctx, req.ProductModelID)

	if err != nil {
		return nil, err
	}

	return a.ufw.OrderRepository.Create(&a.db, ctx, &models.Order{
		Product:      *product,
		ProductModel: *productModel,
		Status:       constants.OrderStatusConfirmed,
		Quantity:     req.Quantity,
		TotalPrice:   req.TotalPrice,
	})
}

func (a *orderService) UpdateOrder(ctx context.Context, req request.UpdateOrderRequest, orderID string) (order *models.Order, err error) {
	product, err := a.productService.GetProductByID(ctx, req.ProductID)

	if err != nil {
		return nil, err
	}

	productModel, err := a.productModelService.GetProductModelByID(ctx, req.ProductModelID)

	if err != nil {
		return nil, err
	}

	return a.ufw.OrderRepository.Update(&a.db, ctx, &models.Order{
		Product:      *product,
		ProductModel: *productModel,
		Status:       req.Status,
		Quantity:     req.Quantity,
		TotalPrice:   req.TotalPrice,
		Reason:       req.Reason,
	}, orderID)
}

func (a *orderService) GetList(ctx context.Context, req request.GetListRequest, userID string) (res []*models.Order, err error) {
	return a.ufw.OrderRepository.GetList(&a.db, ctx, req, userID)
}

func (a *orderService) Delete(ctx context.Context, orderID string) (err error) {
	return a.ufw.OrderRepository.Delete(&a.db, ctx, orderID)
}
