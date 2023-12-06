package subscriber

import (
	"context"
	"go.uber.org/zap"
	"productservice/internal/constants"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/messaging/mapper"
)

type UpdateProductSubscribe struct {
	kafkaSubscribe    infrastructure.KafkaSubscribe
	cmsProductService domain.CmsProductService
	logger            *zap.Logger
}

func NewUpdateProductSubscribe(cmsProductService domain.CmsProductService, logger *zap.Logger) {
	kafkaSub := infrastructure.NewKafkaSubscribe(constants.TopicCmsProductDecreaseProductQuantity)
	s := UpdateProductSubscribe{kafkaSubscribe: kafkaSub, cmsProductService: cmsProductService, logger: logger}
	go s.Subscribe()
}

func (s *UpdateProductSubscribe) Subscribe() {
	s.logger.Info("Waiting message from topic " + constants.TopicCmsProductDecreaseProductQuantity)
	for {
		m, err := s.kafkaSubscribe.FetchMessage(context.Background())
		if err != nil {
			s.logger.Error("Error message", zap.Error(err))
			continue
		}
		msg := mappermessage.DecreaseProductQuantityMessage(m)
		ctx := context.WithValue(context.Background(), "x-user-id", msg.UserID)

		err = s.cmsProductService.DecreaseProductQuantity(ctx, msg.ProductID, msg.Quantity)
		if err != nil {
			s.logger.Error("Error decrease product quantity", zap.Error(err))
			continue
		}
		err = s.kafkaSubscribe.CommitMessages(context.Background(), m)
		if err != nil {
			s.logger.Error("Error commit message", zap.Error(err))
			continue
		}
	}
}
