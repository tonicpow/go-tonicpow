package tonicpow

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gojek/heimdall"
	"github.com/gojek/heimdall/httpclient"
)

// Client is the parent struct that wraps the heimdall client
type Client struct {
	httpClient  heimdall.Client // carries out the POST operations
	LastRequest *LastRequest    // is the raw information from the last request
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
	Method     string `json:"method"`      // method is the HTTP method used
	PostData   string `json:"post_data"`   // postData is the post data submitted if POST/PUT request
	StatusCode int    `json:"status_code"` // statusCode is the last code from the request
	URL        string `json:"url"`         // url is the url used for the request
}

// Parameters are application specific values for requests
type Parameters struct {
	apiKey            string         // is the given api key for the user
	apiSessionCookie  *http.Cookie   // is the current session cookie for the api key
	environment       APIEnvironment // is the current api environment to use
	UserAgent         string         // (optional for changing user agents)
	UserSessionCookie *http.Cookie   // is the current session cookie for a user (on behalf)
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
		UserAgent:                      DefaultUserAgent,
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

	// Create a last request and parameters struct
	c.LastRequest = new(LastRequest)
	c.Parameters = &Parameters{
		UserAgent: options.UserAgent,
	}
	return
}

// request is a generic wrapper for all api requests
func (c *Client) request(endpoint string, method string, params *url.Values, customSessionToken string) (response string, err error) {

	// Set reader
	var bodyReader io.Reader

	// Add the network value
	endpoint = fmt.Sprintf("%s%s", c.Parameters.environment, endpoint)

	// Switch on methods
	switch method {
	case http.MethodPost, http.MethodPut:
		{
			encodedParams := params.Encode()
			bodyReader = strings.NewReader(encodedParams)
			c.LastRequest.PostData = encodedParams
		}
	case http.MethodGet:
		{
			if params != nil {
				endpoint += "?" + params.Encode()
			}
		}
	}

	// Store for debugging purposes
	c.LastRequest.Method = method
	c.LastRequest.URL = endpoint

	// Start the request
	var request *http.Request
	if request, err = http.NewRequest(method, endpoint, bodyReader); err != nil {
		return
	}

	// Change the user agent
	request.Header.Set("User-Agent", c.Parameters.UserAgent)

	// Set the content type
	if method == http.MethodPost || method == http.MethodPut {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	// Custom token, used for user related requests
	if len(customSessionToken) > 0 {
		request.AddCookie(&http.Cookie{
			Name:     SessionCookie,
			Value:    customSessionToken,
			MaxAge:   60 * 60 * 24,
			HttpOnly: true,
		})
	} else if c.Parameters.apiSessionCookie != nil {
		request.AddCookie(c.Parameters.apiSessionCookie)
	}

	// Fire the http request
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

	// Got a session token? Set the session token for the api user or user via on behalf
	for _, cookie := range resp.Cookies() {
		if cookie.Name == SessionCookie {
			if cookie.MaxAge <= 0 {
				cookie = nil
			}
			if len(customSessionToken) > 0 {
				c.Parameters.UserSessionCookie = cookie
			} else {
				c.Parameters.apiSessionCookie = cookie
			}
			break
		}
	}

	// Parse the response
	response = string(body)
	return
}

// error will handle all basic error cases
func (c *Client) error(expectedStatusCode int, response string) (err error) {
	if c.LastRequest.StatusCode != expectedStatusCode {
		resp := new(Error)
		if err = json.Unmarshal([]byte(response), resp); err != nil {
			return
		}
		err = fmt.Errorf("error code: %d message: %s", resp.Code, resp.Message)
	}
	return
}
