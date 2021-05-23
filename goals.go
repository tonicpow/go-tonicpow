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
//
// For more information: https://docs.tonicpow.com/#29a93e9b-9726-474c-b25e-92586200a803
func (c *Client) CreateGoal(goal *Goal) (createdGoal *Goal, err error) {

	// Basic requirements
	if goal.CampaignID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldCampaignID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(modelGoal, http.MethodPost, goal); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.Error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &createdGoal); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "goal"), http.StatusExpectationFailed)
	}
	return
}

// GetGoal will get an existing goal
// This will return an Error if the goal is not found (404)
//
// For more information: https://docs.tonicpow.com/#48d7bbc8-5d7b-4078-87b7-25f545c3deaf
func (c *Client) GetGoal(goalID uint64) (goal *Goal, err error) {

	// Must have an id
	if goalID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/details/%d", modelGoal, goalID), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &goal); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "goal"), http.StatusExpectationFailed)
	}
	return
}

// UpdateGoal will update an existing goal
//
// For more information: https://docs.tonicpow.com/#395f5b7d-6a5d-49c8-b1ae-abf7f90b42a2
func (c *Client) UpdateGoal(goal *Goal) (updatedGoal *Goal, err error) {

	// Basic requirements
	if goal.ID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldID), http.StatusBadRequest)
		return
	}

	// Permit fields
	goal.permitFields()

	// Fire the Request
	var response string
	if response, err = c.Request(modelGoal, http.MethodPut, goal); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &updatedGoal); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "goal"), http.StatusExpectationFailed)
	}
	return
}

// DeleteGoal will delete an existing goal
//
// For more information: https://docs.tonicpow.com/#38605b65-72c9-4fc8-87a7-bc644bc89a96
func (c *Client) DeleteGoal(goal *Goal) (deleted bool, err error) {

	// Basic requirements
	if goal.ID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(modelGoal+fmt.Sprintf("?id=%d", goal.ID), http.MethodDelete, goal); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	deleted = true
	return
}
