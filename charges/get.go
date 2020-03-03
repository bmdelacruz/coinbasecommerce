package charges

import (
	"encoding/json"
	"errors"

	"github.com/bmdelacruz/coinbasecommerce"
	"github.com/bmdelacruz/coinbasecommerce/internal"
)

// GetResponse represents the data that may be returned when a
// request to get a specific charge object to the Coinbase Commerce API
// is made.
type GetResponse struct {
	Charge   coinbasecommerce.Charge   `json:"data"`
	Error    *coinbasecommerce.Error   `json:"error,omitempty"`
	Warnings coinbasecommerce.Warnings `json:"warnings"`
}

const (
	getEndpointMethod = "GET"
	getEndpoint       = "https://api.commerce.coinbase.com/charges/"
)

// Get retrieves a specific charge object using Coinbase Commerce API.
func Get(
	apiCallContext coinbasecommerce.APICallContext,
	idOrCode string,
) (GetResponse, error) {
	if idOrCode == "" {
		return GetResponse{}, errors.New("idOrCode cannot be empty")
	}
	response, err := internal.CreateAndDoHTTPRequest(
		apiCallContext,
		getEndpointMethod,
		getEndpoint+idOrCode,
	)
	if err != nil {
		return GetResponse{}, err
	}
	defer response.Body.Close()

	var responseBody GetResponse
	return responseBody, json.NewDecoder(response.Body).Decode(&responseBody)
}
