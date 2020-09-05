package tonicpow

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// permitFields will remove fields that cannot be used
func (a *AdvertiserProfile) permitFields() {
	a.UserID = 0
}

// CreateAdvertiserProfile will make a new advertiser profile
//
// For more information: https://docs.tonicpow.com/#153c0b65-2d4c-4972-9aab-f791db05b37b
func (c *Client) CreateAdvertiserProfile(profile *AdvertiserProfile) (createdProfile *AdvertiserProfile, err error) {

	// Basic requirements
	if profile.UserID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldUserID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(modelAdvertiser, http.MethodPost, profile); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.Error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &createdProfile); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "advertiser"), http.StatusExpectationFailed)
	}
	return
}

// GetAdvertiserProfile will get an existing advertiser profile
// This will return an Error if the profile is not found (404)
//
// For more information: https://docs.tonicpow.com/#b3a62d35-7778-4314-9321-01f5266c3b51
func (c *Client) GetAdvertiserProfile(profileID uint64) (profile *AdvertiserProfile, err error) {

	// Must have an id
	if profileID == 0 {
		err = c.createError(fmt.Sprintf("missing field: %s", fieldID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/details/%d", modelAdvertiser, profileID), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &profile); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "advertiser"), http.StatusExpectationFailed)
	}
	return
}

// UpdateAdvertiserProfile will update an existing profile
//
// For more information: https://docs.tonicpow.com/#0cebd1ff-b1ce-4111-aff6-9d586f632a84
func (c *Client) UpdateAdvertiserProfile(profile *AdvertiserProfile) (updatedProfile *AdvertiserProfile, err error) {

	// Basic requirements
	if profile.ID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldID), http.StatusBadRequest)
		return
	}

	// Permit fields
	profile.permitFields()

	// Fire the Request
	var response string
	if response, err = c.Request(modelAdvertiser, http.MethodPut, profile); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &updatedProfile); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "advertiser"), http.StatusExpectationFailed)
	}
	return
}

// ListCampaignsByAdvertiserProfile will return a list of campaigns
// This will return an Error if the campaign is not found (404)
//
// For more information: https://docs.tonicpow.com/#98017e9a-37dd-4810-9483-b6c400572e0c
func (c *Client) ListCampaignsByAdvertiserProfile(profileID uint64, page, resultsPerPage int, sortBy, sortOrder string) (results *CampaignResults, err error) {

	// Basic requirements
	if profileID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldAdvertiserProfileID), http.StatusBadRequest)
		return
	}

	// Do we know this field?
	if len(sortBy) > 0 {
		if !isInList(strings.ToLower(sortBy), campaignSortFields) {
			err = c.createError(fmt.Sprintf("sort by %s is not valid", sortBy), http.StatusBadRequest)
			return
		}
	} else {
		sortBy = SortByFieldCreatedAt
		sortOrder = SortOrderDesc
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/campaigns/%d?%s=%d&%s=%d&%s=%s&%s=%s", modelAdvertiser, profileID, fieldCurrentPage, page, fieldResultsPerPage, resultsPerPage, fieldSortBy, sortBy, fieldSortOrder, sortOrder), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &results); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "advertiser"), http.StatusExpectationFailed)
	}
	return
}

// ListAppsByAdvertiserProfile will return a list of apps
// This will return an Error if the campaign is not found (404)
//
// For more information: (todo)
func (c *Client) ListAppsByAdvertiserProfile(profileID uint64, page, resultsPerPage int, sortBy, sortOrder string) (results *AppResults, err error) {

	// Basic requirements
	if profileID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldAdvertiserProfileID), http.StatusBadRequest)
		return
	}

	// Do we know this field?
	if len(sortBy) > 0 {
		if !isInList(strings.ToLower(sortBy), appSortFields) {
			err = c.createError(fmt.Sprintf("sort by %s is not valid", sortBy), http.StatusBadRequest)
			return
		}
	} else {
		sortBy = SortByFieldCreatedAt
		sortOrder = SortOrderDesc
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/apps/%d?%s=%d&%s=%d&%s=%s&%s=%s", modelAdvertiser, profileID, fieldCurrentPage, page, fieldResultsPerPage, resultsPerPage, fieldSortBy, sortBy, fieldSortOrder, sortOrder), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &results); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "advertiser"), http.StatusExpectationFailed)
	}
	return
}
