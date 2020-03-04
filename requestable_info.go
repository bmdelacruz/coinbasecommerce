package coinbasecommerce

// RequestableInfo represents information that can be requested from customers
type RequestableInfo string

// RequestableInfo constants.
const (
	RequestableInfoEmail RequestableInfo = "email"
	RequestableInfoName  RequestableInfo = "name"
)
