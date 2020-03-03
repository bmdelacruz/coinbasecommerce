package charges

import (
	"bytes"
	"encoding/json"

	"github.com/bmdelacruz/coinbasecommerce"
	"github.com/bmdelacruz/coinbasecommerce/internal"
)

// CreateRequest contains fields that are needed to create a charge
// using Coinbase Commerce API.
type CreateRequest struct {
	// Name is the name that should be given to the charge (required, <=100 characters).
	Name string `json:"name"`
	// Description describes the charge (required, <=200 characters).
	Description string `json:"description"`
	// PricingType is the type of how an entity should be charged.
	PricingType coinbasecommerce.PricingType `json:"pricing_type"`
	// LocalPrice represents the amount that should be charged from an entity
	// (required if pricing type is equal to 'fixed_price').
	LocalPrice *coinbasecommerce.Money `json:"local_price,omitempty"`
	// Metadata represents an arbitrary data that will be included in the charge object.
	Metadata    map[string]string `json:"metadata,omitempty"`
	RedirectURL string            `json:"redirect_url,omitempty"`
	CancelURL   string            `json:"cancel_url,omitempty"`
}

const (
	createEndpointMethod = "POST"
	createEndpoint       = "https://api.commerce.coinbase.com/charges"
)

// Create creates a charge by sending a request to the Coinbase Commerce API.
func Create(
	apiCallContext coinbasecommerce.APICallContext,
	request CreateRequest,
) (coinbasecommerce.Charge, coinbasecommerce.Warnings, error) {
	bodyBuffer := new(bytes.Buffer)
	err := json.NewEncoder(bodyBuffer).Encode(request)
	if err != nil {
		return coinbasecommerce.Charge{}, nil,
			coinbasecommerce.LocalError{Inner: err}
	}

	response, err := internal.CreateAndDoHTTPRequest(
		apiCallContext,
		createEndpointMethod,
		createEndpoint,
		internal.CreateAndDoHTTPRequestOptionsJSONBody(bodyBuffer),
	)
	if err != nil {
		return coinbasecommerce.Charge{}, nil,
			coinbasecommerce.LocalError{Inner: err}
	}
	defer response.Body.Close()

	var responseBody struct {
		Charge   coinbasecommerce.Charge    `json:"data"`
		Error    *coinbasecommerce.APIError `json:"error,omitempty"`
		Warnings coinbasecommerce.Warnings  `json:"warnings"`
	}
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return coinbasecommerce.Charge{}, nil,
			coinbasecommerce.LocalError{Inner: err}
	}

	return responseBody.Charge, responseBody.Warnings, responseBody.Error
}
