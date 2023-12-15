package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type ProductModels struct {
	BaseModel `bson:",inline"`
	ProductID uuid.UUID    `json:"product_id" gorm:"column:product_id;type:uuid;not null" bson:"-"`
	Models    ProductModel `json:"models" gorm:"column:model;type:jsonb" bson:"-"`
}

type ProductModel struct {
	Name                string `json:"name" gorm:"column:name;type:varchar(255);not null" bson:"-"`
	TierIndex           string `json:"tier_index" gorm:"column:item_index;type:varchar(255);not null" bson:"-"`
	Sold                int64  `json:"sold" gorm:"column:sold;type:bigint" bson:"-"`
	Price               int64  `json:"price" gorm:"column:price;type:bigint" bson:"-"`
	PriceBeforeDiscount int64  `json:"price_before_discount" gorm:"column:price_before_discount;type:bigint" bson:"-"`
	Stock               int64  `json:"stock" gorm:"column:stock;type:bigint" bson:"-"`
}

type ArrayProductModel []ProductModel

// Value Marshal
func (a ArrayProductModel) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *ArrayProductModel) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
