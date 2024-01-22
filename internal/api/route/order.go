package route

import (
	controller "productservice/internal/api/controllers"
)

type OrderRoutes struct {
}

func NewOderRoutes(c *CmsGroupRoutes, controller *controller.OrderController) *OrderRoutes {
	gr := c.g.Group("/order")
	gr.POST("", controller.CreateOrder)
	gr.PUT("/:id", controller.UpdateOrder)
	gr.GET("", controller.GetList)
	gr.DELETE("/:id", controller.Delete)

	return &OrderRoutes{}
}
