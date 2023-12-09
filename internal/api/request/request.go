package request

type PageOptions struct {
	Page    int     `form:"page" json:"page"`
	Limit   int     `form:"limit" json:"limit"`
	Sort    string  `form:"sort" json:"sort"`
	Search  *string `form:"search" json:"search"`
	IsCount bool    `form:"is_count" json:"is_count"`
}
