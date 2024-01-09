package route

import (
	controller "productservice/internal/api/controllers"
)

type ProductModelRoutes struct {
}

func NewProductModelsRoutes(c *CmsGroupRoutes, controller *controller.ProductModelsController) *ProductModelRoutes {
	gr := c.g.Group("product/product_models")
	gr.GET("/:product_id", controller.GetListByProductId)

	return &ProductModelRoutes{}
}
