package producer

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"productservice/internal/constants"
	"productservice/internal/infrastructure"
	"productservice/internal/messaging/message"
)

type CmsProductProducer interface {
	DecreaseProductQuantity(ctx context.Context, msg message.DecreaseProductQuantity) error
}

type cmsProductProducer struct {
	infrastructure.Producer
	logger *zap.Logger
}

func NewCmsProductProducer(kafkaProducer infrastructure.Producer, logger *zap.Logger) CmsProductProducer {
	return &cmsProductProducer{
		Producer: kafkaProducer,
		logger:   logger,
	}
}

func (c cmsProductProducer) DecreaseProductQuantity(ctx context.Context, msg message.DecreaseProductQuantity) error {
	sm, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "Error marshal message")
	}

	go func() {
		err := c.PublishMessage(context.Background(), kafka.Message{
			Value: sm,
			Topic: constants.TopicCmsProductDecreaseProductQuantity,
		})
		if err != nil {
			c.logger.Error("Error publish message", zap.Error(err))
		}
	}()
	return nil
}
