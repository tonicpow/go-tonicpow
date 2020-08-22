package tonicpow

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateVisitorSession will make a new session for a visitor (used for goal conversions)
//
// For more information: https://docs.tonicpow.com/#29a93e9b-9726-474c-b25e-92586200a803
func (c *Client) CreateVisitorSession(visitorSession *VisitorSession) (createdSession *VisitorSession, err error) {

	// Basic requirements
	if visitorSession.LinkID == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldLinkID)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/%s", modelVisitors, modelVisitorSession), http.MethodPost, visitorSession); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &createdSession)
	return
}

// GetVisitorSession will get a visitor session
// This will return an error if the session is not found (404)
//
// For more information: https://docs.tonicpow.com/#cf560448-6dda-42a6-9051-136afabe78e6
func (c *Client) GetVisitorSession(visitorSessionGUID string) (visitorSession *VisitorSession, err error) {

	// Must have an id
	if len(visitorSessionGUID) == 0 {
		err = fmt.Errorf("missing field: %s", fieldVisitorSessionGUID)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/%s/details/%s", modelVisitors, modelVisitorSession, visitorSessionGUID), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &visitorSession)
	return
}
