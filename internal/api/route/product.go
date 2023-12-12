package route

import (
	controller "productservice/internal/api/controllers"
	"productservice/internal/lib"
)

type ProductRoutes struct {
}

func NewProductRoutes(handler *lib.Handler, controller *controller.ProductController) *ProductRoutes {
	gr := handler.Group("/product")
	gr.GET("/:id", controller.GetProductById)
	return &ProductRoutes{}
}
