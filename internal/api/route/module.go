package route

import (
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(NewCmsRoutes), fx.Invoke(NewCmsProductRoutes, NewHealthRoutes))
