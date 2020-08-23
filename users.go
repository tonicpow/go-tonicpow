package tonicpow

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
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
		err = fmt.Errorf("missing required attribute: %s", fieldEmail)
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
	err = json.Unmarshal([]byte(response), &createdUser)
	return
}

// UpdateUser will update an existing user model
//
// For more information: https://docs.tonicpow.com/#7c3c3c3a-f636-469f-a884-449cf6fb35fe
func (c *Client) UpdateUser(user *User) (updatedUser *User, err error) {

	// Basic requirements
	if user.ID == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldID)
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
	err = json.Unmarshal([]byte(response), &updatedUser)
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
	err = json.Unmarshal([]byte(response), &user)
	return
}

// GetUserBalance will update a user's balance from the chain
//
// For more information: https://docs.tonicpow.com/#8478765b-95b8-47ad-8b86-2db5bce54924
func (c *Client) GetUserBalance(userID uint64, lastBalance int64) (user *User, err error) {

	// Basic requirements
	if userID == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldID)
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
	err = json.Unmarshal([]byte(response), &user)
	return
}

// CurrentUser will the current user based on token
//
// For more information: https://docs.tonicpow.com/#7f6e9b5d-8c7f-4afc-8e07-7aafdd891521
func (c *Client) CurrentUser(userID uint64) (user *User, err error) {

	// No current user
	if userID == 0 {
		err = fmt.Errorf("missing user id")
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
	err = json.Unmarshal([]byte(response), &user)
	return
}

// LoginUser will authenticate a given user
//
// For more information: https://docs.tonicpow.com/#5cad3e9a-5931-44bf-b110-4c4b74c7a070
func (c *Client) LoginUser(email, password string) (user *User, err error) {

	// Basic requirements
	if len(email) == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldEmail)
		return
	} else if len(password) == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldPassword)
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
	err = json.Unmarshal([]byte(response), &user)
	return
}

// ForgotPassword will fire a forgot password Request
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

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/password/forgot", modelUser), http.MethodPost, data); err != nil {
		return
	}

	// Only a 200 is treated as a success
	err = c.Error(http.StatusOK, response)
	return
}

// ResetPassword will reset a password from a ForgotPassword() Request
//
// For more information: https://docs.tonicpow.com/#370fbeec-adb2-4ed3-82dc-2dffa840e490
func (c *Client) ResetPassword(token, password, passwordConfirm string) (err error) {

	// Basic requirements
	if len(token) == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldToken)
		return
	} else if len(token) <= 10 {
		err = fmt.Errorf("invalid token: %s", token)
		return
	} else if len(password) == 0 || len(passwordConfirm) == 0 {
		err = fmt.Errorf("missing required attribute: %s or %s", fieldPassword, fieldPasswordConfirm)
		return
	} else if password != passwordConfirm {
		err = fmt.Errorf("given passwords don't match")
		return
	}

	// Start the post data
	data := map[string]string{fieldToken: token, fieldPassword: password, fieldPasswordConfirm: passwordConfirm}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/password/reset", modelUser), http.MethodPut, data); err != nil {
		return
	}

	// Only a 200 is treated as a success
	err = c.Error(http.StatusOK, response)
	return
}

// AcceptUser will accept a user (if approval is required for new users)
// ID or email address
// Reason field is optional
//
// For more information: https://docs.tonicpow.com/#65c3962d-c309-4ef4-b85f-7ec1f08f031b
func (c *Client) AcceptUser(userID uint64, email string, reason string) (err error) {

	var data map[string]string

	// Basic requirements
	if userID > 0 {
		data = map[string]string{fieldID: fmt.Sprintf("%d", userID), fieldReason: reason}
	} else if len(email) > 0 {
		data = map[string]string{fieldEmail: email, fieldReason: reason}
	} else {
		err = fmt.Errorf("missing required attribute: %s or %s", fieldUserID, fieldEmail)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/status/accept", modelUser), http.MethodPut, data); err != nil {
		return
	}

	// Only a 200 is treated as a success
	err = c.Error(http.StatusOK, response)
	return
}

// ActivateUser will activate a user (if all application criteria is met)
// ID or email address
//
// For more information: https://docs.tonicpow.com/#aa499fdf-2492-43ee-99d4-fc9735676431
func (c *Client) ActivateUser(userID uint64, email string) (err error) {

	var data map[string]string

	// Basic requirements
	if userID > 0 {
		data = map[string]string{fieldID: fmt.Sprintf("%d", userID)}
	} else if len(email) > 0 {
		data = map[string]string{fieldEmail: email}
	} else {
		err = fmt.Errorf("missing required attribute: %s or %s", fieldUserID, fieldEmail)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/status/activate", modelUser), http.MethodPut, data); err != nil {
		return
	}

	// Only a 200 is treated as a success
	err = c.Error(http.StatusOK, response)
	return
}

// PauseUser will pause a user account (all payouts go to internal address)
// ID or email address
//
// For more information: https://docs.tonicpow.com/#3307310d-86a9-4a5c-84ff-c38c581c77e5
func (c *Client) PauseUser(userID uint64, email string, reason string) (err error) {

	var data map[string]string

	// Basic requirements
	if userID > 0 {
		data = map[string]string{fieldID: fmt.Sprintf("%d", userID), fieldReason: reason}
	} else if len(email) > 0 {
		data = map[string]string{fieldEmail: email, fieldReason: reason}
	} else {
		err = fmt.Errorf("missing required attribute: %s or %s", fieldUserID, fieldEmail)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/status/pause", modelUser), http.MethodPut, data); err != nil {
		return
	}

	// Only a 200 is treated as a success
	err = c.Error(http.StatusOK, response)
	return
}

// UserExists will check if a user exists by email address
// This will return an Error if the user is not found (404)
//
// For more information: https://docs.tonicpow.com/#2d8c37d4-c88b-4cec-83ad-fa72b0f41f17
func (c *Client) UserExists(byEmail string) (existsResponse *UserExists, err error) {

	// Must have email
	if len(byEmail) == 0 {
		err = fmt.Errorf("missing %s", fieldEmail)
		return
	}

	// Set the values
	params := url.Values{}
	params.Add(fieldEmail, byEmail)

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/exists", modelUser), http.MethodGet, params); err != nil {
		return
	}

	// Only a 200 is treated as a success, 404 is false and no Error
	if c.LastRequest.StatusCode == http.StatusNotFound {
		return
	}
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &existsResponse)
	return
}

// ReleaseUserBalance will send the internal balance to the user's payout_address
// Reason field is optional
//
// For more information: https://docs.tonicpow.com/#be82b6cb-7fe8-4f03-9b0c-dbade8f2d40f
func (c *Client) ReleaseUserBalance(userID uint64, reason string) (err error) {

	var data map[string]string

	// Basic requirements
	if userID > 0 {
		data = map[string]string{fieldID: fmt.Sprintf("%d", userID), fieldReason: reason}
	} else {
		err = fmt.Errorf("missing required attribute: %s", fieldUserID)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/wallet/release", modelUser), http.MethodPut, data); err != nil {
		return
	}

	// Only a 200 is treated as a success
	err = c.Error(http.StatusOK, response)
	return
}

// RefundUserBalance will send the internal balance back to the corresponding campaigns
// Reason field is required
//
// For more information: https://docs.tonicpow.com/#c373c7ed-189d-4aa6-88da-c4a58955fd28
func (c *Client) RefundUserBalance(userID uint64, reason string) (err error) {

	var data map[string]string

	// Basic requirements
	if userID > 0 {
		data = map[string]string{fieldID: fmt.Sprintf("%d", userID), fieldReason: reason}
	} else {
		err = fmt.Errorf("missing required attribute: %s", fieldUserID)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/wallet/refund", modelUser), http.MethodPut, data); err != nil {
		return
	}

	// Only a 200 is treated as a success
	err = c.Error(http.StatusOK, response)
	return
}

// GetUserReferrals will return all the related referrals to the given user
// Use either an ID or email to get an existing user
//
// For more information: https://docs.tonicpow.com/#fa7ee5a6-c87d-4e01-8ad3-ef6bda39533b
func (c *Client) GetUserReferrals(byID uint64, byEmail string) (referrals []*UserReferral, err error) {

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

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/referred", modelUser), http.MethodGet, params); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &referrals)
	return
}

// ListUserReferrals will return a list of active users that have referrals
// This will return an Error if no users are found (404)
//
// For more information: https://docs.tonicpow.com/#3fd8e647-abfa-422f-90af-952cccd3be7c
func (c *Client) ListUserReferrals(page, resultsPerPage int, sortBy, sortOrder string) (results *ReferralResults, err error) {

	// Do we know this field?
	if len(sortBy) > 0 {
		if !isInList(strings.ToLower(sortBy), referralSortFields) {
			err = fmt.Errorf("sort by %s is not valid", sortBy)
			return
		}
	} else {
		sortBy = SortByFieldCreatedAt
		sortOrder = SortOrderDesc
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/referrals?%s=%d&%s=%d&%s=%s&%s=%s", modelUser, fieldCurrentPage,
		page, fieldResultsPerPage, resultsPerPage, fieldSortBy, sortBy, fieldSortOrder, sortOrder), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.Error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &results)
	return
}

// RequestActivation will send a Request for activation
//
// For more information: https://docs.tonicpow.com/#c3d2f569-dc5e-4885-9701-a58522cb92cf
func (c *Client) RequestActivation(userID uint64) (err error) {

	// Basic requirements
	if userID == 0 {
		err = fmt.Errorf("missing required attribute: %s", fieldID)
		return
	}

	// Fire the Request
	var response string
	if response, err = c.Request(fmt.Sprintf("%s/status/Request?id=%d", modelUser, userID), http.MethodPut, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	err = c.Error(http.StatusOK, response)
	return
}
