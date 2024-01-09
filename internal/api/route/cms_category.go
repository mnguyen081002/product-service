package route

import (
	"productservice/internal/api/controllers"
)

type CmsCategoryRoutes struct {
}

func NewCmsCategoryRoutes(c *CmsGroupRoutes, controller *controller.CmsCategoryController) *CmsCategoryRoutes {
	gr := c.g.Group("/category")
	gr.GET("/", controller.ListCategories)
	gr.POST("/", controller.CreateCategory)
	gr.PUT("/:id", controller.UpdateCategory)
	return &CmsCategoryRoutes{}
}
