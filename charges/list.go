package charges

import (
	"encoding/json"

	"github.com/bmdelacruz/coinbasecommerce"
	"github.com/bmdelacruz/coinbasecommerce/internal"
)

// ListResponse represents the data that may be returned when a
// request to get a list of charges to the Coinbase Commerce API
// is made.
type ListResponse struct {
	Pagination coinbasecommerce.Pagination `json:"pagination"`
	Charges    []coinbasecommerce.Charge   `json:"data"`
	Error      *coinbasecommerce.Error     `json:"error,omitempty"`
	Warnings   coinbasecommerce.Warnings   `json:"warnings"`
}

const (
	listEndpointMethod = "GET"
	listEndpoint       = "https://api.commerce.coinbase.com/charges?"
)

// List retrieves a list of charges from the Coinbase Commerce API.
func List(
	apiCallContext coinbasecommerce.APICallContext,
	paginationOption coinbasecommerce.PaginationOption,
) (ListResponse, error) {
	response, err := internal.CreateAndDoHTTPRequest(
		apiCallContext,
		listEndpointMethod,
		listEndpoint+paginationOption.MakeQueryString(),
	)
	if err != nil {
		return ListResponse{}, err
	}
	defer response.Body.Close()

	var responseBody ListResponse
	return responseBody, json.NewDecoder(response.Body).Decode(&responseBody)
}
