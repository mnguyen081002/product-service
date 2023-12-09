package gormlib

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
	"productservice/internal/api/request"
	"productservice/internal/api_errors"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
	"productservice/internal/utils"
)

type categoryRepository struct {
	logger *zap.Logger
}

func NewCategoryRepository() domain.CategoryRepository {
	return categoryRepository{}
}

func (u categoryRepository) Create(db *infrastructure.Database, ctx context.Context, cat *models.Category) (res *models.Category, err error) {
	if err := db.RDBMS.WithContext(ctx).Create(&cat).Error; err != nil {
		return nil, errors.Cause(err)
	}

	return cat, nil
}

func (u categoryRepository) Update(db *infrastructure.Database, ctx context.Context, update map[string]interface{}, id uuid.UUID) (*models.Category, error) {
	res := &models.Category{}
	if err := db.RDBMS.WithContext(ctx).Clauses(clause.Returning{}).Model(res).Where("id = ?", id).Updates(update).Error; err != nil {
		return nil, errors.Cause(err)
	}
	fmt.Println(res)
	return res, nil
}

func (u categoryRepository) UpdateNoSub(db *infrastructure.Database, ctx context.Context, noSub bool, id uuid.UUID) (err error) {
	if err = db.RDBMS.WithContext(ctx).Model(&models.Category{}).Where("id = ?", id).Update("no_sub", noSub).Error; err != nil {
		return errors.Cause(err)
	}
	return nil
}

func (u categoryRepository) CountSubCategoryOfCategory(db *infrastructure.Database, ctx context.Context, id uuid.UUID) (total int64, err error) {
	if err = db.RDBMS.WithContext(ctx).Model(&models.Category{}).Where("parent_id = ?", id).Count(&total).Error; err != nil {
		return 0, errors.Cause(err)
	}
	return total, nil
}

func (u categoryRepository) FindByID(db *infrastructure.Database, ctx context.Context, id uuid.UUID) (cat *models.Category, err error) {
	if err = db.RDBMS.WithContext(ctx).Model(cat).Where("id = ?", id).First(&cat).Error; err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New(api_errors.ErrCategoryNotFound)
		}
		return nil, errors.Cause(err)
	}
	return cat, nil
}

func (u categoryRepository) ListCategories(db *infrastructure.Database, ctx context.Context, options request.ListCategoryRequest) (categories []*models.Category, total int64, err error) {
	qb := db.RDBMS.WithContext(ctx).Model(&models.Category{})

	if options.IsParent {
		qb = qb.Where("parent_id IS NULL")
	}

	if options.ParentID != nil {
		qb = qb.Where("parent_id = ?", *options.ParentID)
	}

	if options.Search != nil {
		qb = qb.Where("name ILIKE ?", "%"+*options.Search+"%")
	}

	qpb := GormQueryPagination(qb, options.PageOptions, &categories)
	if options.IsCount {
		qpb.Count(&total)
	}

	if err = qpb.Error(); err != nil {
		return nil, 0, errors.Cause(err)
	}
	return categories, total, nil
}
