package request

type ClientListProductRequest struct {
	PageOptions
	Price *int64 `form:"price"`
}
