package tonicpow

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// permitFields will remove fields that cannot be used
func (c *Campaign) permitFields() {
	c.AdvertiserProfileID = 0
	c.Balance = 0
	c.BalanceSatoshis = 0
	c.Clicks = 0
	c.FundingAddress = ""
	c.LinksCreated = 0
	c.PublicGUID = ""
}

// CreateCampaign will make a new campaign
// Use the userSessionToken if making request on behalf of another user
//
// For more information: https://docs.tonicpow.com/#b67e92bf-a481-44f6-a31d-26e6e0c521b1
func (c *Client) CreateCampaign(campaign *Campaign, userSessionToken string) (createdCampaign *Campaign, err error) {

	// Basic requirements
	if campaign.AdvertiserProfileID == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldAdvertiserProfileID)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(modelCampaign, http.MethodPost, campaign, userSessionToken); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &createdCampaign)
	return
}

// GetCampaign will get an existing campaign
// This will return an error if the campaign is not found (404)
// Use the userSessionToken if making request on behalf of another user
//
// For more information: https://docs.tonicpow.com/#b827446b-be34-4678-b347-33c4f63dbf9e
func (c *Client) GetCampaign(campaignID uint64, userSessionToken string) (campaign *Campaign, err error) {

	// Must have an id
	if campaignID == 0 {
		err = fmt.Errorf("missing field: %s", fieldID)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/details/%d", modelCampaign, campaignID), http.MethodGet, nil, userSessionToken); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &campaign)
	return
}

// GetCampaignBalance will update the models's balance from the chain
// This will return an error if the campaign is not found (404)
//
// For more information: https://docs.tonicpow.com/#b6c60c63-8ac5-4c74-a4a2-cf3e858e5a8d
func (c *Client) GetCampaignBalance(campaignID uint64, lastBalance int64) (campaign *Campaign, err error) {

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/balance/%d?%s=%d", modelCampaign, campaignID, fieldLastBalance, lastBalance), http.MethodGet, nil, ""); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &campaign)
	return
}

// UpdateCampaign will update an existing campaign
// Use the userSessionToken if making request on behalf of another user
//
// For more information: https://docs.tonicpow.com/#665eefd6-da42-4ca9-853c-fd8ca1bf66b2
func (c *Client) UpdateCampaign(campaign *Campaign, userSessionToken string) (updatedCampaign *Campaign, err error) {

	// Basic requirements
	if campaign.ID == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldID)
		return
	}

	// Permit fields
	campaign.permitFields()

	// Fire the request
	var response string
	if response, err = c.request(modelCampaign, http.MethodPut, campaign, userSessionToken); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &updatedCampaign)
	return
}

// ListCampaigns will return a list of active campaigns
// Use the customSessionToken if making request on behalf of another user / advertiser
// This will return an error if the campaign is not found (404)
//
// For more information: https://docs.tonicpow.com/#c1b17be6-cb10-48b3-a519-4686961ff41c
func (c *Client) ListCampaigns(customSessionToken string, page, resultsPerPage int) (results *CampaignResults, err error) {

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/list?%s=%d&%s=%d", modelCampaign, fieldCurrentPage, page, fieldResultsPerPage, resultsPerPage), http.MethodGet, nil, customSessionToken); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &results)
	return
}

// ListCampaignsByURL will return a list of active campaigns
// This will return an error if the url is not found (404)
//
// For more information: https://docs.tonicpow.com/#30a15b69-7912-4e25-ba41-212529fba5ff
func (c *Client) ListCampaignsByURL(targetURL string, page, resultsPerPage int) (results *CampaignResults, err error) {

	// Must have a value
	if len(targetURL) == 0 {
		err = fmt.Errorf("missing field: %s", fieldTargetURL)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/find?%s=%s&%s=%d&%s=%d", modelCampaign, fieldTargetURL, targetURL, fieldCurrentPage, page, fieldResultsPerPage, resultsPerPage), http.MethodGet, nil, ""); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &results)
	return
}
