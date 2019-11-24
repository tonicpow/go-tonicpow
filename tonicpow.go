/*
Package tonicpow is the official golang implementation for the TonicPow API
*/
package tonicpow

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gojek/heimdall"
	"github.com/gojek/heimdall/httpclient"
)

// Client holds client configuration settings
type Client struct {

	// HTTPClient carries out the POST operations
	HTTPClient heimdall.Client

	// Parameters contains the search parameters that are submitted with your query,
	// which may affect the data returned
	Parameters *RequestParameters

	// LastRequest is the raw information from the last request
	LastRequest *LastRequest
}

// RequestParameters holds options that can affect data returned by a request.
type RequestParameters struct {

	// AdvertiserSecretKey
	AdvertiserSecretKey string

	// Environment
	Environment APIEnvironment

	// UserAgent (optional for changing user agents)
	UserAgent string
}

// LastRequest is used to track what was submitted to whatsonchain on the Request()
type LastRequest struct {

	// Method is either POST or GET
	Method string

	// PostData is the post data submitted if POST request
	PostData string

	// StatusCode is the last code from the request
	StatusCode int

	// URL is the url used for the request
	URL string
}

// NewClient creates a new client to submit queries with.
// Parameters values are set to the defaults defined by TonicPow.
//
// For more information: https://tonicpow.com
func NewClient(advertiserSecretKey string) (c *Client, err error) {

	// Create a client
	c = new(Client)

	// Create exponential backoff
	backOff := heimdall.NewExponentialBackoff(
		ConnectionInitialTimeout,
		ConnectionMaxTimeout,
		ConnectionExponentFactor,
		ConnectionMaximumJitterInterval,
	)

	// Create the http client
	//c.HTTPClient = new(http.Client) (@mrz this was the original HTTP client)
	c.HTTPClient = httpclient.NewClient(
		httpclient.WithHTTPTimeout(ConnectionWithHTTPTimeout),
		httpclient.WithRetrier(heimdall.NewRetrier(backOff)),
		httpclient.WithRetryCount(ConnectionRetryCount),
		httpclient.WithHTTPClient(&http.Client{
			Transport: ClientDefaultTransport,
			Timeout:   ConnectionWithHTTPTimeout,
		}),
	)

	// Create default parameters
	c.Parameters = new(RequestParameters)
	c.Parameters.UserAgent = DefaultUserAgent
	c.Parameters.AdvertiserSecretKey = advertiserSecretKey
	c.Parameters.Environment = LiveEnvironment
	if len(advertiserSecretKey) == 0 {
		err = fmt.Errorf("parameter %s cannot be empty", "advertiserSecretKey")
		return
	}

	// Create a last request struct
	c.LastRequest = new(LastRequest)

	// Return the client
	return
}

// Request is a generic TonicPow request wrapper that can be used without constraints
func (c *Client) Request(endpoint string, method string, params *url.Values) (response string, err error) {

	// Set reader
	var bodyReader io.Reader

	// Add the network value
	endpoint = fmt.Sprintf("%s%s", c.Parameters.Environment, endpoint)

	// Switch on POST vs GET
	switch method {
	case "POST":
		{
			encodedParams := params.Encode()
			bodyReader = strings.NewReader(encodedParams)
			c.LastRequest.PostData = encodedParams
		}
	case "GET":
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

	// Change the header (user agent is in case they block default Go user agents)
	request.Header.Set("User-Agent", c.Parameters.UserAgent)

	// Set the content type on POST
	if method == "POST" {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	// Fire the http request
	var resp *http.Response
	if resp, err = c.HTTPClient.Do(request); err != nil {
		return
	}

	// Close the response body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("error closing response body: %s", err.Error())
		}
	}()

	// Read the body
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	// Save the status
	c.LastRequest.StatusCode = resp.StatusCode

	// Parse the response
	response = string(body)

	// Done
	return
}

// ConvertGoal fires a conversion on a given goal name
func (c *Client) ConvertGoal(goalName string, sessionTxID string, userID string, additionalData string) (response *ConversionResponse, err error) {

	// Start the post data
	postData := url.Values{}

	// Add the key
	postData.Add("private_guid", c.Parameters.AdvertiserSecretKey)

	// Add the goal name
	postData.Add("conversion_goal_name", goalName)
	if len(goalName) == 0 {
		err = fmt.Errorf("parameter %s cannot be empty", "goalName")
		return
	}

	// Add the session/click
	postData.Add("click_tx_id", sessionTxID)
	if len(sessionTxID) == 0 {
		err = fmt.Errorf("parameter %s cannot be empty", "sessionTxID")
		return
	}

	// Add the user id if not found
	if len(userID) > 0 {
		postData.Add("user_id", userID)
	}

	// Add the additional data if found
	if len(additionalData) > 0 {
		postData.Add("additional_data", additionalData)
	}

	// Fire the request
	var resp string
	resp, err = c.Request("conversions", "POST", &postData)
	if err != nil {
		return
	}

	// Convert the response
	response = new(ConversionResponse)
	if err = json.Unmarshal([]byte(resp), response); err != nil {
		return
	}

	// Internal error from API
	if len(response.Message) > 0 && response.Code > 201 {
		err = fmt.Errorf("error from TonicPow API - code: %d message: %s", response.Code, response.Message)
	}
	return
}
