package tonicpow

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// permitFields will remove fields that cannot be used
func (g *Goal) permitFields() {
	g.CampaignID = 0
	g.Payouts = 0
}

// CreateGoal will make a new goal
// Use the userSessionToken if making request on behalf of another user
//
// For more information: https://docs.tonicpow.com/#29a93e9b-9726-474c-b25e-92586200a803
func (c *Client) CreateGoal(goal *Goal, userSessionToken string) (createdGoal *Goal, err error) {

	// Basic requirements
	if goal.CampaignID == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldAdvertiserProfileID)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(modelGoal, http.MethodPost, goal, userSessionToken); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	createdGoal = new(Goal)
	err = json.Unmarshal([]byte(response), createdGoal)
	return
}

// GetGoal will get an existing goal
// This will return an error if the goal is not found (404)
// Use the userSessionToken if making request on behalf of another user
//
// For more information: https://docs.tonicpow.com/#48d7bbc8-5d7b-4078-87b7-25f545c3deaf
func (c *Client) GetGoal(goalID uint64, userSessionToken string) (goal *Goal, err error) {

	// Must have an id
	if goalID == 0 {
		err = fmt.Errorf("missing field: %s", fieldID)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/details/%d", modelGoal, goalID), http.MethodGet, nil, userSessionToken); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	goal = new(Goal)
	err = json.Unmarshal([]byte(response), goal)
	return
}

// UpdateGoal will update an existing goal
// Use the userSessionToken if making request on behalf of another user
//
// For more information: https://docs.tonicpow.com/#395f5b7d-6a5d-49c8-b1ae-abf7f90b42a2
func (c *Client) UpdateGoal(goal *Goal, userSessionToken string) (updatedGoal *Goal, err error) {

	// Basic requirements
	if goal.ID == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldID)
		return
	}

	// Permit fields
	goal.permitFields()

	// Fire the request
	var response string
	if response, err = c.request(modelGoal, http.MethodPut, goal, userSessionToken); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	updatedGoal = new(Goal)
	err = json.Unmarshal([]byte(response), updatedGoal)
	return
}

// ConvertGoal will fire a conversion for a given goal, if successful it will make a new Conversion
//
// For more information: https://docs.tonicpow.com/#caeffdd5-eaad-4fc8-ac01-8288b50e8e27
func (c *Client) ConvertGoal(goalName, visitorSessionID, additionalData, customUserID string) (conversion *Conversion, err error) {

	// Must have a name
	if len(goalName) == 0 {
		err = fmt.Errorf("missing field: %s", fieldName)
		return
	}

	// Must have a session id
	if len(visitorSessionID) == 0 {
		err = fmt.Errorf("missing field: %s", fieldVisitorSessionID)
		return
	}

	// Start the post data
	data := map[string]string{fieldName: goalName, fieldVisitorSessionID: visitorSessionID, fieldAdditionalData: additionalData, fieldUserID: customUserID}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/convert", modelGoal), http.MethodPost, data, ""); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	conversion = new(Conversion)
	err = json.Unmarshal([]byte(response), conversion)
	return
}
