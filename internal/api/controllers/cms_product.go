package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"productservice/config"
	"productservice/internal/api/request"
	"productservice/internal/domain"
	"productservice/internal/messaging/message"
	"productservice/internal/messaging/producer"
	"productservice/internal/utils"
)

type CmsProductController struct {
	cmsProductService  domain.CmsProductService
	cmsProductProducer producer.CmsProductProducer
	logger             *zap.Logger
	config             *config.Config
}

func NewCmsProductController(authService domain.CmsProductService, logger *zap.Logger, cmsProductProducer producer.CmsProductProducer, config *config.Config) *CmsProductController {
	controller := &CmsProductController{
		cmsProductService:  authService,
		logger:             logger,
		cmsProductProducer: cmsProductProducer,
		config:             config,
	}
	return controller
}

func (b *CmsProductController) CreateProduct(c *gin.Context) {
	var req request.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, err)
		return
	}

	p, err := b.cmsProductService.CreateProduct(c.Request.Context(), req)

	if err != nil {
		ResponseError(c, err)
		return
	}
	Response(c, http.StatusOK, "success", p.ID)
}

func (b *CmsProductController) GetProductById(c *gin.Context) {
	id := c.Param("id")

	p, err := b.cmsProductService.GetProductById(c.Request.Context(), id)

	if err != nil {
		ResponseError(c, err)
		return
	}
	Response(c, http.StatusOK, "success", p)
}

func (b *CmsProductController) ListProduct(c *gin.Context) {
	var req request.ListProductRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		ResponseError(c, err)
		return
	}

	products, total, err := b.cmsProductService.ListProduct(c.Request.Context(), req)
	if err != nil {
		return
	}

	if err != nil {
		ResponseError(c, err)
		return
	}
	ResponseList(c, "success", total, products)
}

func (b *CmsProductController) DecreaseProductQuantity(c *gin.Context) {
	id := c.Param("id")
	var req request.UpdateProductQuantityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseValidationError(c, err)
		return
	}
	uid, err := utils.GetUserUUIDFromContext(c.Request.Context())
	if err != nil {
		ResponseError(c, err)
		return
	}

	if b.config.Kafka.Enable {
		err = b.cmsProductProducer.DecreaseProductQuantity(c.Request.Context(), message.DecreaseProductQuantity{
			ProductID: id,
			Quantity:  req.Quantity,
			UserID:    uid.String(),
		})
	} else {
		err = b.cmsProductService.DecreaseProductQuantity(c.Request.Context(), id, req.Quantity)
	}

	if err != nil {
		ResponseError(c, err)
		return
	}
	Response(c, http.StatusOK, "success", nil)
}

func (b *CmsProductController) DecreaseProductQuantityMutex(c *gin.Context) {
	id := c.Param("id")
	var req request.UpdateProductQuantityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseValidationError(c, err)
		return
	}

	err := b.cmsProductService.DecreaseProductQuantityMutex(c.Request.Context(), id, req.Quantity)

	if err != nil {
		ResponseError(c, err)
		return
	}
	Response(c, http.StatusOK, "success", nil)
}
