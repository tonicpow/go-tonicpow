package tonicpow

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateUser will make a new user
//
// For more information: https://docs.tonicpow.com/#8de84fb5-ba77-42cc-abb0-f3044cc871b6
func (c *Client) CreateUser(user *User) (createdUser *User, err error) {

	// Basic requirements
	if len(user.Email) == 0 {
		err = fmt.Errorf("missing required attribute: %s", "email")
		return
	}

	// Fire the request
	var response string
	if response, err = c.request("users", http.MethodPost, user, ""); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	createdUser = new(User)
	err = json.Unmarshal([]byte(response), createdUser)
	return
}

// UpdateUser will update an existing user model
// Use the customSessionToken if the current user is editing their own user model
//
// For more information: https://docs.tonicpow.com/#7c3c3c3a-f636-469f-a884-449cf6fb35fe
func (c *Client) UpdateUser(user *User, customSessionToken string) (updatedUser *User, err error) {

	// Basic requirements
	if user.ID == 0 {
		err = fmt.Errorf("missing required attribute: %s", "id")
		return
	}

	// Fire the request
	var response string
	if response, err = c.request("users", http.MethodPut, user, customSessionToken); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	updatedUser = new(User)
	err = json.Unmarshal([]byte(response), updatedUser)
	return
}
