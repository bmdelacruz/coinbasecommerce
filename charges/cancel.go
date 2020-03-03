package charges

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bmdelacruz/coinbasecommerce"
	"github.com/bmdelacruz/coinbasecommerce/internal"
)

// CancelResponse represents the data that may be returned when a request
// to cancel a charge object to the Coinbase Commerce API Is made.
type CancelResponse struct {
	Charge   coinbasecommerce.Charge   `json:"data"`
	Error    *coinbasecommerce.Error   `json:"error,omitempty"`
	Warnings coinbasecommerce.Warnings `json:"warnings"`
}

const (
	cancelEndpointMethod = "POST"
	cancelEndpointFmt    = "https://api.commerce.coinbase.com/charges/%s/cancel"
)

// Cancel cancels a charge object using the Coinbase Commerce API.
func Cancel(
	apiCallContext coinbasecommerce.APICallContext,
	idOrCode string,
) (CancelResponse, error) {
	if idOrCode == "" {
		return CancelResponse{}, errors.New("idOrCode cannot be empty")
	}
	response, err := internal.CreateAndDoHTTPRequest(
		apiCallContext,
		cancelEndpointMethod,
		fmt.Sprintf(cancelEndpointFmt, idOrCode),
	)
	if err != nil {
		return CancelResponse{}, err
	}
	defer response.Body.Close()

	var responseBody CancelResponse
	return responseBody, json.NewDecoder(response.Body).Decode(&responseBody)
}
