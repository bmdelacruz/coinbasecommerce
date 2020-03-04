package coinbasecommerce

// Checkout represents a Coinbase Commerce API checkout object
type Checkout struct {
	ID            string            `json:"id"`
	Resource      string            `json:"resource"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	LogoURL       string            `json:"logo_url,omitempty"`
	PricingType   PricingType       `json:"pricing_type"`
	LocalPrice    *Money            `json:"local_price,omitempty"`
	RequestedInfo []RequestableInfo `json:"requested_info,omitempty"`
}
