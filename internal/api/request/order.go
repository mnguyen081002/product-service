package request

type CreateOrderRequest struct {
	ProductID      string `json:"product_id" gorm:"column:product_id;type:uuid;not null" bson:"-"`
	ProductModelID string `json:"product_model_id" gorm:"column:product_model_id;type:uuid;not null" bson:"-"`
	Quantity       int64  `json:"quantity" gorm:"column:quantity;type:bigint;not null" bson:"-"`
	TotalPrice     int64  `json:"total_price" gorm:"column:total_price;type:bigint" bson:"-"`
}

type UpdateOrderRequest struct {
	ProductID      string `json:"product_id" gorm:"column:product_id;type:uuid;not null" bson:"-"`
	ProductModelID string `json:"product_model_id" gorm:"column:product_model_id;type:uuid;not null" bson:"-"`
	Quantity       int64  `json:"quantity" gorm:"column:quantity;type:bigint;not null" bson:"-"`
	TotalPrice     int64  `json:"total_price" gorm:"column:total_price;type:bigint" bson:"-"`
	Status         string `json:"status" gorm:"column:status;type:varchar(255)" bson:"-"`
	Reason         string `json:"reason,omitempty" gorm:"column:reasion;type:varchar(255)" bson:"-"`
}

type GetListRequest struct {
	PageOptions
}
