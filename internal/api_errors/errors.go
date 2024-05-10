package api_errors

type ExternalError interface {
	Error() string
	Status() int
	Code() string
}

type externalError struct {
	message string
	status  int
	code    string
}

func NewExternalError(message string, status int, code string) ExternalError {
	return externalError{
		message: message,
		status:  status,
		code:    code,
	}
}

func (e externalError) Error() string {
	return e.message
}

func (e externalError) Status() int {
	return e.status
}

func (e externalError) Code() string {
	return e.code
}

var (
	ErrInternalServerError        = "10000"
	ErrUnauthorizedAccess         = "10001"
	ErrCreateProduct              = "10002"
	ErrInvalidUserID              = "10003"
	ErrValidation                 = "10004"
	ErrProductNotFound            = "10005"
	ErrQuantityMustHigherThanZero = "10006"
	ErrQuantityNotEnough          = "10007"
	ErrIdNotFound                 = "10008"
	ErrProductAttributesNotFound  = "10009"
	ErrTierVariationNotFound      = "10010"
	ErrDeleteFailed               = "10011"
	ErrInvalidProductID           = "10012"
	ErrCategoryNotFound           = "10013"
	ErrInvalidCategoryName        = "10014"
	ErrInvalidUUID                = "10015"
	ErrProductAttributeNotFound   = "10016"

	ErrCartQuantityMustHigherThanZero = "70000"
)

var MapErrCodeMsg = map[string]string{
	ErrInternalServerError:        "Internal Server Error",
	ErrUnauthorizedAccess:         "Unauthorized Access",
	ErrCreateProduct:              "Error Create Product",
	ErrInvalidUserID:              "Invalid User ID",
	ErrValidation:                 "Validation Error",
	ErrProductNotFound:            "Product Not Found",
	ErrQuantityMustHigherThanZero: "Quantity must higher than zero",
	ErrQuantityNotEnough:          "Quantity not enough",
	ErrIdNotFound:                 "Invalid id",
	ErrProductAttributesNotFound:  "Product Attributes not found",
	ErrTierVariationNotFound:      "Tier Variation not found",
	ErrDeleteFailed:               "Delete Failed",
	ErrInvalidProductID:           "Invalid Product ID",
	ErrCategoryNotFound:           "Category Not Found",
	ErrInvalidCategoryName:        "Invalid Category Name",
	ErrInvalidUUID:                "Invalid UUID",
	ErrProductAttributeNotFound:   "Product Attribute Not Found",
}
