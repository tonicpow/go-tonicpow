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
	u.Earned = 0
	u.InternalAddress = ""
	u.Status = ""
	u.ReferralLinkID = 0
	u.ReferralURL = ""
}

// CreateUser will make a new user
//
// For more information: https://docs.tonicpow.com/#8de84fb5-ba77-42cc-abb0-f3044cc871b6
func (c *Client) CreateUser(user *User) (createdUser *User, err error) {

	// Basic requirements
	if len(user.Email) == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldEmail), http.StatusBadRequest)
		return
	}

	// Permit fields before creating
	user.permitFields()

	// Fire the Request
	var response string
	if response, err = c.Request(modelUser, http.MethodPost, user); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.Error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &createdUser); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "user"), http.StatusExpectationFailed)
	}
	return
}

// UpdateUser will update an existing user model
//
// For more information: https://docs.tonicpow.com/#7c3c3c3a-f636-469f-a884-449cf6fb35fe
func (c *Client) UpdateUser(user *User) (updatedUser *User, err error) {

	// Basic requirements
	if user.ID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldID), http.StatusBadRequest)
		return
	}

	// Permit fields before updating
	user.permitFields()

	// Fire the Request
	var response string
	if response, err = c.Request(modelUser, http.MethodPut, user); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &updatedUser); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "user"), http.StatusExpectationFailed)
	}
	return
}

// GetUser will get an existing user
// This will return an Error if the user is not found (404)
// Use either an ID or email to get an existing user
//
// For more information: https://docs.tonicpow.com/#e6f764a2-5a91-4680-aa5e-03409dd878d8
func (c *Client) GetUser(byID uint64, byEmail string) (user *User, err error) {

	// Must have either an ID or email
	if byID == 0 && len(byEmail) == 0 {
		err = c.createError(fmt.Sprintf("missing either %s or %s", fieldID, fieldEmail), http.StatusBadRequest)
		return
	}

	// Set the values
	params := url.Values{}
	if byID > 0 {
		params.Add(fieldID, fmt.Sprintf("%d", byID))
	} else {
		params.Add(fieldEmail, byEmail)
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/details", modelUser), http.MethodGet, params); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &user); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "user"), http.StatusExpectationFailed)
	}
	return
}

// GetUserBalance will update a user's balance from the chain
//
// For more information: https://docs.tonicpow.com/#8478765b-95b8-47ad-8b86-2db5bce54924
func (c *Client) GetUserBalance(userID uint64, lastBalance int64) (user *User, err error) {

	// Basic requirements
	if userID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/balance/%d?%s=%d", modelUser, userID, fieldLastBalance, lastBalance),
		http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &user); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "user"), http.StatusExpectationFailed)
	}
	return
}

// CurrentUser will the current user based on token
//
// For more information: https://docs.tonicpow.com/#7f6e9b5d-8c7f-4afc-8e07-7aafdd891521
func (c *Client) CurrentUser(userID uint64) (user *User, err error) {

	// No current user
	if userID == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldID), http.StatusBadRequest)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/account?id=%d", modelUser, userID), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &user); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "user"), http.StatusExpectationFailed)
	}
	return
}

// LoginUser will authenticate a given user
//
// For more information: https://docs.tonicpow.com/#5cad3e9a-5931-44bf-b110-4c4b74c7a070
func (c *Client) LoginUser(email, password string) (user *User, err error) {

	// Basic requirements
	if len(email) == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldEmail), http.StatusBadRequest)
		return
	} else if len(password) == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldPassword), http.StatusBadRequest)
		return
	}

	// Set the fields
	user = &User{
		Email:    email,
		Password: password,
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/login", modelUser), http.MethodPost, user); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.Error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	if err = json.Unmarshal([]byte(response), &user); err != nil {
		err = c.createError(fmt.Sprintf("failed unmarshaling data: %s", "user"), http.StatusExpectationFailed)
	}
	return
}

// ForgotPassword will fire a forgot password Request
//
// For more information: https://docs.tonicpow.com/#2c33dae4-d6b1-4949-9e84-fb02157ab7cd
func (c *Client) ForgotPassword(emailAddress string) (err error) {

	// Basic requirements
	if len(emailAddress) == 0 {
		err = c.createError(fmt.Sprintf("missing required attribute: %s", fieldEmail), http.StatusBadRequest)
		return
	}

	// Start the post data
	data := map[string]string{fieldEmail: emailAddress}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/password/forgot", modelUser), http.MethodPost, data); err != nil {
		return
	}

	// Only a 200 is treated as a success
	err = c.Error(http.StatusOK, response)
	return
}
