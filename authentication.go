package tonicpow

import (
	"net/http"
)

// createSession will establish a new session with the api
// This is run in the NewClient() method
//
// For more information: https://docs.tonicpow.com/#632ed94a-3afd-4323-af91-bdf307a399d2
func (c *Client) createSession() (err error) {

	// Start the post data with api key
	data := map[string]string{APIKeyName: c.Parameters.apiKey}

	// Fire the request
	var response string
	if response, err = c.request("auth/session", http.MethodPost, data, ""); err != nil {
		return
	}

	// Only a 201 is treated as a success
	err = c.error(http.StatusCreated, response)
	return
}

// ProlongSession will a session alive based on the forUser (user vs api session)
// Use customSessionToken for any token, user token, if empty it will use current api session token
//
// For more information: https://docs.tonicpow.com/#632ed94a-3afd-4323-af91-bdf307a399d2
func (c *Client) ProlongSession(customSessionToken string) (err error) {

	// Fire the request
	var response string
	if response, err = c.request("auth/session", http.MethodGet, nil, customSessionToken); err != nil {
		return
	}

	// Only a 200 is treated as a success
	err = c.error(http.StatusOK, response)
	return
}

// EndSession will end a session based on the forUser (user vs api session)
// Use customSessionToken for any token, user token, if empty it will use current api session token
//
// For more information: https://docs.tonicpow.com/#632ed94a-3afd-4323-af91-bdf307a399d2
func (c *Client) EndSession(customSessionToken string) (err error) {

	// Fire the request
	var response string
	if response, err = c.request("auth/session", http.MethodDelete, nil, customSessionToken); err != nil {
		return
	}

	// Only a 200 is treated as a success
	err = c.error(http.StatusOK, response)
	return
}
