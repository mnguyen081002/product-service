package service

import "go.uber.org/fx"

var Module = fx.Provide(
	NewCmsProductService,
	NewCmsProductAttributesService,
	NewTierVariationService,
	NewProductModel,
	NewCmsCategoryService,
	NewRatingService,
	NewProductService,
)
