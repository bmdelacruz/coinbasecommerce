package checkouts

import (
	"encoding/json"

	"github.com/bmdelacruz/coinbasecommerce"
	"github.com/bmdelacruz/coinbasecommerce/internal"
)

const (
	listEndpointMethod = "GET"
	listEndpoint       = "https://api.commerce.coinbase.com/checkouts?"
)

// List retrieves a list of checkouts from the Coinbase Commerce API.
func List(
	apiCallContext coinbasecommerce.APICallContext,
	paginationOption coinbasecommerce.PaginationOption,
) (
	[]coinbasecommerce.Checkout,
	coinbasecommerce.Pagination,
	coinbasecommerce.Warnings,
	error,
) {
	response, err := internal.CreateAndDoHTTPRequest(
		apiCallContext,
		listEndpointMethod,
		listEndpoint+paginationOption.MakeQueryString(),
	)
	if err != nil {
		return nil, coinbasecommerce.Pagination{}, nil,
			coinbasecommerce.LocalError{Inner: err}
	}
	defer response.Body.Close()

	var responseBody struct {
		Pagination coinbasecommerce.Pagination `json:"pagination"`
		Checkouts  []coinbasecommerce.Checkout `json:"data"`
		Error      *coinbasecommerce.APIError  `json:"error,omitempty"`
		Warnings   coinbasecommerce.Warnings   `json:"warnings"`
	}
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return nil, coinbasecommerce.Pagination{}, nil,
			coinbasecommerce.LocalError{Inner: err}
	}

	return responseBody.Checkouts, responseBody.Pagination, responseBody.Warnings,
		coinbasecommerce.ReturnAPIErrorAsError(responseBody.Error)
}
