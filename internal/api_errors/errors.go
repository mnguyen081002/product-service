package api_errors

import "net/http"

var (
	ErrInternalServerError = "10000"
	ErrUnauthorizedAccess  = "10001"
	ErrInvalidUserID       = "10002"
	ErrValidation          = "10003"
	ErrDeleteFailed        = "10004"

	ErrCreateProduct              = "40000"
	ErrProductNotFound            = "40001"
	ErrQuantityMustHigherThanZero = "40002"
	ErrQuantityNotEnough          = "40003"
	ErrIdNotFound                 = "40004"
	ErrProductAttributesNotFound  = "40005"
	ErrTierVariationNotFound      = "40006"
	ErrInvalidProductID           = "40007"
	ErrCategoryNotFound           = "40008"
	ErrInvalidCategoryName        = "40009"
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
	ErrIdNotFound:                 {"Invalid id", http.StatusBadRequest},
	ErrProductAttributesNotFound:  {"Product Attributes not found", http.StatusBadRequest},
	ErrTierVariationNotFound:      {"Tier Variation not found", http.StatusBadRequest},
	ErrDeleteFailed:               {"Delete Failed", http.StatusInternalServerError},
	ErrInvalidProductID:           {"Invalid Product ID", http.StatusBadRequest},
	ErrCategoryNotFound:           {"Category Not Found", http.StatusNotFound},
	ErrInvalidCategoryName:        {"Invalid Category Name", http.StatusBadRequest},
}
