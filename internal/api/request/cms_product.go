package request

type CmsCreateProductRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Price       int64    `json:"price" binding:"required"`
	Quantity    int64    `json:"quantity" binding:"required"`
	Images      []string `json:"images" binding:"required"`
	CategoryID  string   `json:"category_id" binding:"required,uuid"`
}

type CmsListProductRequest struct {
	PageOptions
	Price *int64 `form:"price"`
}

type CmsUpdateProductQuantityRequest struct {
	Quantity int64 `json:"quantity" binding:"required" validate:"min=1"`
}
