package tonicpow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gojektech/heimdall/v6"
	"github.com/gojektech/heimdall/v6/httpclient"
)

// Client is the parent struct that wraps the heimdall client
type Client struct {
	httpClient  heimdall.Client // carries out the http operations
	LastRequest *LastRequest    // is the raw information from the last Request
	Parameters  *Parameters     // contains application specific values
}

// Options holds all the configuration for connection, dialer and transport
type Options struct {
	BackOffExponentFactor          float64       `json:"back_off_exponent_factor"`
	BackOffInitialTimeout          time.Duration `json:"back_off_initial_timeout"`
	BackOffMaximumJitterInterval   time.Duration `json:"back_off_maximum_jitter_interval"`
	BackOffMaxTimeout              time.Duration `json:"back_off_max_timeout"`
	DialerKeepAlive                time.Duration `json:"dialer_keep_alive"`
	DialerTimeout                  time.Duration `json:"dialer_timeout"`
	RequestRetryCount              int           `json:"request_retry_count"`
	RequestTimeout                 time.Duration `json:"request_timeout"`
	TransportExpectContinueTimeout time.Duration `json:"transport_expect_continue_timeout"`
	TransportIdleTimeout           time.Duration `json:"transport_idle_timeout"`
	TransportMaxIdleConnections    int           `json:"transport_max_idle_connections"`
	TransportTLSHandshakeTimeout   time.Duration `json:"transport_tls_handshake_timeout"`
	UserAgent                      string        `json:"user_agent"`
}

// LastRequest is used to track what was submitted via the Request()
type LastRequest struct {
	Error      *Error `json:"Error"`       // Error is the last Error response from the api
	Method     string `json:"method"`      // method is the HTTP method used
	PostData   string `json:"post_data"`   // postData is the post data submitted if POST/PUT Request
	StatusCode int    `json:"status_code"` // statusCode is the last code from the Request
	URL        string `json:"url"`         // url is the url used for the Request
}

// Parameters are application specific values for requests
type Parameters struct {
	apiKey        string              // is the given api key for the user
	CustomHeaders map[string][]string // is used for setting custom header values on requests
	environment   APIEnvironment      // is the current api environment to use
	UserAgent     string              // (optional for changing user agents)
}

// ClientDefaultOptions will return an Options struct with the default settings
// Useful for starting with the default and then modifying as needed
func ClientDefaultOptions() (clientOptions *Options) {
	return &Options{
		BackOffExponentFactor:          2.0,
		BackOffInitialTimeout:          2 * time.Millisecond,
		BackOffMaximumJitterInterval:   2 * time.Millisecond,
		BackOffMaxTimeout:              10 * time.Millisecond,
		DialerKeepAlive:                20 * time.Second,
		DialerTimeout:                  5 * time.Second,
		RequestRetryCount:              2,
		RequestTimeout:                 10 * time.Second,
		TransportExpectContinueTimeout: 3 * time.Second,
		TransportIdleTimeout:           20 * time.Second,
		TransportMaxIdleConnections:    10,
		TransportTLSHandshakeTimeout:   5 * time.Second,
		UserAgent:                      defaultUserAgent,
	}
}

// createClient will make a new http client based on the options provided
func createClient(options *Options) (c *Client) {

	// Create a client
	c = new(Client)

	// Set options (either default or user modified)
	if options == nil {
		options = ClientDefaultOptions()
	}

	// dial is the net dialer for clientDefaultTransport
	dial := &net.Dialer{KeepAlive: options.DialerKeepAlive, Timeout: options.DialerTimeout}

	// clientDefaultTransport is the default transport struct for the HTTP client
	clientDefaultTransport := &http.Transport{
		DialContext:           dial.DialContext,
		ExpectContinueTimeout: options.TransportExpectContinueTimeout,
		IdleConnTimeout:       options.TransportIdleTimeout,
		MaxIdleConns:          options.TransportMaxIdleConnections,
		Proxy:                 http.ProxyFromEnvironment,
		TLSHandshakeTimeout:   options.TransportTLSHandshakeTimeout,
	}

	// Determine the strategy for the http client (no retry enabled)
	if options.RequestRetryCount <= 0 {
		c.httpClient = httpclient.NewClient(
			httpclient.WithHTTPTimeout(options.RequestTimeout),
			httpclient.WithHTTPClient(&http.Client{
				Transport: clientDefaultTransport,
				Timeout:   options.RequestTimeout,
			}),
		)
	} else { // Retry enabled
		// Create exponential back-off
		backOff := heimdall.NewExponentialBackoff(
			options.BackOffInitialTimeout,
			options.BackOffMaxTimeout,
			options.BackOffExponentFactor,
			options.BackOffMaximumJitterInterval,
		)

		c.httpClient = httpclient.NewClient(
			httpclient.WithHTTPTimeout(options.RequestTimeout),
			httpclient.WithRetrier(heimdall.NewRetrier(backOff)),
			httpclient.WithRetryCount(options.RequestRetryCount),
			httpclient.WithHTTPClient(&http.Client{
				Transport: clientDefaultTransport,
				Timeout:   options.RequestTimeout,
			}),
		)
	}

	// Create a last Request and parameters struct
	c.LastRequest = new(LastRequest)
	c.LastRequest.Error = new(Error)
	c.Parameters = &Parameters{
		UserAgent: options.UserAgent,
	}
	return
}

// Request is a generic wrapper for all api requests
func (c *Client) Request(endpoint string, method string, payload interface{}) (response string, err error) {

	// Set post value
	var jsonValue []byte

	// Add the environment
	endpoint = fmt.Sprintf("%s%s", c.Parameters.environment, endpoint)

	// Switch on methods
	switch method {
	case http.MethodPost, http.MethodPut:
		{
			if jsonValue, err = json.Marshal(payload); err != nil {
				return
			}
		}
	case http.MethodGet:
		{
			if payload != nil {
				params := payload.(url.Values)
				endpoint += "?" + params.Encode()
			}
		}
	}

	// Store for debugging purposes
	c.LastRequest.Method = method
	c.LastRequest.PostData = string(jsonValue)
	c.LastRequest.URL = endpoint

	// Start the Request
	var request *http.Request
	if request, err = http.NewRequest(method, endpoint, bytes.NewBuffer(jsonValue)); err != nil {
		return
	}

	// Set the auth header
	request.Header.Set(fieldAPIKey, c.Parameters.apiKey)

	// Change the user agent
	request.Header.Set("User-Agent", c.Parameters.UserAgent)

	// Custom headers?
	for key, headers := range c.Parameters.CustomHeaders {
		for _, value := range headers {
			request.Header.Set(key, value)
		}
	}

	// Set the content type
	if method == http.MethodPost || method == http.MethodPut {
		request.Header.Set("Content-Type", "application/json")
	}

	// Fire the http Request
	var resp *http.Response
	if resp, err = c.httpClient.Do(request); err != nil {
		return
	}

	// Close the response body
	defer func() {
		_ = resp.Body.Close()
	}()

	// Save the status
	c.LastRequest.StatusCode = resp.StatusCode

	// Read the body
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	// Clear headers
	c.Parameters.CustomHeaders = make(map[string][]string)

	// Parse the response
	response = string(body)
	return
}

// Error will handle all basic error cases
func (c *Client) Error(expectedStatusCode int, response string) (err error) {
	if c.LastRequest.StatusCode != expectedStatusCode {
		c.LastRequest.Error = new(Error)
		if err = json.Unmarshal([]byte(response), c.LastRequest.Error); err != nil {
			return
		}
		err = fmt.Errorf("%s", c.LastRequest.Error.Message)
	}
	return
}
