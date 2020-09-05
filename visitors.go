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
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldLinkID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/%s", modelVisitors, modelVisitorSession), http.MethodPost, visitorSession); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.Error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &createdSession); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "visitor_session"), http.StatusExpectationFailed)
	}
	return
}

// GetVisitorSession will get a visitor session
// This will return an Error if the session is not found (404)
//
// For more information: https://docs.tonicpow.com/#cf560448-6dda-42a6-9051-136afabe78e6
func (c *Client) GetVisitorSession(visitorSessionGUID string) (visitorSession *VisitorSession, err error) {

	// Must have an id
	if len(visitorSessionGUID) == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldVisitorSessionGUID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/%s/details/%s", modelVisitors, modelVisitorSession, visitorSessionGUID), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &visitorSession); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "visitor_session"), http.StatusExpectationFailed)
	}
	return
}
