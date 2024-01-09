package route

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewCmsRoutes),
	fx.Invoke(
		NewHealthRoutes,
		NewCmsProductRoutes,
		NewCmsCategoryRoutes,
		NewRatingRoutes,
		NewProductRoutes,
		NewCmsProductAttributesRoutes,
		NewTierVariationRoutes,
		NewProductModelsRoutes,
	),
)
