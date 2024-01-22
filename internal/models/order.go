package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"
)

type Order struct {
	BaseModel    `bson:",inline"`
	Product      Product      `json:"product" gorm:"column:product;type:jsonb;not null" bson:"-"`
	ProductModel ProductModel `json:"product_model" gorm:"product_model:details;type:jsonb;not null" bson:"-"`
	Status       string       `json:"status" gorm:"column:status;type:varchar(255);not null" bson:"-"`
	Quantity     int64        `json:"quantity" gorm:"column:quantity;type:bigint;not null" bson:"-"`
	TotalPrice   int64        `json:"total_price" gorm:"column:total_price;type:bigint" bson:"-"`
	Reason       string       `json:"reason,omitempty" gorm:"column:reasion;type:varchar(255)" bson:"-"`
}

// Value Marshal
func (a Order) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *Order) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

// Value Marshal Product
func (a Product) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal Product
func (a *Product) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

// Value Marshal ProductModel
func (a ProductModel) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal ProductModel
func (a *ProductModel) Scan(value interface{}) error {

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
