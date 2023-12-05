package producer

import (
	"context"
	"encoding/json"
	"productservice/internal/constants"
	"productservice/internal/infrastructure"
	"productservice/internal/messaging/message"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

type CmsProductProducer interface {
	PublishDecreaseProductQuantity(ctx context.Context, msg message.DecreaseProductQuantity) error
}

type cmsProductProducer struct {
	infrastructure.Producer
}

func NewCmsProductProducer(kafkaProducer infrastructure.Producer) CmsProductProducer {
	return &cmsProductProducer{
		Producer: kafkaProducer,
	}
}

func (c cmsProductProducer) PublishDecreaseProductQuantity(ctx context.Context, msg message.DecreaseProductQuantity) error {
	sm, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "Error marshal message")
	}

	return c.PublishMessage(ctx, kafka.Message{
		Value: sm,
		Topic: constants.TopicCmsProductDecreaseProductQuantity,
	})
}
