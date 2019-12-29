package tonicpow

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateLink will make a new link
// Use the userSessionToken if making request on behalf of another user
//
// For more information: https://docs.tonicpow.com/#154bf9e1-6047-452f-a289-d21f507b0f1d
func (c *Client) CreateLink(link *Link, userSessionToken string) (createdLink *Link, err error) {

	// Basic requirements
	if link.CampaignID == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldCampaignID)
		return
	}

	if link.UserID == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldUserID)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(modelLink, http.MethodPost, link, userSessionToken); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &createdLink)
	return
}

// GetLink will get an existing link
// This will return an error if the link is not found (404)
// Use the userSessionToken if making request on behalf of another user
//
// For more information: https://docs.tonicpow.com/#c53add03-303e-4f72-8847-2adfdb992eb3
func (c *Client) GetLink(linkID uint64, userSessionToken string) (link *Link, err error) {

	// Must have an id
	if linkID == 0 {
		err = fmt.Errorf("missing field: %s", fieldID)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/details/%d", modelLink, linkID), http.MethodGet, nil, userSessionToken); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &link)
	return
}

// CheckLink will check for an existing link with a short_code
// This will return an error if the link is not found (404)
//
// For more information: https://docs.tonicpow.com/#cc9780b7-0d84-4a60-a28f-664b2ecb209b
func (c *Client) CheckLink(shortCode string) (link *Link, err error) {

	// Must have a short code
	if len(shortCode) == 0 {
		err = fmt.Errorf("missing field: %s", fieldShortCode)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/check/%s", modelLink, shortCode), http.MethodGet, nil, ""); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &link)
	return
}
