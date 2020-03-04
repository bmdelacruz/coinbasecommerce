package checkouts

import (
	"encoding/json"
	"fmt"

	"github.com/bmdelacruz/coinbasecommerce"
	"github.com/bmdelacruz/coinbasecommerce/internal"
)

const (
	deleteEndpointMethod = "DELETE"
	deleteEndpointFmt    = "https://api.commerce.coinbase.com/checkouts/%s"
)

// Delete deletes a checkout object using the Coinbase Commerce API.
func Delete(
	apiCallContext coinbasecommerce.APICallContext,
	id string,
) (coinbasecommerce.Warnings, error) {
	if id == "" {
		return nil, coinbasecommerce.LocalError{
			Inner: coinbasecommerce.ErrInvalidCheckoutID,
		}
	}

	response, err := internal.CreateAndDoHTTPRequest(
		apiCallContext,
		deleteEndpointMethod,
		fmt.Sprintf(deleteEndpointFmt, id),
	)
	if err != nil {
		return nil, coinbasecommerce.LocalError{Inner: err}
	}
	defer response.Body.Close()

	var responseBody struct {
		Error    *coinbasecommerce.APIError `json:"error,omitempty"`
		Warnings coinbasecommerce.Warnings  `json:"warnings"`
	}
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return nil, coinbasecommerce.LocalError{Inner: err}
	}

	return responseBody.Warnings, coinbasecommerce.ReturnAPIErrorAsError(responseBody.Error)
}
