package tonicpow

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateLink will make a new link
//
// For more information: https://docs.tonicpow.com/#154bf9e1-6047-452f-a289-d21f507b0f1d
func (c *Client) CreateLink(link *Link) (createdLink *Link, err error) {

	// Basic requirements
	if link.CampaignID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldCampaignID), http.StatusBadRequest)
		return
	}

	if link.UserID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldUserID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(modelLink, http.MethodPost, link); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.Error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &createdLink); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "link"), http.StatusExpectationFailed)
	}
	return
}

// CreateLinkByURL will make a new link
//
// For more information: https://docs.tonicpow.com/#d5a22343-c580-43cc-8e27-dd131896ea3b
func (c *Client) CreateLinkByURL(link *Link) (createdLink *Link, err error) {

	// Basic requirements
	if len(link.TargetURL) == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldTargetURL), http.StatusBadRequest)
		return
	}

	if link.UserID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldUserID), http.StatusBadRequest)
		return
	}

	// Force campaign ID to zero
	link.CampaignID = 0

	// Fire the Request
	var response string
	if response, err = c.Request(modelLink, http.MethodPost, link); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.Error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &createdLink); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "link"), http.StatusExpectationFailed)
	}
	return
}

// GetLink will get an existing link
// This will return an Error if the link is not found (404)
//
// For more information: https://docs.tonicpow.com/#c53add03-303e-4f72-8847-2adfdb992eb3
func (c *Client) GetLink(linkID uint64) (link *Link, err error) {

	// Must have an id
	if linkID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/details/%d", modelLink, linkID), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &link); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "link"), http.StatusExpectationFailed)
	}
	return
}

// CheckLink will check for an existing link with a short_code
// This will return an Error if the link is not found (404)
//
// For more information: https://docs.tonicpow.com/#cc9780b7-0d84-4a60-a28f-664b2ecb209b
func (c *Client) CheckLink(shortCode string) (link *Link, err error) {

	// Must have a short code
	if len(shortCode) == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldShortCode), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/check/%s", modelLink, shortCode), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &link); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "link"), http.StatusExpectationFailed)
	}
	return
}

// ListLinksByUserID will get links associated to the user id
// This will return an Error if the link(s) are not found (404)
//
// For more information: https://docs.tonicpow.com/#23d068f1-4f0e-476a-a802-50b7edccd0b2
func (c *Client) ListLinksByUserID(userID uint64, page, resultsPerPage int) (results *LinkResults, err error) {

	// Must have an id
	if userID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldUserID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/links/?id=%d&%s=%d&%s=%d", modelUser, userID, fieldCurrentPage, page, fieldResultsPerPage, resultsPerPage), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &results); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "link"), http.StatusExpectationFailed)
	}
	return
}
