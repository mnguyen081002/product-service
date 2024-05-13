package controller

import (
	"productservice/internal/domain"

	"go.uber.org/zap"
)

type ProductModelsController struct {
	productModelService domain.ProductModelService
	// cmsProductAttributesProducer producer.CmsProductAttributeProducer
	logger *zap.Logger
}

func NewProductModelsController(authService domain.ProductModelService, logger *zap.Logger) *ProductModelsController {
	controller := &ProductModelsController{
		productModelService: authService,
		logger:              logger,
		// cmsProductAttributesProducer: cmsProductAttributesProducer,
	}
	return controller
}
