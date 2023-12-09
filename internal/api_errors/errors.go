package api_errors

import "net/http"

var (
	ErrInternalServerError        = "10000"
	ErrUnauthorizedAccess         = "10001"
	ErrCreateProduct              = "10002"
	ErrInvalidUserID              = "10003"
	ErrValidation                 = "10004"
	ErrProductNotFound            = "10005"
	ErrQuantityMustHigherThanZero = "10006"
	ErrQuantityNotEnough          = "10007"
	ErrInvalidProductID           = "10008"
	ErrCategoryNotFound           = "10009"
	ErrInvalidCategoryName        = "10010"
)

type MessageAndStatus struct {
	Message string
	Status  int
}

var MapErrorCodeMessage = map[string]MessageAndStatus{
	ErrInternalServerError:        {"Internal Server Error", http.StatusInternalServerError},
	ErrUnauthorizedAccess:         {"Unauthorized Access", http.StatusUnauthorized},
	ErrCreateProduct:              {"Error Create Product", http.StatusInternalServerError},
	ErrInvalidUserID:              {"Invalid User ID", http.StatusBadRequest},
	ErrValidation:                 {"Validation Error", http.StatusBadRequest},
	ErrProductNotFound:            {"Product Not Found", http.StatusNotFound},
	ErrQuantityMustHigherThanZero: {"Quantity must higher than zero", http.StatusBadRequest},
	ErrQuantityNotEnough:          {"Quantity not enough", http.StatusBadRequest},
	ErrInvalidProductID:           {"Invalid Product ID", http.StatusBadRequest},
	ErrCategoryNotFound:           {"Category Not Found", http.StatusNotFound},
	ErrInvalidCategoryName:        {"Invalid Category Name", http.StatusBadRequest},
}
