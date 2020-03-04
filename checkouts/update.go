package checkouts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/bmdelacruz/coinbasecommerce"
	"github.com/bmdelacruz/coinbasecommerce/internal"
)

// UpdateOptions contains the checkout field data which should be updated
type UpdateOptions struct {
	updateName          bool
	name                string
	updateDescription   bool
	description         string
	updateLocalPrice    bool
	localPrice          coinbasecommerce.Money
	updateRequestedInfo bool
	requestedInfo       []coinbasecommerce.RequestableInfo
}

// UpdateOptionsFunc is a function that can modify the UpdateOptions
type UpdateOptionsFunc func(*UpdateOptions)

// UpdateOptionsName sets the name field of the UpdateOptions
func UpdateOptionsName(name string) UpdateOptionsFunc {
	return func(options *UpdateOptions) {
		options.updateName = true
		options.name = name
	}
}

// UpdateOptionsDescription sets the description field of the UpdateOptions
func UpdateOptionsDescription(description string) UpdateOptionsFunc {
	return func(options *UpdateOptions) {
		options.updateDescription = true
		options.description = description
	}
}

// UpdateOptionsLocalPrice sets the local price field of the UpdateOptions
func UpdateOptionsLocalPrice(localPrice coinbasecommerce.Money) UpdateOptionsFunc {
	return func(options *UpdateOptions) {
		options.updateLocalPrice = true
		options.localPrice = localPrice
	}
}

// UpdateOptionsRequestedInfo sets the requested info field of the UpdateOptions
func UpdateOptionsRequestedInfo(requestedInfo []coinbasecommerce.RequestableInfo) UpdateOptionsFunc {
	return func(options *UpdateOptions) {
		options.updateRequestedInfo = true
		options.requestedInfo = requestedInfo
	}
}

// NewUpdateOptions creates a new UpdateOptions
func NewUpdateOptions(optionFuncs ...UpdateOptionsFunc) UpdateOptions {
	options := UpdateOptions{
		updateName:          false,
		updateDescription:   false,
		updateLocalPrice:    false,
		updateRequestedInfo: false,
	}
	for _, optionFunc := range optionFuncs {
		optionFunc(&options)
	}
	return options
}

// HasUpdates returns if the UpdateOptions object has fields to update
func (options *UpdateOptions) HasUpdates() bool {
	return options.updateName || options.updateDescription ||
		options.updateLocalPrice || options.updateRequestedInfo
}

// WriteJSON writes a JSON representation of the update to the writer
func (options *UpdateOptions) WriteJSON(w io.Writer) error {
	optionsMap := make(map[string]interface{})

	if options.updateName {
		optionsMap["name"] = options.name
	}
	if options.updateDescription {
		optionsMap["description"] = options.description
	}
	if options.updateLocalPrice {
		optionsMap["local_price"] = &options.localPrice
	}
	if options.updateRequestedInfo {
		if len(options.requestedInfo) == 0 {
			optionsMap["requested_info"] = make([]coinbasecommerce.RequestableInfo, 0)
		} else {
			optionsMap["requested_info"] = options.requestedInfo
		}
	}

	return json.NewEncoder(w).Encode(&optionsMap)
}

const (
	updateEndpointMethod = "PUT"
	updateEndpointFmt    = "https://api.commerce.coinbase.com/checkouts/%s"
)

// Errors related to Update function
var (
	ErrNothingToUpdate error = errors.New("nothing to update")
)

// Update updates the fields of a checkout object with matching id using the Coinbase
// Commerce API
func Update(
	apiCallContext coinbasecommerce.APICallContext,
	id string,
	options UpdateOptions,
) (coinbasecommerce.Checkout, coinbasecommerce.Warnings, error) {
	if id == "" {
		return coinbasecommerce.Checkout{}, nil,
			coinbasecommerce.LocalError{
				Inner: coinbasecommerce.ErrInvalidCheckoutID,
			}
	} else if !options.HasUpdates() {
		return coinbasecommerce.Checkout{}, nil,
			coinbasecommerce.LocalError{
				Inner: ErrNothingToUpdate,
			}
	}

	bodyBuffer := new(bytes.Buffer)
	if err := options.WriteJSON(bodyBuffer); err != nil {
		return coinbasecommerce.Checkout{}, nil,
			coinbasecommerce.LocalError{Inner: err}
	}

	response, err := internal.CreateAndDoHTTPRequest(
		apiCallContext,
		updateEndpointMethod,
		fmt.Sprintf(updateEndpointFmt, id),
		internal.CreateAndDoHTTPRequestOptionsJSONBody(bodyBuffer),
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
