package charges

import (
	"encoding/json"

	"github.com/bmdelacruz/coinbasecommerce"
	"github.com/bmdelacruz/coinbasecommerce/internal"
)

const (
	listEndpointMethod = "GET"
	listEndpoint       = "https://api.commerce.coinbase.com/charges?"
)

// List retrieves a list of charges from the Coinbase Commerce API.
func List(
	apiCallContext coinbasecommerce.APICallContext,
	paginationOption coinbasecommerce.PaginationOption,
) (
	[]coinbasecommerce.Charge,
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
		Charges    []coinbasecommerce.Charge   `json:"data"`
		Error      *coinbasecommerce.APIError  `json:"error,omitempty"`
		Warnings   coinbasecommerce.Warnings   `json:"warnings"`
	}
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return nil, coinbasecommerce.Pagination{}, nil,
			coinbasecommerce.LocalError{Inner: err}
	}

	return responseBody.Charges, responseBody.Pagination, responseBody.Warnings,
		coinbasecommerce.ReturnAPIErrorAsError(responseBody.Error)
}
