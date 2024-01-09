package service

import (
	"context"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"productservice/config"
	"productservice/internal/api/request"
	"productservice/internal/api_errors"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
	"productservice/internal/repository"
	"productservice/internal/utils"
	"sync"
)

type ratingService struct {
	db            infrastructure.Database
	dbTransaction infrastructure.DatabaseTransaction
	ratingService domain.RatingService
	ufw           *repository.UnitOfWork
	config        *config.Config
	logger        *zap.Logger
	mu            sync.Mutex
}

func NewRatingService(
	db infrastructure.Database,
	dbTransaction infrastructure.DatabaseTransaction,
	ufw *repository.UnitOfWork,
	config *config.Config,
	logger *zap.Logger,
) domain.RatingService {
	return &ratingService{
		db:            db,
		dbTransaction: dbTransaction,
		ufw:           ufw,
		config:        config,
		logger:        logger,
	}
}

func (s *ratingService) CreateRating(ctx context.Context, req request.CreateRating) (r *models.Rating, err error) {
	raterId, err := utils.GetUserUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	uProductId, err := uuid.FromString(req.ProductID)
	if err != nil {
		return nil, errors.New(api_errors.ErrInvalidProductID)
	}
	err = s.dbTransaction.WithTransaction(func(tx *infrastructure.Database) error {
		r, err = s.ufw.RatingRepository.Create(tx, ctx, &models.Rating{
			ProductID: &uProductId,
			RaterID:   &raterId,
			Rating:    req.Rate,
			Comment:   req.Comment,
			Images:    req.Images,
		})
		if err != nil {
			return err
		}

		p, err := s.ufw.ProductRepository.GetById(tx, ctx, req.ProductID)
		if err != nil {
			return err
		}

		p.MedRating = (p.MedRating*float64(p.RatingCount) + float64(req.Rate)) / float64(p.RatingCount+1)
		p.RatingCount += 1

		if err := s.ufw.ProductRepository.Update(tx, ctx, req.ProductID, map[string]interface{}{
			"rating_count": p.RatingCount,
			"med_rating":   p.MedRating,
		}); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *ratingService) ListRatingByProductID(ctx context.Context, params request.ListRatingByProductID) (ratings []*models.Rating, total int64, err error) {
	return s.ufw.RatingRepository.ListByProductID(&s.db, ctx, params)
}
