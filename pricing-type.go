package coinbasecommerce

// PricingType represents how an entity should be charged.
type PricingType string

// PricingType constants.
const (
	PricingTypeNone  PricingType = "no_price"
	PricingTypeFixed PricingType = "fixed_price"
)
