package models

import uuid "github.com/satori/go.uuid"

type ProductModel struct {
	BaseModel           `bson:",inline"`
	ProductID           uuid.UUID `json:"product_id" gorm:"column:product_id;type:uuid;not null" bson:"-"`
	ItemIndex           int       `json:"item_index" gorm:"column:item_index;type:bigint;not null" bson:"-"`
	Sold                int64     `json:"sold" gorm:"column:sold;type:bigint;not null" bson:"-"`
	Price               int64     `json:"price" gorm:"column:price;type:bigint;not null" bson:"-"`
	PriceBeforeDiscount int64     `json:"price_before_discount" gorm:"column:price_before_discount;type:bigint;not null" bson:"-"`
	Stock               int64     `json:"stock" gorm:"column:stock;type:bigint;not null" bson:"-"`
}
