package request

import uuid "github.com/satori/go.uuid"

type OrderItem struct {
	ProductModelID uuid.UUID `json:"product_model_id"`
	Quantity       int64     `json:"quantity"`
}

type CreateOrderRequest struct {
	Items []OrderItem `json:"items"`
}
