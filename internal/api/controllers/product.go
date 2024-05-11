package controller

import (
	"productservice/config"
	"productservice/internal/api/request"
	"productservice/internal/domain"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductController struct {
	productService domain.ProductService
	logger         *zap.Logger
	config         *config.Config
}

func NewProductController(productService domain.ProductService, logger *zap.Logger, config *config.Config) *ProductController {
	controller := &ProductController{
		productService: productService,
		logger:         logger,
		config:         config,
	}
	return controller
}

func (b *ProductController) GetProductList(c *gin.Context) {
	var req request.ClientListProductRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		ResponseValidationError(c, err)
		return
	}
	p, total, err := b.productService.GetProductList(c.Request.Context(), req)

	if err != nil {
		ResponseError(c, err)
		return
	}

	ResponseList(c, "success", total, p)
}
