package route

import (
	controller "productservice/internal/api/controllers"
)

type TierVariation struct {
}

func NewTierVariationRoutes(c *CmsGroupRoutes, controller *controller.TierVariationController) *TierVariation {
	gr := c.g.Group("/product/tier_variation")
	gr.GET("/:product_id", controller.GetTierVariationByProductID)
	gr.POST("", controller.CreaTierVariation)
	gr.PUT("/options/:id", controller.UpdateTierVariation)
	gr.DELETE("/options", controller.DeleteTierVariationOptions)

	return &TierVariation{}
}
