package coinbasecommerce

import (
	"errors"
	"fmt"
)

// APIError contains the details of the error that was received
// from the Coinbase Commerce API.
type APIError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("coinbase commerce api error: %s (%s)", e.Type, e.Message)
}

// Is checks if e and target are of the same type
func (e APIError) Is(target error) bool {
	t, ok := target.(*APIError)
	return ok && t.Type == e.Type
}

// API Errors
var (
	ErrAPIInvalidRequest error = &APIError{Type: "invalid_request"}
)

// LocalError is an error that occured here in the package
type LocalError struct {
	Inner error
}

func (e LocalError) Error() string {
	return fmt.Sprintf("local error: %v", e.Inner)
}

func (e LocalError) Unwrap() error {
	return e.Inner
}

// Is checks if e and target are of the same type
func (e LocalError) Is(target error) bool {
	_, ok := target.(*LocalError)
	return ok
}

// Local inner errors
var (
	ErrInvalidChargeIDOrCode = errors.New("invalid charge id or code")
)
