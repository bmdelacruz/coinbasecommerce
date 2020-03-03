package coinbasecommerce

import "fmt"

// Error contains the details of the error that was received
// from the Coinbase Commerce API.
type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (e *Error) String() string {
	return fmt.Sprintf("%s (%s)", e.Type, e.Message)
}
