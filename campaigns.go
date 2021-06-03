package tonicpow

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// permitFields will remove fields that cannot be used
func (c *Campaign) permitFields() {
	c.AdvertiserProfileID = 0
	c.Balance = 0
	c.BalanceSatoshis = 0
	c.PaidClicks = 0
	c.FundingAddress = ""
	c.LinksCreated = 0
	c.PublicGUID = ""
}

// CreateCampaign will make a new campaign
//
// For more information: https://docs.tonicpow.com/#b67e92bf-a481-44f6-a31d-26e6e0c521b1
func (c *Client) CreateCampaign(campaign *Campaign) (createdCampaign *Campaign, err error) {

	// Basic requirements
	if campaign.AdvertiserProfileID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldAdvertiserProfileID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(modelCampaign, http.MethodPost, campaign); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.Error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &createdCampaign); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "campaign"), http.StatusExpectationFailed)
	}
	return
}

// GetCampaign will get an existing campaign
// This will return an Error if the campaign is not found (404)
//
// For more information: https://docs.tonicpow.com/#b827446b-be34-4678-b347-33c4f63dbf9e
func (c *Client) GetCampaign(campaignID uint64) (campaign *Campaign, err error) {

	// Must have an id
	if campaignID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/details/?id=%d", modelCampaign, campaignID), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &campaign); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "campaign"), http.StatusExpectationFailed)
	}
	return
}

// GetCampaignBySlug will get an existing campaign
// This will return an Error if the campaign is not found (404)
//
// For more information: https://docs.tonicpow.com/#b827446b-be34-4678-b347-33c4f63dbf9e
func (c *Client) GetCampaignBySlug(slug string) (campaign *Campaign, err error) {

	// Must have an id
	if len(slug) == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldSlug), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/details/?slug=%s", modelCampaign, slug), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &campaign); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "campaign"), http.StatusExpectationFailed)
	}
	return
}

// UpdateCampaign will update an existing campaign
//
// For more information: https://docs.tonicpow.com/#665eefd6-da42-4ca9-853c-fd8ca1bf66b2
func (c *Client) UpdateCampaign(campaign *Campaign) (updatedCampaign *Campaign, err error) {

	// Basic requirements
	if campaign.ID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldID), http.StatusBadRequest)
		return
	}

	// Permit fields
	campaign.permitFields()

	// Fire the Request
	var response string
	if response, err = c.Request(modelCampaign, http.MethodPut, campaign); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &updatedCampaign); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "campaign"), http.StatusExpectationFailed)
	}
	return
}

// ListCampaigns will return a list of active campaigns
// This will return an Error if the campaign is not found (404)
//
// For more information: https://docs.tonicpow.com/#c1b17be6-cb10-48b3-a519-4686961ff41c
func (c *Client) ListCampaigns(page, resultsPerPage int, sortBy, sortOrder, searchQuery string,
	minimumBalance uint64, includeExpired bool) (results *CampaignResults, err error) {

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
	if response, err = c.Request(fmt.Sprintf(
		"%s/list?%s=%d&%s=%d&%s=%s&%s=%s&%s=%s&%s=%d&%s=%t",
		modelCampaign,
		fieldCurrentPage, page,
		fieldResultsPerPage, resultsPerPage,
		fieldSortBy, sortBy,
		fieldSortOrder, sortOrder,
		fieldSearchQuery, searchQuery,
		fieldMinimumBalance, minimumBalance,
		fieldExpired, includeExpired,
	), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &results); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "campaign"), http.StatusExpectationFailed)
	}
	return
}

// ListCampaignsByURL will return a list of active campaigns
// This will return an Error if the url is not found (404)
//
// todo: update with list campaigns functionality
// For more information: https://docs.tonicpow.com/#30a15b69-7912-4e25-ba41-212529fba5ff
func (c *Client) ListCampaignsByURL(targetURL string, page, resultsPerPage int,
	sortBy, sortOrder string) (results *CampaignResults, err error) {

	// Must have a value
	if len(targetURL) == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldTargetURL), http.StatusBadRequest)
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
	if response, err = c.Request(
		fmt.Sprintf("%s/find?%s=%s&%s=%d&%s=%d&%s=%s&%s=%s",
			modelCampaign, fieldTargetURL, targetURL, fieldCurrentPage,
			page, fieldResultsPerPage, resultsPerPage,
			fieldSortBy, sortBy,
			fieldSortOrder, sortOrder,
		), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &results); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "campaign"), http.StatusExpectationFailed)
	}
	return
}

// CampaignsFeed will return a feed of active campaigns
// This will return an Error if no campaigns are found (404)
//
// For more information: https://docs.tonicpow.com/#b3fe69d3-24ba-4c2a-a485-affbb0a738de
func (c *Client) CampaignsFeed(feedType string) (feed string, err error) {

	// Must have a value (force to rss if invalid)
	feedType = strings.ToLower(strings.TrimSpace(feedType))
	if len(feedType) == 0 || (feedType != FeedTypeRSS && feedType != FeedTypeAtom && feedType != FeedTypeJSON) {
		// Default feed type
		feedType = FeedTypeRSS
	}

	// Fire the Request
	if feed, err = c.Request(
		fmt.Sprintf("%s/feed?%s=%s", modelCampaign, fieldFeedType, feedType),
		http.MethodGet, nil,
	); err != nil {
		return
	}

	// Only a 200 is treated as a success
	err = c.Error(http.StatusOK, feed)

	return
}
