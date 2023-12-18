package controller

import (
	"net/http"
	"productservice/internal/api/request"
	"productservice/internal/domain"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TierVariationController struct {
	TierVariationService domain.TierVariationService
	// cmsProductAttributesProducer producer.CmsProductAttributeProducer
	logger *zap.Logger
}

func NewTierVariationController(authService domain.TierVariationService, logger *zap.Logger) *TierVariationController {
	controller := &TierVariationController{
		TierVariationService: authService,
		logger:               logger,
		// cmsProductAttributesProducer: cmsProductAttributesProducer,
	}
	return controller
}

func (controller *TierVariationController) CreaTierVariation(ctx *gin.Context) {
	var request request.TierVariationCreate
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ResponseValidationError(ctx, err)
		return
	}

	data, err := controller.TierVariationService.CreateTierVariation(ctx.Request.Context(), request)
	if err != nil {
		ResponseError(ctx, err)
		return
	}

	Response(ctx, http.StatusOK, "success", data)
}

func (controller *TierVariationController) UpdateTierVariation(ctx *gin.Context) {
	var request request.TierVariationUpdate
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ResponseValidationError(ctx, err)
		return
	}
	id := ctx.Param("product_id")

	err = controller.TierVariationService.UpdateTierVariationOptions(ctx.Request.Context(), request, id)
	if err != nil {
		ResponseError(ctx, err)
		return
	}

	Response(ctx, http.StatusOK, "success", nil)
}

func (controller *TierVariationController) DeleteTierVariationOptions(ctx *gin.Context) {
	var request request.TierVariationDelete
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ResponseValidationError(ctx, err)
		return
	}
	id := ctx.Param("id")

	err = controller.TierVariationService.DeleteTierVariationOptions(ctx.Request.Context(), request, id)
	if err != nil {
		ResponseError(ctx, err)
		return
	}

	Response(ctx, http.StatusOK, "success", nil)
}

// GetTierVariationByProductID
func (controller *TierVariationController) GetTierVariationByProductID(ctx *gin.Context) {
	id := ctx.Param("product_id")

	data, err := controller.TierVariationService.GetByProductID(ctx.Request.Context(), id)
	if err != nil {
		ResponseError(ctx, err)
		return
	}

	Response(ctx, http.StatusOK, "success", data)
}
