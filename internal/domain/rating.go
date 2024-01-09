package domain

import (
	"context"
	"productservice/internal/api/request"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
)

type RatingRepository interface {
	Create(db *infrastructure.Database, ctx context.Context, m *models.Rating) (r *models.Rating, err error)
	ListByProductID(db *infrastructure.Database, ctx context.Context, options request.ListRatingByProductID) (ratings []*models.Rating, total int64, err error)
}

type RatingService interface {
	CreateRating(ctx context.Context, req request.CreateRating) (r *models.Rating, err error)
	ListRatingByProductID(ctx context.Context, params request.ListRatingByProductID) (ratings []*models.Rating, total int64, err error)
}
