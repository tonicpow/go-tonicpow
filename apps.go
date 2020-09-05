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
		err = fmt.Errorf("missing required attribute: %s", fieldAdvertiserProfileID)
		return
	} else if len(app.Name) == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldName)
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
	err = json.Unmarshal([]byte(response), &createdApp)
	return
}

// GetApp will get an existing app
// This will return an Error if the profile is not found (404)
//
// For more information: (todo)
func (c *Client) GetApp(appID uint64) (app *App, err error) {

	// Must have an id
	if appID == 0 {
		err = fmt.Errorf("missing field: %s", fieldID)
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
	err = json.Unmarshal([]byte(response), &app)
	return
}

// UpdateApp will update an existing app
//
// For more information: (todo)
func (c *Client) UpdateApp(app *App) (updatedApp *App, err error) {

	// Basic requirements
	if app.ID == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldID)
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
	err = json.Unmarshal([]byte(response), &updatedApp)
	return
}
