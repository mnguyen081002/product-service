package service

import (
	"context"
	"productservice/config"
	"productservice/internal/api/request"
	"productservice/internal/api_errors"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
	"productservice/internal/repository"
	"sync"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type cmsCategoryService struct {
	db                 infrastructure.Database
	dbTransaction      infrastructure.DatabaseTransaction
	cmsCategoryService domain.CmsCategoryService
	ufw                *repository.UnitOfWork
	config             *config.Config
	logger             *zap.Logger
	mu                 sync.Mutex
}

func NewCmsCategoryService(
	db infrastructure.Database,
	dbTransaction infrastructure.DatabaseTransaction,
	ufw *repository.UnitOfWork,
	config *config.Config,
	logger *zap.Logger,
) domain.CmsCategoryService {
	return &cmsCategoryService{
		db:            db,
		dbTransaction: dbTransaction,
		ufw:           ufw,
		config:        config,
		logger:        logger,
	}
}

func (a *cmsCategoryService) CreateCategory(ctx context.Context, req request.CreateCategoryRequest) (cat *models.Category, err error) {
	if req.Name == "" {
		return nil, errors.New(api_errors.ErrInvalidCategoryName)
	}
	err = a.dbTransaction.WithTransaction(func(tx *infrastructure.Database) error {
		cat, err = a.ufw.CategoryRepository.Create(tx, ctx, &models.Category{
			Name:     req.Name,
			NoSub:    true,
			ParentID: req.ParentID,
		})
		if err != nil {
			return err
		}
		if req.ParentID != nil {
			err = a.ufw.CategoryRepository.UpdateNoSub(&a.db, ctx, false, *req.ParentID)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return cat, err
}

func (a *cmsCategoryService) UpdateCategory(ctx context.Context, input request.UpdateCategoryRequest, id uuid.UUID) (category *models.Category, err error) {
	if input.Name == "" {
		return nil, errors.New(api_errors.ErrInvalidCategoryName)
	}

	err = a.dbTransaction.WithTransaction(func(tx *infrastructure.Database) error {
		foundCat, err := a.ufw.CategoryRepository.FindByID(tx, ctx, id)
		if err != nil {
			return err
		}
		category, err = a.ufw.CategoryRepository.Update(tx, ctx, map[string]interface{}{
			"name":      input.Name,
			"parent_id": input.ParentID,
		}, id)
		if err != nil {
			return err
		}
		if input.ParentID == nil && foundCat.ParentID != nil {
			count, err := a.ufw.CategoryRepository.CountSubCategoryOfCategory(tx, ctx, *foundCat.ParentID)
			if err != nil {
				return err
			}
			// if count = 0 => category doesn't have sub category
			if count == 0 {
				err = a.ufw.CategoryRepository.UpdateNoSub(tx, ctx, true, *foundCat.ParentID)
				if err != nil {
					return err
				}
			}
		} else if foundCat.ParentID != nil && *input.ParentID != *foundCat.ParentID {
			err = a.ufw.CategoryRepository.UpdateNoSub(tx, ctx, false, *foundCat.ParentID)
			if err != nil {
				return err
			}
		}

		return nil
	})
	return category, err
}

func (a *cmsCategoryService) ListCategories(ctx context.Context, options request.ListCategoryRequest) (categories []*models.Category, total int64, err error) {
	categories, total, err = a.ufw.CategoryRepository.ListCategories(&a.db, ctx, options)
	return categories, total, err
}
