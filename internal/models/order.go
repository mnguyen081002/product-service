package models

import (
	"productservice/internal/constants"
)

type Order struct {
	BaseModel  `json:",inline"`
	UserID     string                `json:"user_id"`
	Items      []*OrderItem          `json:"items"`
	TotalPrice int64                 `json:"total_price"`
	Status     constants.OrderStatus `json:"status"`
}
