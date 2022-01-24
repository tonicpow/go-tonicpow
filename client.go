package tonicpow

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

type (
	// Client is the TonicPow client/configuration
	Client struct {
		httpClient *resty.Client
		options    *ClientOptions // Options are all the default settings / configuration
	}

	// ClientOptions holds all the configuration for client requests and default resources
	ClientOptions struct {
		apiKey         string              // API key
		env            Environment         // Environment
		customHeaders  map[string][]string // Custom headers on outgoing requests
		httpTimeout    time.Duration       // Default timeout in seconds for GET requests
		requestTracing bool                // If enabled, it will trace the request timing
		retryCount     int                 // Default retry count for HTTP requests
		userAgent      string              // User agent for all outgoing requests
	}

	// StandardResponse is the standard fields returned on all responses
	StandardResponse struct {
		Body       []byte          `json:"-"` // Body of the response request
		Error      *Error          `json:"-"` // API error response
		StatusCode int             `json:"-"` // Status code returned on the request
		Tracing    resty.TraceInfo `json:"-"` // Trace information if enabled on the request
	}
)

// NewClient creates a new client for all TonicPow requests
//
// If no options are given, it will use the DefaultClientOptions()
// If there is no client is supplied, it will use a default Resty HTTP client.
func NewClient(opts ...ClientOps) (ClientInterface, error) {
	defaults := defaultClientOptions()

	// Create a new client
	client := &Client{
		options: defaults,
	}
	// overwrite defaults with any set by user
	for _, opt := range opts {
		opt(client.options)
	}
	// Check for API key
	if client.options.apiKey == "" {
		return nil, errors.New("missing an API Key")
	}
	// Set the Resty HTTP client
	if client.httpClient == nil {
		client.httpClient = resty.New()
		// Set defaults (for GET requests)
		client.httpClient.SetTimeout(client.options.httpTimeout)
		client.httpClient.SetRetryCount(client.options.retryCount)
	}
	return client, nil
}

// WithCustomHTTPClient will overwrite the default client with a custom client.
func (c *Client) WithCustomHTTPClient(client *resty.Client) *Client {
	c.httpClient = client
	return c
}

// GetUserAgent will return the user agent string of the client
func (c *Client) GetUserAgent() string {
	return c.options.userAgent
}

// GetEnvironment will return the Environment of the client
func (c *Client) GetEnvironment() Environment {
	return c.options.env
}

// Options will return the clients current options
func (c *Client) Options() *ClientOptions {
	return c.options
}

// Request is a standard GET / POST / PUT / DELETE request for all outgoing HTTP requests
// Omit the data attribute if using a GET request
func (c *Client) Request(httpMethod string, requestEndpoint string,
	data interface{}, expectedCode int) (response *StandardResponse, err error) {

	// Set the user agent
	req := c.httpClient.R().SetHeader("User-Agent", c.options.userAgent)

	// Set the body if (PUT || POST)
	if httpMethod != http.MethodGet && httpMethod != http.MethodDelete {
		var j []byte
		if j, err = json.Marshal(data); err != nil {
			return
		}
		req = req.SetBody(string(j))
		req.Header.Add("Content-Length", strconv.Itoa(len(j)))
		req.Header.Set("Content-Type", "application/json")
	}

	// Enable tracing
	if c.options.requestTracing {
		req.EnableTrace()
	}

	// Set the authorization and content type
	req.Header.Set(fieldAPIKey, c.options.apiKey)

	// Custom headers?
	for key, headers := range c.options.customHeaders {
		for _, value := range headers {
			req.Header.Set(key, value)
		}
	}

	// Fire the request
	var resp *resty.Response
	switch httpMethod {
	case http.MethodPost:
		resp, err = req.Post(c.options.env.URL() + requestEndpoint)
	case http.MethodPut:
		resp, err = req.Put(c.options.env.URL() + requestEndpoint)
	case http.MethodDelete:
		resp, err = req.Delete(c.options.env.URL() + requestEndpoint)
	case http.MethodGet:
		resp, err = req.Get(c.options.env.URL() + requestEndpoint)
	}
	if err != nil {
		return
	}

	// Start the response
	response = new(StandardResponse)

	// Tracing enabled?
	if c.options.requestTracing {
		response.Tracing = resp.Request.TraceInfo()
	}

	// Set the status code & body
	response.StatusCode = resp.StatusCode()
	response.Body = resp.Body()

	// Check expected code if set
	if expectedCode > 0 && response.StatusCode != expectedCode {
		response.Error = new(Error)
		if err = json.Unmarshal(response.Body, &response.Error); err != nil {
			return
		}
		err = fmt.Errorf("%s", response.Error.Message)
	}

	return
}
