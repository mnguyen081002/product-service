package controller

import (
	"net/http"
	"productservice/internal/domain"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductModelsController struct {
	productModels domain.ProductModelService
	// cmsProductAttributesProducer producer.CmsProductAttributeProducer
	logger *zap.Logger
}

func NewProductModelsController(authService domain.ProductModelService, logger *zap.Logger) *ProductModelsController {
	controller := &ProductModelsController{
		productModels: authService,
		logger:        logger,
		// cmsProductAttributesProducer: cmsProductAttributesProducer,
	}
	return controller
}

// GetListByProductId
func (controller *ProductModelsController) GetListByProductId(ctx *gin.Context) {
	productID := ctx.Param("product_id")

	data, err := controller.productModels.GetListByProductId(ctx.Request.Context(), productID)
	if err != nil {
		ResponseError(ctx, err)
		return
	}

	Response(ctx, http.StatusOK, "success", data)
}
