package checkouts

import (
	"encoding/json"

	"github.com/bmdelacruz/coinbasecommerce"
	"github.com/bmdelacruz/coinbasecommerce/internal"
)

const (
	getEndpointMethod = "GET"
	getEndpoint       = "https://api.commerce.coinbase.com/checkouts/"
)

// Get retrieves a specific checkout object using Coinbase Commerce API.
func Get(
	apiCallContext coinbasecommerce.APICallContext, id string,
) (coinbasecommerce.Checkout, coinbasecommerce.Warnings, error) {
	if id == "" {
		return coinbasecommerce.Checkout{}, nil,
			coinbasecommerce.LocalError{
				Inner: coinbasecommerce.ErrInvalidCheckoutID,
			}
	}

	response, err := internal.CreateAndDoHTTPRequest(
		apiCallContext,
		getEndpointMethod,
		getEndpoint+id,
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
