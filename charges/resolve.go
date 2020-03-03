package charges

import (
	"encoding/json"
	"fmt"

	"github.com/bmdelacruz/coinbasecommerce"
	"github.com/bmdelacruz/coinbasecommerce/internal"
)

const (
	resolveEndpointMethod = "POST"
	resolveEndpointFmt    = "https://api.commerce.coinbase.com/charges/%s/resolve"
)

// Resolve resolves a charge object using the Coinbase Commerce API.
func Resolve(
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
		resolveEndpointMethod,
		fmt.Sprintf(resolveEndpointFmt, idOrCode),
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
