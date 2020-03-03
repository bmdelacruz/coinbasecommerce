package charges

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bmdelacruz/coinbasecommerce"
	"github.com/bmdelacruz/coinbasecommerce/internal"
)

// ResolveResponse represents the data that may be returned when a request to
// resolve a charge object to the Coinbase Commerce API is made.
type ResolveResponse struct {
	Charge   coinbasecommerce.Charge   `json:"data"`
	Error    *coinbasecommerce.Error   `json:"error,omitempty"`
	Warnings coinbasecommerce.Warnings `json:"warnings"`
}

const (
	resolveEndpointMethod = "POST"
	resolveEndpointFmt    = "https://api.commerce.coinbase.com/charges/%s/resolve"
)

// Resolve resolves a charge object using the Coinbase Commerce API.
func Resolve(
	apiCallContext coinbasecommerce.APICallContext,
	idOrCode string,
) (ResolveResponse, error) {
	if idOrCode == "" {
		return ResolveResponse{}, errors.New("idOrCode cannot be empty")
	}
	response, err := internal.CreateAndDoHTTPRequest(
		apiCallContext,
		resolveEndpointMethod,
		fmt.Sprintf(resolveEndpointFmt, idOrCode),
	)
	if err != nil {
		return ResolveResponse{}, err
	}
	defer response.Body.Close()

	var responseBody ResolveResponse
	return responseBody, json.NewDecoder(response.Body).Decode(&responseBody)
}
