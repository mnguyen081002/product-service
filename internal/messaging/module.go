package messaging

import (
	"go.uber.org/fx"
	"productservice/internal/messaging/producer"
	"productservice/internal/messaging/subscriber"
)

type Subscriber interface {
	Subscribe()
}

var Module = fx.Options(
	fx.Provide(producer.NewCmsProductProducer),
	fx.Invoke(subscriber.NewUpdateProductSubscribe),
)
