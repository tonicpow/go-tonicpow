package tonicpow

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateWidget will make a new widget
//
// For more information: (todo)
func (c *Client) CreateWidget(widget *Widget) (createdWidget *Widget, err error) {

	// Basic requirements
	if len(widget.Label) == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldLabel), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(modelWidget, http.MethodPost, widget); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.Error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &createdWidget); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "widget"), http.StatusExpectationFailed)
	}
	return
}

// GetWidget will get an existing widget
// This will return an Error if the profile is not found (404)
//
// For more information: (todo)
func (c *Client) GetWidget(widgetID uint64) (widget *Widget, err error) {

	// Must have an id
	if widgetID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldWidgetID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/details/%d", modelWidget, widgetID), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &widget); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "widget"), http.StatusExpectationFailed)
	}
	return
}

// UpdateWidget will update an existing widget
//
// For more information: (todo)
func (c *Client) UpdateWidget(widget *Widget) (updatedWidget *Widget, err error) {

	// Basic requirements
	if widget.ID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldWidgetID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(modelWidget, http.MethodPut, widget); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &updatedWidget); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "widget"), http.StatusExpectationFailed)
	}
	return
}
