package controller

import (
	"net/http"
	"productservice/config"
	"productservice/internal/domain"
	"productservice/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CartController struct {
	cartService domain.CartService
	logger      *zap.Logger
	config      *config.Config
}

func NewCartController(productService domain.CartService, logger *zap.Logger, config *config.Config) *CartController {
	controller := &CartController{
		cartService: productService,
		logger:      logger,
		config:      config,
	}
	return controller
}

func (b *CartController) GetUserCart(c *gin.Context) {
	uid, err := utils.GetUserUUIDFromContext(c.Request.Context())
	if err != nil {
		ResponseError(c, err)
		return
	}

	p, err := b.cartService.GetUserCart(c.Request.Context(), uid)

	if err != nil {
		ResponseError(c, err)
		return
	}

	Response(c, http.StatusOK, "success", p)
}
