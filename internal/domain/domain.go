package domain

type QueryOptions struct {
	Page    int     `json:"page"`
	Limit   int     `json:"limit"`
	Sort    string  `json:"sort"`
	Search  *string `json:"search"`
	IsCount bool    `json:"is_count"`
}
