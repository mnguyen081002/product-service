package gormlib

import (
	"context"
	"productservice/internal/api/request"
	"productservice/internal/api_errors"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
	"productservice/internal/utils"

	"github.com/pkg/errors"
)

type orderRepository struct {
}

func NewOrderRepository() domain.OrderRepository {
	return orderRepository{}
}

func (u orderRepository) Create(db *infrastructure.Database, ctx context.Context, order *models.Order) (res *models.Order, err error) {
	if err := db.RDBMS.WithContext(ctx).Create(&order).Error; err != nil {
		return nil, errors.Cause(err)
	}

	return order, nil
}

func (u orderRepository) Update(db *infrastructure.Database, ctx context.Context, order *models.Order, orderID string) (res *models.Order, err error) {
	if err := db.RDBMS.WithContext(ctx).Where("id = ?", orderID).Updates(&order).Error; err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New(api_errors.ErrOrderNotFound)
		}
		return nil, errors.Cause(err)
	}

	return order, nil
}

func (u orderRepository) GetList(db *infrastructure.Database, ctx context.Context, req request.GetListRequest, userId string) (res []*models.Order, err error) {

	dbQuery := db.RDBMS.WithContext(ctx).Model(&models.Order{})
	if req.Search != nil {
		dbQuery = dbQuery.Where("name ILIKE ?", "%"+*req.Search+"%")
	}

	dbQuery = dbQuery.Where("updater_id = ?", userId)

	err = GormQueryPagination(dbQuery, req.PageOptions, &res).Error()

	if err != nil {
		return nil, errors.Cause(err)
	}

	return res, nil

}

func (u orderRepository) Delete(db *infrastructure.Database, ctx context.Context, orderID string) (err error) {
	pm := db.RDBMS.WithContext(ctx).Debug().Where("id = ?", orderID).Delete(&models.Order{})

	if pm.Error != nil {
		if utils.ErrNoRows(pm.Error) {
			return errors.New(api_errors.ErrOrderNotFound)
		}
		return errors.WithStack(pm.Error)

	}

	if pm.RowsAffected == 0 {
		return errors.New(api_errors.ErrDeleteOrderFailed)
	}

	return nil
}
