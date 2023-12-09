package producer

import (
	"context"
	"productservice/internal/infrastructure"
	"productservice/internal/messaging/message"

	"go.uber.org/zap"
)

type CmsProductAttributeProducer interface {
	DecreaseProductQuantity(ctx context.Context, msg message.DecreaseProductQuantity) error
}

type cmsProductAttributeProducer struct {
	infrastructure.Producer
	logger *zap.Logger
}

func NewCmsProductAttributeProducer(kafkaProducer infrastructure.Producer, logger *zap.Logger) CmsProductProducer {
	return &cmsProductProducer{
		Producer: kafkaProducer,
		logger:   logger,
	}
}
