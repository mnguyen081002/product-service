package messaging

import (
	"go.uber.org/fx"
	"productservice/internal/messaging/subscriber"
)

type Subscriber interface {
	Subscribe()
}

var Module = fx.Invoke(subscriber.NewUpdateProductSubscribe)
