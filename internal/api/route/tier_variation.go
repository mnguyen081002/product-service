package route

import (
	controller "productservice/internal/api/controllers"
)

type TierVariation struct {
}

func NewTierVariationRoutes(c *CmsGroupRoutes, controller *controller.TierVariationController) *TierVariation {
	gr := c.g.Group("/product/tier_variation")
	gr.POST("", controller.CreaTierVariation)
	gr.PUT("/options/:id", controller.UpdateTierVariation)
	gr.DELETE("/options/:id", controller.DeleteTierVariationOptions)

	return &TierVariation{}
}
