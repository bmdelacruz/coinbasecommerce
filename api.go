package coinbasecommerce

import (
	"context"
	"net/http"
)

// Header keys that are required by the Coinbase Commerce API.
const (
	APIHeaderAPIKey  = "X-CC-Api-Key"
	APIHeaderVersion = "X-CC-Version"
)

// APIConfig contains configuration that's needed by the Coinbase Commerce API.
type APIConfig struct {
	apiKey  string
	version string
}

// APIKey returns the API key from the API configuration object.
func (cfg *APIConfig) APIKey() string {
	return cfg.apiKey
}

// Version returns the API version from the API configuration object.
func (cfg *APIConfig) Version() string {
	return cfg.version
}

// NewAPIConfig creates a new API configuration.
func NewAPIConfig(apiKey, version string) *APIConfig {
	return &APIConfig{apiKey, version}
}

// APICallContext contains objects that will be used during the execution
// of a request to the Coinbase Commerce API.
type APICallContext struct {
	apiConfig  *APIConfig
	httpClient *http.Client
	context    context.Context
}

// APIConfig returns the API configuration object that will be used to
// execute a request to the Coinbase Commerce API.
func (acc *APICallContext) APIConfig() *APIConfig {
	return acc.apiConfig
}

// HTTPClient returns the HTTP client that will be used to send the request
// to the Coinbase Commerce API.
func (acc *APICallContext) HTTPClient() *http.Client {
	return acc.httpClient
}

// Context returns context that will be used when sending the request to the
// Coinbase Commerce API; may be equal to nil.
func (acc *APICallContext) Context() context.Context {
	return acc.context
}

// APICallContextOptions contains options for the Create API call.
type APICallContextOptions struct {
	httpClient *http.Client
	context    context.Context
}

// APICallContextOptionFunc represents a function that can modify the contents
// of the APICallContextOptions.
type APICallContextOptionFunc func(*APICallContextOptions)

// APICallContextOptionHTTPClient sets the HTTP client that will be used to send
// a request to the Coinbase Commerce API.
func APICallContextOptionHTTPClient(httpClient *http.Client) APICallContextOptionFunc {
	return func(options *APICallContextOptions) {
		options.httpClient = httpClient
	}
}

// APICallContextOptionContext sets the context that will be used when sending a
// request to the Coinbase Commerce API.
func APICallContextOptionContext(context context.Context) APICallContextOptionFunc {
	return func(options *APICallContextOptions) {
		options.context = context
	}
}

// NewAPICallContext creates a new API call context.
func NewAPICallContext(
	apiConfig *APIConfig,
	optionFuncs ...APICallContextOptionFunc,
) APICallContext {
	if apiConfig == nil {
		panic("apiConfig cannot be equal to nil")
	}

	options := APICallContextOptions{
		httpClient: http.DefaultClient,
		context:    nil,
	}
	for _, optionFunc := range optionFuncs {
		optionFunc(&options)
	}

	return APICallContext{
		apiConfig:  apiConfig,
		httpClient: options.httpClient,
		context:    options.context,
	}
}
