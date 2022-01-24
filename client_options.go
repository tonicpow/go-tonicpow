package tonicpow

import (
	"strings"
	"time"
)

// ClientOps allow functional options to be supplied
// that overwrite default client options.
type ClientOps func(c *ClientOptions)

// defaultClientOptions will return a ClientOptions struct with the default settings
//
// Useful for starting with the default and then modifying as needed
func defaultClientOptions() (opts *ClientOptions) {
	// Set the default options
	opts = &ClientOptions{
		env:            EnvironmentLive,
		httpTimeout:    defaultHTTPTimeout,
		requestTracing: false,
		retryCount:     defaultRetryCount,
		userAgent:      defaultUserAgent,
	}
	return
}

// WithEnvironment will change the Environment
func WithEnvironment(e Environment) ClientOps {
	return func(c *ClientOptions) {
		c.env = e
	}
}

// WithCustomEnvironment will set a custom Environment
func WithCustomEnvironment(name, alias, apiURL string) ClientOps {
	return WithEnvironment(Environment{
		alias:  alias,
		apiURL: apiURL,
		name:   name,
	})
}

// WithEnvironmentString will change the Environment
func WithEnvironmentString(e string) ClientOps {
	e = strings.ToLower(strings.TrimSpace(e))
	if e == environmentStagingName || e == environmentStagingAlias {
		return WithEnvironment(EnvironmentStaging)
	} else if e == environmentDevelopmentName || e == environmentDevelopmentAlias {
		return WithEnvironment(EnvironmentDevelopment)
	}
	return WithEnvironment(EnvironmentLive)
}

// WithHTTPTimeout can be supplied to adjust the default http client timeouts.
// The http client is used when creating requests
// Default timeout is 10 seconds.
func WithHTTPTimeout(timeout time.Duration) ClientOps {
	return func(c *ClientOptions) {
		c.httpTimeout = timeout
	}
}

// WithRequestTracing will enable tracing.
// Tracing is disabled by default.
func WithRequestTracing() ClientOps {
	return func(c *ClientOptions) {
		c.requestTracing = true
	}
}

// WithRetryCount will overwrite the default retry count for http requests.
// Default retries is 2.
func WithRetryCount(retries int) ClientOps {
	return func(c *ClientOptions) {
		c.retryCount = retries
	}
}

// WithUserAgent will overwrite the default useragent.
// Default is package name + version.
func WithUserAgent(userAgent string) ClientOps {
	return func(c *ClientOptions) {
		c.userAgent = userAgent
	}
}

// WithAPIKey provides the API key
func WithAPIKey(appAPIKey string) ClientOps {
	return func(c *ClientOptions) {
		c.apiKey = appAPIKey
	}
}

// WithCustomHeaders will add custom headers to outgoing requests
// Custom headers is empty by default
func WithCustomHeaders(headers map[string][]string) ClientOps {
	return func(c *ClientOptions) {
		c.customHeaders = headers
	}
}
