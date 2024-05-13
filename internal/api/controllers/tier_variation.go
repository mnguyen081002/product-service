package controller

import (
	"net/http"
	"productservice/internal/api/request"
	"productservice/internal/domain"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TierVariationController struct {
	tierVariationService domain.TierVariationService
	// cmsProductAttributesProducer producer.CmsProductAttributeProducer
	logger *zap.Logger
}

func NewTierVariationController(tierVariationService domain.TierVariationService, logger *zap.Logger) *TierVariationController {
	controller := &TierVariationController{
		tierVariationService: tierVariationService,
		logger:               logger,
		// cmsProductAttributesProducer: cmsProductAttributesProducer,
	}
	return controller
}

func (controller *TierVariationController) UpdateTierVariation(ctx *gin.Context) {
	var request request.UpdateTierVariation
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ResponseValidationError(ctx, err)
		return
	}
	id := ctx.Param("product_id")

	err = controller.tierVariationService.UpdateTierVariationOptions(ctx.Request.Context(), request, id)
	if err != nil {
		ResponseError(ctx, err)
		return
	}

	Response(ctx, http.StatusOK, "success", nil)
}

func (controller *TierVariationController) DeleteTierVariationOptions(ctx *gin.Context) {
	var request request.DeleteTierVariation
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ResponseValidationError(ctx, err)
		return
	}

	err = controller.tierVariationService.DeleteTierVariationOptions(ctx.Request.Context(), request)
	if err != nil {
		ResponseError(ctx, err)
		return
	}

	Response(ctx, http.StatusOK, "success", nil)
}

// GetTierVariationByProductID
func (controller *TierVariationController) GetTierVariationByProductID(ctx *gin.Context) {
	id := ctx.Param("product_id")

	data, err := controller.tierVariationService.GetByProductID(ctx.Request.Context(), id)
	if err != nil {
		ResponseError(ctx, err)
		return
	}

	Response(ctx, http.StatusOK, "success", data)
}
