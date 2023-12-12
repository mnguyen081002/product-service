package controller

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewHealthController,
		NewCmsProductController,
		NewCmsCategoryController,
		NewRatingController,
		NewProductController,
	),
)
