package request

type CreateRating struct {
	ProductID string   `json:"product_id" binding:"required"`
	Rate      int8     `json:"rate" binding:"required" validate:"min=1,max=5"`
	Comment   string   `json:"comment"`
	Images    []string `json:"images"`
}

type ListRatingByProductID struct {
	PageOptions
	Rate      *int8  `form:"rate" json:"rate"`
	ProductID string `form:"product_id" json:"product_id" binding:"required,uuid"`
}
