package gormlib

import (
	"context"
	"fmt"
	"productservice/internal/api_errors"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"

	"github.com/pkg/errors"
)

type productModelRepository struct {
}

func NewProductModelRepository() domain.ProductModelRepository {
	return productModelRepository{}
}

// Create
func (u productModelRepository) Create(db *infrastructure.Database, ctx context.Context, productModel *models.ProductModel) (res *models.ProductModel, err error) {
	if err := db.RDBMS.WithContext(ctx).Create(&productModel).Error; err != nil {
		return nil, errors.Cause(err)
	}

	return productModel, nil
}

// BulkCreate
func (u productModelRepository) BulkCreate(db *infrastructure.Database, ctx context.Context, productModels []*models.ProductModel) (res []*models.ProductModel, err error) {

	if err := db.RDBMS.WithContext(ctx).Create(&productModels).Error; err != nil {
		return nil, errors.Cause(err)
	}

	return productModels, nil
}

// Update
func (u productModelRepository) Update(db *infrastructure.Database, ctx context.Context, id string, updates map[string]interface{}) (err error) {
	if err := db.RDBMS.WithContext(ctx).Model(&models.ProductModel{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// CountWithCondition
func (u productModelRepository) CountWithCondition(db *infrastructure.Database, ctx context.Context, condition map[string]interface{}) (count int64, err error) {

	if err := db.RDBMS.WithContext(ctx).Model(&models.ProductModel{}).Where(condition).Count(&count).Error; err != nil {
		return 0, errors.WithStack(err)
	}

	return count, nil
}

// BulkDeleteWithCondition
func (u productModelRepository) BulkDeleteByProductIdAndItemIndex(db *infrastructure.Database, ctx context.Context, productID string, firstChar string, optionID int) (err error) {
	likeCondition := "%" + firstChar
	if optionID == 0 {
		likeCondition = firstChar + "%"
	}

	fmt.Println(optionID == 0, likeCondition)

	pm := db.RDBMS.WithContext(ctx).Debug().Where("product_id = ?", productID).Where("item_index LIKE ?", likeCondition).Delete(&models.ProductModel{})

	if pm.Error != nil {
		return errors.WithStack(err)
	}

	if pm.RowsAffected == 0 {
		return errors.New(api_errors.ErrDeleteFailed)
	}

	return nil
}

// GetListByProductId
func (u productModelRepository) GetListByProductId(db *infrastructure.Database, ctx context.Context, productID string) (res []*models.ProductModel, err error) {
	if err := db.RDBMS.WithContext(ctx).Where("product_id = ?", productID).Find(&res).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}
