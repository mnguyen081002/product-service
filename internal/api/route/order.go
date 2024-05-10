package route

import (
	controller "productservice/internal/api/controllers"
	"productservice/internal/api/middlewares"
	"productservice/internal/lib"
)

type OrderRoutes struct {
}

func NewOrderRoutes(handler *lib.Handler, middleware *middlewares.GinMiddleware, controller *controller.OrderController) *OrderRoutes {
	g := handler.Group("/order", middleware.Auth(true))
	g.POST("/", controller.CreateOrder)
	return &OrderRoutes{}
}
