package tonicpow

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// permitFields will remove fields that cannot be used
func (a *AdvertiserProfile) permitFields() {
	a.UserID = 0
}

// CreateAdvertiserProfile will make a new advertiser profile
// Use the userSessionToken if making request on behalf of another user
//
// For more information: https://docs.tonicpow.com/#153c0b65-2d4c-4972-9aab-f791db05b37b
func (c *Client) CreateAdvertiserProfile(profile *AdvertiserProfile, userSessionToken string) (createdProfile *AdvertiserProfile, err error) {

	// Basic requirements
	if profile.UserID == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldUserID)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(modelAdvertiser, http.MethodPost, profile, userSessionToken); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &createdProfile)
	return
}

// GetAdvertiserProfile will get an existing advertiser profile
// This will return an error if the profile is not found (404)
// Use the userSessionToken if making request on behalf of another user
//
// For more information: https://docs.tonicpow.com/#b3a62d35-7778-4314-9321-01f5266c3b51
func (c *Client) GetAdvertiserProfile(profileID uint64, userSessionToken string) (profile *AdvertiserProfile, err error) {

	// Must have an id
	if profileID == 0 {
		err = fmt.Errorf("missing field: %s", fieldID)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/details/%d", modelAdvertiser, profileID), http.MethodGet, nil, userSessionToken); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &profile)
	return
}

// UpdateAdvertiserProfile will update an existing profile
// Use the userSessionToken if making request on behalf of another user
//
// For more information: https://docs.tonicpow.com/#0cebd1ff-b1ce-4111-aff6-9d586f632a84
func (c *Client) UpdateAdvertiserProfile(profile *AdvertiserProfile, userSessionToken string) (updatedProfile *AdvertiserProfile, err error) {

	// Basic requirements
	if profile.ID == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldID)
		return
	}

	// Permit fields
	profile.permitFields()

	// Fire the request
	var response string
	if response, err = c.request(modelAdvertiser, http.MethodPut, profile, userSessionToken); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &updatedProfile)
	return
}

// GetCampaignsByAdvertiserProfile will return a list of campaigns
// This will return an error if the campaign is not found (404)
//
// For more information: https://docs.tonicpow.com/#98017e9a-37dd-4810-9483-b6c400572e0c
func (c *Client) GetCampaignsByAdvertiserProfile(profileID uint64) (campaigns []*Campaign, err error) {

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/campaigns/%d", modelAdvertiser, profileID), http.MethodGet, nil, ""); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &campaigns)
	return
}
