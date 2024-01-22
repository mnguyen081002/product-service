package controller

import (
	"net/http"
	"productservice/internal/api/request"
	"productservice/internal/domain"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrderController struct {
	orderService domain.OrderService
	logger       *zap.Logger
}

func NewOrderController(authService domain.OrderService, logger *zap.Logger) *OrderController {
	controller := &OrderController{
		orderService: authService,
		logger:       logger,
	}
	return controller
}

func (controller *OrderController) CreateOrder(ctx *gin.Context) {
	var request request.CreateOrderRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ResponseValidationError(ctx, err)
		return
	}

	uid := ctx.Request.Header.Get("x-user-id")

	data, err := controller.orderService.CreateOrder(ctx.Request.Context(), request, uid)
	if err != nil {
		ResponseError(ctx, err)
		return
	}

	Response(ctx, http.StatusOK, "success", data)
}

func (controller *OrderController) UpdateOrder(ctx *gin.Context) {
	var request request.UpdateOrderRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ResponseValidationError(ctx, err)
		return
	}

	order_id := ctx.Param("id")

	data, err := controller.orderService.UpdateOrder(ctx.Request.Context(), request, order_id)
	if err != nil {
		ResponseError(ctx, err)
		return
	}

	Response(ctx, http.StatusOK, "success", data)
}

func (controller *OrderController) GetList(ctx *gin.Context) {
	var request request.GetListRequest
	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		ResponseValidationError(ctx, err)
		return
	}

	userId := ctx.Request.Header.Get("x-user-id")

	data, err := controller.orderService.GetList(ctx.Request.Context(), request, userId)
	if err != nil {
		ResponseError(ctx, err)
		return
	}

	Response(ctx, http.StatusOK, "success", data)
}

func (controller *OrderController) Delete(ctx *gin.Context) {
	order_id := ctx.Param("id")

	err := controller.orderService.Delete(ctx.Request.Context(), order_id)
	if err != nil {
		ResponseError(ctx, err)
		return
	}

	Response(ctx, http.StatusOK, "success", nil)
}
