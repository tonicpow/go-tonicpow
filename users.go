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
		err = fmt.Errorf("missing required attribute: %s", fieldEmail)
		return
	}

	// Permit fields before creating
	user.permitFields()

	// Fire the request
	var response string
	if response, err = c.request(modelUser, http.MethodPost, user, ""); err != nil {
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
// Use the userSessionToken if the current user is editing their own user model
//
// For more information: https://docs.tonicpow.com/#7c3c3c3a-f636-469f-a884-449cf6fb35fe
func (c *Client) UpdateUser(user *User, userSessionToken string) (updatedUser *User, err error) {

	// Basic requirements
	if user.ID == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldID)
		return
	}

	// Permit fields before updating
	user.permitFields()

	// Fire the request
	var response string
	if response, err = c.request(modelUser, http.MethodPut, user, userSessionToken); err != nil {
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
		err = fmt.Errorf("missing either %s or %s", fieldID, fieldEmail)
		return
	}

	// Set the values
	params := url.Values{}
	if byID > 0 {
		params.Add(fieldID, fmt.Sprintf("%d", byID))
	} else {
		params.Add(fieldEmail, byEmail)
	}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/details", modelUser), http.MethodGet, params, ""); err != nil {
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

// GetUserBalance will update a user's balance from the chain
//
// For more information: https://docs.tonicpow.com/#8478765b-95b8-47ad-8b86-2db5bce54924
func (c *Client) GetUserBalance(userID uint64) (user *User, err error) {

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/balance/%d", modelUser, userID), http.MethodGet, nil, ""); err != nil {
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

// CurrentUser will the current user based on token
// Required: LoginUser()
//
// For more information: https://docs.tonicpow.com/#7f6e9b5d-8c7f-4afc-8e07-7aafdd891521
func (c *Client) CurrentUser() (user *User, err error) {

	// No current user
	if c.Parameters.UserSessionCookie == nil {
		err = fmt.Errorf("missing user session, use LoginUser() first")
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/account", modelUser), http.MethodGet, nil, c.Parameters.UserSessionCookie.Value); err != nil {
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
func (c *Client) LoginUser(user *User) (userSessionToken string, err error) {

	// Basic requirements
	if len(user.Email) == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldEmail)
		return
	} else if len(user.Password) == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldPassword)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/login", modelUser), http.MethodPost, user, c.Parameters.apiSessionCookie.Value); err != nil {
		return
	}

	// Only a 201 is treated as a success
	if err = c.error(http.StatusCreated, response); err != nil {
		return
	}

	// Convert model response
	userSessionToken = c.Parameters.UserSessionCookie.Value
	return
}

// LogoutUser will logout a given session token
//
// For more information: https://docs.tonicpow.com/#39d65294-376a-4366-8f71-a02b08f9abdf
func (c *Client) LogoutUser(userSessionToken string) (err error) {

	// Basic requirements
	if len(userSessionToken) == 0 {
		err = fmt.Errorf("missing required attribute: %s", sessionCookie)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/logout", modelUser), http.MethodDelete, nil, userSessionToken); err != nil {
		return
	}
	// Only a 200 is treated as a success
	err = c.error(http.StatusOK, response)
	return
}

// ForgotPassword will fire a forgot password request
//
// For more information: https://docs.tonicpow.com/#2c33dae4-d6b1-4949-9e84-fb02157ab7cd
func (c *Client) ForgotPassword(emailAddress string) (err error) {

	// Basic requirements
	if len(emailAddress) == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldEmail)
		return
	}

	// Start the post data
	data := map[string]string{fieldEmail: emailAddress}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/password/forgot", modelUser), http.MethodPost, data, ""); err != nil {
		return
	}

	// Only a 200 is treated as a success
	err = c.error(http.StatusOK, response)
	return
}

// ResetPassword will reset a password from a ForgotPassword() request
//
// For more information: https://docs.tonicpow.com/#370fbeec-adb2-4ed3-82dc-2dffa840e490
func (c *Client) ResetPassword(token, password, passwordConfirm string) (err error) {

	// Basic requirements
	if len(token) == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldToken)
		return
	} else if len(password) == 0 || len(passwordConfirm) == 0 {
		err = fmt.Errorf("missing required attribute: %s or %s", fieldPassword, fieldPasswordConfirm)
		return
	}

	// Start the post data
	data := map[string]string{fieldToken: token, fieldPassword: password, fieldPasswordConfirm: passwordConfirm}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/password/reset", modelUser), http.MethodPut, data, ""); err != nil {
		return
	}

	// Only a 200 is treated as a success
	err = c.error(http.StatusOK, response)
	return
}

// CompleteEmailVerification will complete an email verification with a given token
//
// For more information: https://docs.tonicpow.com/#f5081800-a224-4f36-8014-94981f0bd55d
func (c *Client) CompleteEmailVerification(token string) (err error) {

	// Basic requirements
	if len(token) == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldToken)
		return
	}

	// Start the post data
	data := map[string]string{fieldToken: token}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/verify/%s", modelUser, fieldEmail), http.MethodPut, data, ""); err != nil {
		return
	}

	// Only a 200 is treated as a success
	err = c.error(http.StatusOK, response)
	return
}

// CompletePhoneVerification will complete a phone verification with a given code and number
//
// For more information: https://docs.tonicpow.com/#573403c4-b872-475d-ac04-de32a88ecd19
func (c *Client) CompletePhoneVerification(phone, code string) (err error) {

	// Basic requirements
	if len(phone) == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldPhone)
		return
	} else if len(code) == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldPhoneCode)
		return
	}

	// Start the post data
	data := map[string]string{fieldPhone: phone, fieldPhoneCode: code}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/verify/%s", modelUser, fieldPhone), http.MethodPut, data, ""); err != nil {
		return
	}

	// Only a 200 is treated as a success
	err = c.error(http.StatusOK, response)
	return
}
