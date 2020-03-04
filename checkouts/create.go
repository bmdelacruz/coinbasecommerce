package checkouts

import (
	"bytes"
	"encoding/json"

	"github.com/bmdelacruz/coinbasecommerce"
	"github.com/bmdelacruz/coinbasecommerce/internal"
)

// CreateRequest contains fields that are needed to create a checkout
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
	// RequestedInfo contains the customer's info that needs to be collected (optional).
	RequestedInfo []coinbasecommerce.RequestableInfo `json:"requested_info"`
}

const (
	createEndpointMethod = "POST"
	createEndpoint       = "https://api.commerce.coinbase.com/checkouts"
)

// Create creates a checkout by sending a request to the Coinbase Commerce API.
func Create(
	apiCallContext coinbasecommerce.APICallContext,
	request CreateRequest,
) (coinbasecommerce.Checkout, coinbasecommerce.Warnings, error) {
	bodyBuffer := new(bytes.Buffer)
	err := json.NewEncoder(bodyBuffer).Encode(request)
	if err != nil {
		return coinbasecommerce.Checkout{}, nil,
			coinbasecommerce.LocalError{Inner: err}
	}

	response, err := internal.CreateAndDoHTTPRequest(
		apiCallContext,
		createEndpointMethod,
		createEndpoint,
		internal.CreateAndDoHTTPRequestOptionsJSONBody(bodyBuffer),
	)
	if err != nil {
		return coinbasecommerce.Checkout{}, nil,
			coinbasecommerce.LocalError{Inner: err}
	}
	defer response.Body.Close()

	var responseBody struct {
		Checkout coinbasecommerce.Checkout  `json:"data"`
		Error    *coinbasecommerce.APIError `json:"error,omitempty"`
		Warnings coinbasecommerce.Warnings  `json:"warnings"`
	}
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return coinbasecommerce.Checkout{}, nil,
			coinbasecommerce.LocalError{Inner: err}
	}

	return responseBody.Checkout, responseBody.Warnings,
		coinbasecommerce.ReturnAPIErrorAsError(responseBody.Error)
}
