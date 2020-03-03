package charges

import (
	"encoding/json"

	"github.com/bmdelacruz/coinbasecommerce"
	"github.com/bmdelacruz/coinbasecommerce/internal"
)

const (
	getEndpointMethod = "GET"
	getEndpoint       = "https://api.commerce.coinbase.com/charges/"
)

// Get retrieves a specific charge object using Coinbase Commerce API.
func Get(
	apiCallContext coinbasecommerce.APICallContext,
	idOrCode string,
) (coinbasecommerce.Charge, coinbasecommerce.Warnings, error) {
	if idOrCode == "" {
		return coinbasecommerce.Charge{}, nil,
			coinbasecommerce.LocalError{
				Inner: coinbasecommerce.ErrInvalidChargeIDOrCode,
			}
	}

	response, err := internal.CreateAndDoHTTPRequest(
		apiCallContext,
		getEndpointMethod,
		getEndpoint+idOrCode,
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
