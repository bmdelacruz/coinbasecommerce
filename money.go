package coinbasecommerce

// Money contains details of the money, its amount and currency.
type Money struct {
	Amount   float64  `json:"amount,string"`
	Currency Currency `json:"currency"`
}
