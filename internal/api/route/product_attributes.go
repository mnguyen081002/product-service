package route

import (
	controller "productservice/internal/api/controllers"
)

type CmsProductAttributeRoutes struct {
}

func NewCmsProductAttributesRoutes(c *CmsGroupRoutes, controller *controller.CmsProductAttributesController) *CmsProductAttributeRoutes {
	gr := c.g.Group("/product/attributes")
	gr.POST("", controller.CreateProductAttribute)

	return &CmsProductAttributeRoutes{}
}
