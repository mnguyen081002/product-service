package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"productservice/internal/api/request"
	"productservice/internal/domain"
	"productservice/internal/utils"
)

type CmsCategoryController struct {
	cmsCategoryService domain.CmsCategoryService
	logger             *zap.Logger
}

func NewCmsCategoryController(cmsCategoryService domain.CmsCategoryService, logger *zap.Logger) *CmsCategoryController {
	controller := &CmsCategoryController{
		cmsCategoryService: cmsCategoryService,
		logger:             logger,
	}
	return controller
}

func (b *CmsCategoryController) CreateCategory(c *gin.Context) {
	var req request.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, err)
		return
	}

	p, err := b.cmsCategoryService.CreateCategory(c.Request.Context(), req)

	if err != nil {
		ResponseError(c, err)
		return
	}
	Response(c, http.StatusOK, "success", p.ID)
}

func (b *CmsCategoryController) UpdateCategory(c *gin.Context) {
	paramID, err := utils.GetUUIDFromParam(c, "id")
	if err != nil {
		ResponseError(c, err)
		return
	}
	var req request.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, err)
		return
	}

	_, err = b.cmsCategoryService.UpdateCategory(c.Request.Context(), req, paramID)

	if err != nil {
		ResponseError(c, err)
		return
	}
	Response(c, http.StatusOK, "success", nil)
}

func (b *CmsCategoryController) ListCategories(c *gin.Context) {
	var req request.ListCategoryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		ResponseError(c, err)
		return
	}

	categories, total, err := b.cmsCategoryService.ListCategories(c.Request.Context(), req)

	if err != nil {
		ResponseError(c, err)
		return
	}

	ResponseList(c, "success", total, categories)
}
