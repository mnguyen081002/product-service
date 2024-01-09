package route

import (
	controller "productservice/internal/api/controllers"
)

type CmsProductAttributeRoutes struct {
}

func NewCmsProductAttributesRoutes(c *CmsGroupRoutes, controller *controller.ProductAttributesController) *CmsProductAttributeRoutes {
	gr := c.g.Group("/product/attributes")
	gr.POST("", controller.CreateProductAttribute)
	gr.PUT("/:id", controller.UpdateProductAttribute)

	return &CmsProductAttributeRoutes{}
}
