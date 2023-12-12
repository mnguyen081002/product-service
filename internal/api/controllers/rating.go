package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"productservice/config"
	"productservice/internal/api/request"
	"productservice/internal/domain"
)

type RatingController struct {
	cmsRatingService domain.RatingService
	logger           *zap.Logger
	config           *config.Config
}

func NewRatingController(ratingService domain.RatingService, logger *zap.Logger, config *config.Config) *RatingController {
	controller := &RatingController{
		cmsRatingService: ratingService,
		logger:           logger,
		config:           config,
	}
	return controller
}

func (b *RatingController) CreateRating(c *gin.Context) {
	var req request.CreateRating
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, err)
		return
	}

	p, err := b.cmsRatingService.CreateRating(c.Request.Context(), req)

	if err != nil {
		ResponseError(c, err)
		return
	}
	Response(c, http.StatusOK, "success", p.ID)
}

func (b *RatingController) ListRatingByProductID(c *gin.Context) {
	req := request.ListRatingByProductID{}
	if err := c.ShouldBindQuery(&req); err != nil {
		ResponseValidationError(c, err)
		return
	}

	ratings, total, err := b.cmsRatingService.ListRatingByProductID(c.Request.Context(), req)
	if err != nil {
		ResponseError(c, err)
		return
	}
	ResponseList(c, "success", total, ratings)
}
