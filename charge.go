package coinbasecommerce

import "time"

// ChargeStatus represents the status of a charge.
type ChargeStatus string

// ChargeStatus constants. These are the possible statuses of a charge.
const (
	ChargeStatusNew        ChargeStatus = "NEW"
	ChargeStatusPending    ChargeStatus = "PENDING"
	ChargeStatusCompleted  ChargeStatus = "COMPLETED"
	ChargeStatusExpired    ChargeStatus = "EXPIRED"
	ChargeStatusUnresolved ChargeStatus = "UNRESOLVED"
	ChargeStatusResolved   ChargeStatus = "RESOLVED"
	ChargeStatusCanceled   ChargeStatus = "CANCELED"
)

// ChargeStatusUpdateContext represents the reason why the current status of
// a charge is equal to `UNRESOLVED`.
type ChargeStatusUpdateContext string

// StatusUpdateContext constants. These are the possible context of an
// unresolved charge.
const (
	ChargeStatusUpdateContextUnderpaid ChargeStatusUpdateContext = "UNDERPAID"
	ChargeStatusUpdateContextOverpaid  ChargeStatusUpdateContext = "OVERPAID"
	ChargeStatusUpdateContextDelayed   ChargeStatusUpdateContext = "DELAYED"
	ChargeStatusUpdateContextMultiple  ChargeStatusUpdateContext = "MULTIPLE"
	ChargeStatusUpdateContextManual    ChargeStatusUpdateContext = "MANUAL"
	ChargeStatusUpdateContextOther     ChargeStatusUpdateContext = "OTHER"
)

// ChargeStatusUpdate represents an update regarding a charge. `ChargeStatusUpdate.Context`
// will only be not empty when the `ChargeStatusUpdate.Status` is equal to `StatusUnresolved`.
type ChargeStatusUpdate struct {
	Time    time.Time                 `json:"time"`
	Status  ChargeStatus              `json:"status"`
	Context ChargeStatusUpdateContext `json:"context,omitempty"`
}

// Charge contains full details about a charge.
type Charge struct {
	ID          string               `json:"id"`
	Resource    string               `json:"resource"`
	Code        string               `json:"code"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	LogoURL     string               `json:"logo_url"`
	HostedURL   string               `json:"hosted_url"`
	CreatedAt   time.Time            `json:"created_at"`
	ExpiresAt   time.Time            `json:"expires_at"`
	ConfirmedAt time.Time            `json:"confirmed_at"`
	Checkout    map[string]string    `json:"checkout"`
	Timeline    []ChargeStatusUpdate `json:"timeline"`
	Metadata    map[string]string    `json:"metadata"`
	PricingType PricingType          `json:"pricing_type"`
	Pricing     map[string]Money     `json:"pricing"`
	Payments    []interface{}        `json:"payments"`
}
