package producer

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewCmsProductProducer),
)
