package tonicpow

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateApp will make a new app
//
// For more information: (todo)
func (c *Client) CreateApp(app *App) (createdApp *App, err error) {

	// Basic requirements
	if app.AdvertiserProfileID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldAdvertiserProfileID), http.StatusBadRequest)
		return
	} else if len(app.Name) == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldName), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(modelApp, http.MethodPost, app); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.Error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &createdApp); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "app"), http.StatusExpectationFailed)
	}
	return
}

// GetApp will get an existing app
// This will return an Error if the profile is not found (404)
//
// For more information: (todo)
func (c *Client) GetApp(appID uint64) (app *App, err error) {

	// Must have an id
	if appID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldAppID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/details/%d", modelApp, appID), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &app); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "app"), http.StatusExpectationFailed)
	}
	return
}

// UpdateApp will update an existing app
//
// For more information: (todo)
func (c *Client) UpdateApp(app *App) (updatedApp *App, err error) {

	// Basic requirements
	if app.ID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldAppID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(modelApp, http.MethodPut, app); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &updatedApp); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "app"), http.StatusExpectationFailed)
	}
	return
}
