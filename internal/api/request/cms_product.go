package request

type CreateProductRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Price       int64    `json:"price" binding:"required"`
	Quantity    int64    `json:"quantity" binding:"required"`
	Images      []string `json:"images" binding:"required"`
}

type ListProductRequest struct {
	PageOptions
	Price *int64 `form:"price"`
}

type UpdateProductQuantityRequest struct {
	Quantity int64 `json:"quantity" binding:"required" validate:"min=1"`
}
