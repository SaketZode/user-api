package resterrors

import (
	"net/http"
)

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message:    message,
		HttpStatus: http.StatusBadRequest,
		Error:      "BAD_REQUEST",
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message:    message,
		HttpStatus: http.StatusInternalServerError,
		Error:      "INTERNAL_SERVER_ERROR",
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message:    message,
		HttpStatus: http.StatusNotFound,
		Error:      "NOT_FOUND",
	}
}

func NewNotImplementedError(message string) *RestError {
	return &RestError{
		Message:    message,
		HttpStatus: http.StatusNotImplemented,
		Error:      "METHOD_NOT_IMPLEMENTED",
	}
}
