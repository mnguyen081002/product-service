package route

import (
	"productservice/internal/api/controllers"
)

type CmsProductRoutes struct {
}

func NewCmsProductRoutes(c *CmsGroupRoutes, controller *controller.CmsProductController) *CmsProductRoutes {
	gr := c.g.Group("/product")
	gr.POST("/", controller.CreateProduct)
	gr.GET("/", controller.ListProduct)
	gr.GET("/:id", controller.GetProductById)
	return &CmsProductRoutes{}
}
