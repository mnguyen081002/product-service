package gormlib

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"productservice/internal/api/request"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
)

type ratingRepository struct {
	logger *zap.Logger
}

func NewRatingRepository() domain.RatingRepository {
	return ratingRepository{}
}

func (s ratingRepository) Create(db *infrastructure.Database, ctx context.Context, m *models.Rating) (r *models.Rating, err error) {
	if err := db.RDBMS.WithContext(ctx).Create(&m).Error; err != nil {
		return nil, errors.Cause(err)
	}

	return m, nil
}

func (s ratingRepository) ListByProductID(db *infrastructure.Database, ctx context.Context, options request.ListRatingByProductID) (ratings []*models.Rating, total int64, err error) {
	dbQuery := db.RDBMS.WithContext(ctx).Model(&models.Rating{}).Where("product_id = ?", options.ProductID)

	if options.Rate != nil {
		dbQuery = dbQuery.Where("rating = ?", *options.Rate)
	}

	qb := GormQueryPagination(dbQuery, options.PageOptions, &ratings)

	if options.IsCount {
		qb = qb.Count(&total)
	}

	err = qb.Error()

	if err != nil {
		return nil, 0, errors.Cause(err)
	}
	return ratings, total, nil
}
