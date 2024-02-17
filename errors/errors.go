package application_errors

import "errors"

var (
	ErrInternalServerError  = errors.New("Error getting resources")
	ErrPageNotFound         = errors.New("Page not found")
	ErrBookNotFound         = errors.New("Book not found")
	ErrInvalidContentFormat = errors.New("Invalid content format")
)
