package controller

import (
	"net/http"
	"productservice/config"
	"productservice/internal/api/request"
	"productservice/internal/domain"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrderController struct {
	logger       *zap.Logger
	config       *config.Config
	orderService domain.OrderService
}

func NewOrderController(orderService domain.OrderService, logger *zap.Logger, config *config.Config) *OrderController {
	controller := &OrderController{
		orderService: orderService,
		logger:       logger,
		config:       config,
	}
	return controller
}

func (o *OrderController) GetOrdersOfCurrentUser(c *gin.Context) {
	orders, err := o.orderService.GetOrdersOfCurrentUser(c.Request.Context())
	if err != nil {
		ResponseError(c, err)
		return
	}

	Response(c, http.StatusOK, "success", orders)
}

func (o *OrderController) CreateOrder(c *gin.Context) {
	var req request.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseValidationError(c, err)
		return
	}

	order, err := o.orderService.CreateOrder(c.Request.Context(), req)
	if err != nil {
		ResponseError(c, err)
		return
	}

	Response(c, http.StatusOK, "success", order)
}
