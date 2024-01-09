package controller

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewHealthController,
		NewCmsProductAttributesController,
		NewCmsProductController,
		NewCmsCategoryController,
		NewTierVariationController,
		NewProductModelsController,
		NewRatingController,
		NewProductController,
	),
)
