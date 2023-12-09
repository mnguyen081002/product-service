package controller

import (
	"net/http"
	"productservice/internal/api/request"
	"productservice/internal/domain"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CmsProductAttributesController struct {
	cmsProductAttributesService domain.CmsProductAttributesService
	// cmsProductAttributesProducer producer.CmsProductAttributeProducer
	logger *zap.Logger
}

func NewCmsProductAttributesController(authService domain.CmsProductAttributesService, logger *zap.Logger) *CmsProductAttributesController {
	controller := &CmsProductAttributesController{
		cmsProductAttributesService: authService,
		logger:                      logger,
		// cmsProductAttributesProducer: cmsProductAttributesProducer,
	}
	return controller
}

func (controller *CmsProductAttributesController) CreateProductAttribute(ctx *gin.Context) {
	var request request.ProductAttributesCreate
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ResponseValidationError(ctx, err)
		return
	}

	data, err := controller.cmsProductAttributesService.CreateProductAttributes(ctx.Request.Context(), request)
	if err != nil {
		ResponseError(ctx, err)
		return
	}

	Response(ctx, http.StatusOK, "success", data)
}
