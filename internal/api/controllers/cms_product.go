package controller

import (
	"net/http"
	"productservice/internal/api/request"
	"productservice/internal/domain"
	"productservice/internal/messaging/message"
	"productservice/internal/messaging/producer"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CmsProductController struct {
	cmsProductService  domain.CmsProductService
	cmsProductProducer producer.CmsProductProducer
	logger             *zap.Logger
}

func NewCmsProductController(authService domain.CmsProductService, logger *zap.Logger, cmsProductProducer producer.CmsProductProducer) *CmsProductController {
	controller := &CmsProductController{
		cmsProductService:  authService,
		logger:             logger,
		cmsProductProducer: cmsProductProducer,
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
		ResponseError(c, err)
		return
	}

	err := b.cmsProductProducer.PublishDecreaseProductQuantity(c.Request.Context(), message.DecreaseProductQuantity{
		ProductID: id,
		Quantity:  req.Quantity,
		UserID:    c.GetString("user_id"),
	})

	if err != nil {
		ResponseError(c, err)
		return
	}
	Response(c, http.StatusOK, "success", nil)
}
