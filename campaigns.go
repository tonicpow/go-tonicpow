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
	c.FundingAddress = ""
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
	if response, err = c.request("campaigns", http.MethodPost, campaign, userSessionToken); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	createdCampaign = new(Campaign)
	err = json.Unmarshal([]byte(response), createdCampaign)
	return
}
