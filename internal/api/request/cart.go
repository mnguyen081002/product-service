package request

type AddItemToCartRequest struct {
	ProductModelID string `json:"product_model_id"`
	Quantity       int    `json:"quantity"`
}
