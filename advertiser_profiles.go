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
//
// For more information: https://docs.tonicpow.com/#153c0b65-2d4c-4972-9aab-f791db05b37b
func (c *Client) CreateAdvertiserProfile(profile *AdvertiserProfile) (createdProfile *AdvertiserProfile, err error) {

	// Basic requirements
	if profile.UserID == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldUserID)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request("advertisers", http.MethodPost, profile, ""); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	createdProfile = new(AdvertiserProfile)
	err = json.Unmarshal([]byte(response), createdProfile)
	return
}

// GetAdvertiserProfile will get an existing advertiser profile
// This will return an error if the profile is not found (404)
//
// For more information: https://docs.tonicpow.com/#b3a62d35-7778-4314-9321-01f5266c3b51
func (c *Client) GetAdvertiserProfile(profileID uint64) (profile *AdvertiserProfile, err error) {

	// Must have an id
	if profileID == 0 {
		err = fmt.Errorf("missing field: %s", fieldID)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("advertisers/details/%d", profileID), http.MethodGet, nil, ""); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	profile = new(AdvertiserProfile)
	err = json.Unmarshal([]byte(response), profile)
	return
}
