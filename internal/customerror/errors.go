package customerror

import "errors"

const (
	INVALID_RECEIPT_ERROR   = "The receipt is invalid."
	RECEIPT_NOT_FOUND_ERROR = "No receipt found for that ID."
)

func NewInvalidReceiptError() error {
	return CustomError{
		error:        errors.New(INVALID_RECEIPT_ERROR),
		ErrorMessage: INVALID_RECEIPT_ERROR,
	}
}

func NewReceiptNotFoundError() error {
	return CustomError{
		error:        errors.New(RECEIPT_NOT_FOUND_ERROR),
		ErrorMessage: RECEIPT_NOT_FOUND_ERROR,
	}
}

type CustomError struct {
	error
	ErrorMessage string `json:"error"`
}
