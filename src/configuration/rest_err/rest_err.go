package rest_err

import (
	"net/http"

)

// RestErr represents a standard error response returned by the API
type RestErr struct {
	// Human-readable error message
	// example: invalid user ID
	Message string `json:"message"`

	// Machine-readable error identifier
	// example: bad_request
	Err string `json:"error"`

	// HTTP status code
	// example: 400
	Code int `json:"code"`

	// List of specific causes or validation errors (optional)
	Causes []Causes `json:"causes"`
}

// Causes represents a specific field-level error detail
type Causes struct {
	// Name of the field that caused the error
	// example: email
	Field string `json:"field"`

	// Description of the validation error
	// example: must be a valid email address
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int, causes []Causes) *RestErr{
	return &RestErr{
		Message: message,
		Err: err,
		Code: code,
		Causes: causes,
	}
}

func NewBadRequestError(message string) *RestErr{
	return NewRestErr(message, "bad_request", http.StatusBadRequest, nil)
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr{
	return NewRestErr(message, "bad_request", http.StatusBadRequest, causes)
}

func NewInternalServerError(message string) *RestErr{
	return NewRestErr(message, "internal_server_error", http.StatusInternalServerError, nil)
}

func NewUnauthorizedError(message string) *RestErr{
	return NewRestErr(message, "unauthorized", http.StatusUnauthorized, nil)
}

func NewForbiddenError(message string) *RestErr{
	return NewRestErr(message, "forbidden", http.StatusForbidden, nil)
}

func NewConflictError(message string) *RestErr{
	return NewRestErr(message, "conflict", http.StatusConflict, nil)
}

func NewNotFoundError(message string) *RestErr{
	return NewRestErr(message, "not_found", http.StatusNotFound, nil)
}