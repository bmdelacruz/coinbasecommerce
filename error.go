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
	if target == nil {
		return false
	}
	t, ok := target.(*APIError)
	return ok && t.Type == e.Type
}

// ReturnAPIErrorAsError does what it says it does. We need to explicitly check if
// `apiError` is not equal to nil because, unfortunately, a nil APIError pointer
// that's returned as an error will fail nil check, i.e. `err != nil`.
func ReturnAPIErrorAsError(apiError *APIError) error {
	if apiError == nil {
		return nil
	}
	return apiError
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
	if target == nil {
		return false
	}
	_, ok := target.(*LocalError)
	return ok
}

// Local inner errors
var (
	ErrInvalidChargeIDOrCode = errors.New("invalid charge id or code")
)
