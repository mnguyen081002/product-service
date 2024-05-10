package controller

import (
	"net/http"
	"productservice/config"
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

func (b *ProductController) GetProductById(c *gin.Context) {
	id := c.Param("id")

	p, err := b.productService.GetProductByID(c.Request.Context(), id)

	if err != nil {
		ResponseError(c, err)
		return
	}

	Response(c, http.StatusOK, "success", p)
}
