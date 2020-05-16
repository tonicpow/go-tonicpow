/*
Package tonicpow is the official golang implementation for the TonicPow API
*/
package tonicpow

import (
	"fmt"
)

// NewClient creates a new client to submit requests pre-loaded with the API key
// This will establish a new session given the API key
//
// For more information: https://docs.tonicpow.com
func NewClient(apiKey string, environment APIEnvironment, clientOptions *Options) (c *Client, err error) {

	// apiKey is required
	if len(apiKey) == 0 {
		err = fmt.Errorf("parameter %s cannot be empty", fieldAPIKey)
		return
	}

	// Create a client using the given options
	c = createClient(clientOptions)

	// Set the default parameters
	c.Parameters.apiKey = apiKey
	c.Parameters.environment = environment

	// Start a new api session
	err = c.createSession()

	return
}
