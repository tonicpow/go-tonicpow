package tonicpow

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// permitFields will remove fields that cannot be used
func (u *User) permitFields() {
	u.Balance = 0
	u.InternalAddress = ""
	u.Status = ""
}

// CreateUser will make a new user
//
// For more information: https://docs.tonicpow.com/#8de84fb5-ba77-42cc-abb0-f3044cc871b6
func (c *Client) CreateUser(user *User) (createdUser *User, err error) {

	// Basic requirements
	if len(user.Email) == 0 {
		err = fmt.Errorf("missing required attribute: %s", "email")
		return
	}

	// Permit fields before creating
	user.permitFields()

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

	// Permit fields before updating
	user.permitFields()

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

// GetUser will get an existing user
// This will return an error if the user is not found (404)
// Use either an ID or email to get an existing user
//
// For more information: https://docs.tonicpow.com/#e6f764a2-5a91-4680-aa5e-03409dd878d8
func (c *Client) GetUser(byID uint64, byEmail string) (user *User, err error) {

	// Must have either an ID or email
	if byID == 0 && len(byEmail) == 0 {
		err = fmt.Errorf("missing id/email")
		return
	}

	// Set the values
	params := url.Values{}
	if byID > 0 {
		params.Add("id", fmt.Sprintf("%d", byID))
	} else {
		params.Add("email", byEmail)
	}

	// Fire the request
	var response string
	if response, err = c.request("users/details", http.MethodGet, params, ""); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	user = new(User)
	err = json.Unmarshal([]byte(response), user)
	return
}

// LoginUser will login for a given user
//
// For more information: https://docs.tonicpow.com/#5cad3e9a-5931-44bf-b110-4c4b74c7a070
func (c *Client) LoginUser(user *User) (sessionToken string, err error) {

	// Basic requirements
	if len(user.Email) == 0 {
		err = fmt.Errorf("missing required attribute: %s", "email")
		return
	} else if len(user.Password) == 0 {
		err = fmt.Errorf("missing required attribute: %s", "password")
		return
	}

	// Fire the request
	var response string
	if response, err = c.request("users/login", http.MethodPost, user, c.Parameters.apiSessionCookie.Value); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	sessionToken = c.Parameters.UserSessionCookie.Value
	return
}
