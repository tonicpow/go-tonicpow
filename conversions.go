package tonicpow

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateConversionByGoalID will fire a conversion for a given goal id, if successful it will make a new Conversion
//
// For more information: https://docs.tonicpow.com/#caeffdd5-eaad-4fc8-ac01-8288b50e8e27
func (c *Client) CreateConversionByGoalID(goalID uint64, tncpwSession, customDimensions string, optionalPurchaseAmount float64, delayInMinutes int64) (conversion *Conversion, err error) {

	// Must have a name
	if goalID == 0 {
		err = fmt.Errorf("missing field: %s", fieldID)
		return
	}

	// Must have a session guid
	if len(tncpwSession) == 0 {
		err = fmt.Errorf("missing field: %s", fieldVisitorSessionGUID)
		return
	}

	// Start the post data
	data := map[string]string{fieldGoalID: fmt.Sprintf("%d", goalID), fieldVisitorSessionGUID: tncpwSession, fieldCustomDimensions: customDimensions, fieldDelayInMinutes: fmt.Sprintf("%d", delayInMinutes), fieldAmount: fmt.Sprintf("%f", optionalPurchaseAmount)}

	// Fire the request
	var response string
	if response, err = c.request(modelConversion, http.MethodPost, data); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &conversion)
	return
}

// CreateConversionByGoalName will fire a conversion for a given goal name, if successful it will make a new Conversion
//
// For more information: https://docs.tonicpow.com/#d19c9850-3832-45b2-b880-3ef2f3b7dc37
func (c *Client) CreateConversionByGoalName(goalName, tncpwSession, customDimensions string, optionalPurchaseAmount float64, delayInMinutes int64) (conversion *Conversion, err error) {

	// Must have a name
	if len(goalName) == 0 {
		err = fmt.Errorf("missing field: %s", fieldName)
		return
	}

	// Must have a session guid
	if len(tncpwSession) == 0 {
		err = fmt.Errorf("missing field: %s", fieldVisitorSessionGUID)
		return
	}

	// Start the post data
	data := map[string]string{fieldName: goalName, fieldVisitorSessionGUID: tncpwSession, fieldCustomDimensions: customDimensions, fieldDelayInMinutes: fmt.Sprintf("%d", delayInMinutes), fieldAmount: fmt.Sprintf("%f", optionalPurchaseAmount)}

	// Fire the request
	var response string
	if response, err = c.request(modelConversion, http.MethodPost, data); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &conversion)
	return
}

// CreateConversionByUserID will fire a conversion for a given goal and user id, if successful it will make a new Conversion
//
// For more information: https://docs.tonicpow.com/#d724f762-329e-473d-bdc4-aebc19dd9ea8
func (c *Client) CreateConversionByUserID(goalID, userID uint64, customDimensions string, optionalPurchaseAmount float64, delayInMinutes int64) (conversion *Conversion, err error) {

	// Must have a name
	if goalID == 0 {
		err = fmt.Errorf("missing field: %s", fieldID)
		return
	}

	// Must have a user id
	if userID == 0 {
		err = fmt.Errorf("missing field: %s", fieldUserID)
		return
	}

	// Start the post data
	data := map[string]string{fieldGoalID: fmt.Sprintf("%d", goalID), fieldUserID: fmt.Sprintf("%d", userID), fieldCustomDimensions: customDimensions, fieldDelayInMinutes: fmt.Sprintf("%d", delayInMinutes), fieldAmount: fmt.Sprintf("%f", optionalPurchaseAmount)}

	// Fire the request
	var response string
	if response, err = c.request(modelConversion, http.MethodPost, data); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &conversion)
	return
}

// GetConversion will get an existing conversion
// This will return an error if the goal is not found (404)
//
// For more information: https://docs.tonicpow.com/#fce465a1-d8d5-442d-be22-95169170167e
func (c *Client) GetConversion(conversionID uint64) (conversion *Conversion, err error) {

	// Must have an id
	if conversionID == 0 {
		err = fmt.Errorf("missing field: %s", fieldID)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/details/%d", modelConversion, conversionID), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &conversion)
	return
}

// CancelConversion will cancel an existing conversion (if delay was set and > 1 minute remaining)
//
// For more information: https://docs.tonicpow.com/#e650b083-bbb4-4ff7-9879-c14b1ab3f753
func (c *Client) CancelConversion(conversionID uint64, cancelReason string) (conversion *Conversion, err error) {

	// Must have an id
	if conversionID == 0 {
		err = fmt.Errorf("missing field: %s", fieldID)
		return
	}

	// Start the post data
	data := map[string]string{fieldID: fmt.Sprintf("%d", conversionID), fieldReason: cancelReason}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/cancel", modelConversion), http.MethodPut, data); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &conversion)
	return
}
