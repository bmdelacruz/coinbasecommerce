package internal

import (
	"io"
	"net/http"

	"github.com/bmdelacruz/coinbasecommerce"
)

// CreateAndDoHTTPRequestOptions contains data for a HTTP request.
type CreateAndDoHTTPRequestOptions struct {
	hasBody         bool
	body            io.Reader
	bodyContentType string
}

// CreateAndDoHTTPRequestOptionsFunc represents a function that receives and
// modifies a CreateAndDoHTTPRequestOptions object.
type CreateAndDoHTTPRequestOptionsFunc func(*CreateAndDoHTTPRequestOptions)

// CreateAndDoHTTPRequestOptionsBody creates a function that sets the body
// and content type of the CreateAndDoHTTPRequestOptions object.
func CreateAndDoHTTPRequestOptionsBody(
	body io.Reader,
	contentType string,
) CreateAndDoHTTPRequestOptionsFunc {
	return func(options *CreateAndDoHTTPRequestOptions) {
		options.hasBody = true
		options.body = body
		options.bodyContentType = contentType
	}
}

// CreateAndDoHTTPRequestOptionsJSONBody creates a function that sets a JSON
// content as the body of a CreateAndDoHTTPRequestOptions object.
func CreateAndDoHTTPRequestOptionsJSONBody(body io.Reader) CreateAndDoHTTPRequestOptionsFunc {
	return CreateAndDoHTTPRequestOptionsBody(body, "application/json")
}

// CreateAndDoHTTPRequest creates a HTTP request, executes it, and then
// returns its response.
func CreateAndDoHTTPRequest(
	apiCallContext coinbasecommerce.APICallContext,
	endpointMethod, endpoint string,
	optionFuncs ...CreateAndDoHTTPRequestOptionsFunc,
) (*http.Response, error) {
	options := CreateAndDoHTTPRequestOptions{
		hasBody:         false,
		body:            nil,
		bodyContentType: "",
	}
	for _, optionFunc := range optionFuncs {
		optionFunc(&options)
	}

	var httpRequest *http.Request
	var err error

	if ctx := apiCallContext.Context(); ctx != nil {
		httpRequest, err = http.NewRequestWithContext(
			ctx,
			endpointMethod,
			endpoint,
			options.body,
		)
	} else {
		httpRequest, err = http.NewRequest(
			endpointMethod,
			endpoint,
			options.body,
		)
	}
	if err != nil {
		return nil, err
	}

	if options.hasBody {
		httpRequest.Header.Set("Content-Type", options.bodyContentType)
	}

	httpRequest.Header.Set(
		coinbasecommerce.APIHeaderAPIKey,
		apiCallContext.APIConfig().APIKey(),
	)
	httpRequest.Header.Set(
		coinbasecommerce.APIHeaderVersion,
		apiCallContext.APIConfig().Version(),
	)

	return apiCallContext.HTTPClient().Do(httpRequest)
}
