package route

import (
	controller "productservice/internal/api/controllers"
	"productservice/internal/api/middlewares"
	"productservice/internal/lib"
)

type RatingRoutes struct {
}

func NewRatingRoutes(handler *lib.Handler, controller *controller.RatingController, middleware *middlewares.GinMiddleware) *RatingRoutes {
	gr := handler.Group("/rating")
	gr.POST("/", middleware.Auth(true), controller.CreateRating)
	gr.GET("/", controller.ListRatingByProductID)
	return &RatingRoutes{}
}
