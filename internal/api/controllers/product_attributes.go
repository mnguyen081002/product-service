package controller

import (
	"net/http"
	"productservice/internal/api/request"
	"productservice/internal/api_errors"
	"productservice/internal/domain"
	"productservice/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductAttributesController struct {
	cmsProductAttributesService domain.CmsProductAttributesService
	// cmsProductAttributesProducer producer.CmsProductAttributeProducer
	logger *zap.Logger
}

func NewCmsProductAttributesController(authService domain.CmsProductAttributesService, logger *zap.Logger) *ProductAttributesController {
	controller := &ProductAttributesController{
		cmsProductAttributesService: authService,
		logger:                      logger,
		// cmsProductAttributesProducer: cmsProductAttributesProducer,
	}
	return controller
}

func (controller *ProductAttributesController) CreateProductAttribute(ctx *gin.Context) {
	var request request.CmsCreateProductAttributesRequest
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

// UpdateProductAttribute
func (controller *ProductAttributesController) UpdateProductAttribute(ctx *gin.Context) {
	id := ctx.Param("id")

	if !utils.IsValidUUID(id) {
		ResponseError(ctx, api_errors.NewExternalError(api_errors.MapErrCodeMsg[api_errors.ErrInvalidUUID],
			http.StatusBadRequest, api_errors.ErrInvalidUUID))
		return
	}

	var request request.CmsUpdateProductAttributesRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ResponseValidationError(ctx, err)
		return
	}

	// check if product attribute exist
	_, err = controller.cmsProductAttributesService.GetProductAttributesById(ctx.Request.Context(), id)
	if err != nil {
		ResponseError(ctx, err)
		return
	}

	err = controller.cmsProductAttributesService.UpdateProductAttributes(ctx.Request.Context(), id, request)
	if err != nil {
		ResponseError(ctx, err)
		return
	}

	Response(ctx, http.StatusOK, "success", nil)
}
