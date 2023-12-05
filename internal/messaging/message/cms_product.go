package message

type DecreaseProductQuantity struct {
	ProductID string `json:"product_id"`
	Quantity  int64  `json:"quantity"`
	UserID    string `json:"user_id"`
}
