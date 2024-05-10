package models

import (
	"productservice/internal/constants"

	uuid "github.com/satori/go.uuid"
)

type OrderItem struct {
	BaseModel      `json:",inline"`
	UserID         string                    `json:"user_id"`
	OrderID        uuid.UUID                 `json:"order_id"`
	ProductModelID uuid.UUID                 `json:"product_model_id"`
	Quantity       int64                     `json:"quantity"`
	Price          int64                     `json:"price"`
	Status         constants.OrderItemStatus `json:"status"`
}
