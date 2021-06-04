package tonicpow

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// newTestProfile creates a dummy profile for testing
func newTestProfile() *AdvertiserProfile {
	return &AdvertiserProfile{
		HomepageURL: "https://tonicpow.com",
		IconURL:     "https://i.imgur.com/HvVmeWI.png",
		PublicGUID:  "a4503e16b25c29b9cf58eee3ad353410",
		Name:        "TonicPow",
		ID:          23,
		UserID:      43,
	}
}

// TestClient_GetAdvertiserProfile will test the method GetAdvertiserProfile()
func TestClient_GetAdvertiserProfile(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	t.Run("get an advertiser (success)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		profile := newTestProfile()
		profile.Name = "TonicPow Test"

		endpoint := fmt.Sprintf("%s/%s/details/%d", EnvironmentDevelopment.apiURL, modelAdvertiser, profile.ID)

		err = mockResponseData(http.MethodGet, endpoint, http.StatusOK, profile)
		assert.NoError(t, err)

		profile, err = client.GetAdvertiserProfile(profile.ID)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, "TonicPow Test", profile.Name)
	})

	t.Run("missing advertiser profile id", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		profile := newTestProfile()
		profile.ID = 0

		endpoint := fmt.Sprintf("%s/%s/details/%d", EnvironmentDevelopment.apiURL, modelAdvertiser, profile.ID)

		err = mockResponseData(http.MethodGet, endpoint, http.StatusOK, profile)
		assert.NoError(t, err)

		var realProfile *AdvertiserProfile
		realProfile, err = client.GetAdvertiserProfile(profile.ID)
		assert.Error(t, err)
		assert.Nil(t, realProfile)
	})

	t.Run("error from api (status code)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		profile := newTestProfile()
		endpoint := fmt.Sprintf("%s/%s/details/%d", EnvironmentDevelopment.apiURL, modelAdvertiser, profile.ID)

		err = mockResponseData(http.MethodGet, endpoint, http.StatusBadRequest, profile)
		assert.NoError(t, err)

		var realProfile *AdvertiserProfile
		realProfile, err = client.GetAdvertiserProfile(profile.ID)
		assert.Error(t, err)
		assert.Nil(t, realProfile)
	})

	t.Run("error from api (api error)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		profile := newTestProfile()

		endpoint := fmt.Sprintf("%s/%s/details/%d", EnvironmentDevelopment.apiURL, modelAdvertiser, profile.ID)

		apiError := &Error{
			Code:        123,
			Data:        "field_name",
			IPAddress:   "127.0.0.1",
			Message:     "some error message",
			Method:      http.MethodPut,
			RequestGUID: "7f3d97a8fd67ff57861904df6118dcc8",
			StatusCode:  http.StatusBadRequest,
			URL:         endpoint,
		}

		err = mockResponseData(http.MethodGet, endpoint, http.StatusBadRequest, apiError)
		assert.NoError(t, err)

		var realProfile *AdvertiserProfile
		realProfile, err = client.GetAdvertiserProfile(profile.ID)
		assert.Error(t, err)
		assert.Nil(t, realProfile)
		assert.Equal(t, apiError.Message, err.Error())
	})
}

// ExampleClient_GetAdvertiserProfile example using GetAdvertiserProfile()
//
// See more examples in /examples/
func ExampleClient_GetAdvertiserProfile() {

	// Load the client (using test client for example only)
	client, err := newTestClient()
	if err != nil {
		fmt.Printf("error loading client: %s", err.Error())
		return
	}

	// For mocking
	responseProfile := newTestProfile()

	// Mock response (for example only)
	_ = mockResponseData(
		http.MethodGet,
		fmt.Sprintf("%s/%s/details/%d", EnvironmentDevelopment.apiURL, modelAdvertiser, responseProfile.ID),
		http.StatusOK,
		responseProfile,
	)

	// Get profile (using mocking response)
	var profile *AdvertiserProfile
	if profile, err = client.GetAdvertiserProfile(23); err != nil {
		fmt.Printf("error getting profile: " + err.Error())
		return
	}
	fmt.Printf("advertiser profile: %s", profile.Name)
	// Output:advertiser profile: TonicPow
}

// BenchmarkClient_GetAdvertiserProfile benchmarks the method GetAdvertiserProfile()
func BenchmarkClient_GetAdvertiserProfile(b *testing.B) {
	client, _ := newTestClient()
	profile := newTestProfile()
	_ = mockResponseData(
		http.MethodGet,
		fmt.Sprintf("%s/%s/details/%d", EnvironmentDevelopment.apiURL, modelAdvertiser, profile.ID),
		http.StatusOK,
		profile,
	)
	for i := 0; i < b.N; i++ {
		_, _ = client.GetAdvertiserProfile(profile.ID)
	}
}

// TestClient_UpdateAdvertiserProfile will test the method UpdateAdvertiserProfile()
func TestClient_UpdateAdvertiserProfile(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelAdvertiser)

	t.Run("update an advertiser (success)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		profile := newTestProfile()

		profile.Name = "TonicPow Test"
		err = mockResponseData(http.MethodPut, endpoint, http.StatusOK, profile)
		assert.NoError(t, err)

		err = client.UpdateAdvertiserProfile(profile)
		assert.NoError(t, err)
		assert.Equal(t, "TonicPow Test", profile.Name)
	})

	t.Run("missing advertiser profile id", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		profile := newTestProfile()

		profile.ID = 0
		err = mockResponseData(http.MethodPut, endpoint, http.StatusOK, profile)
		assert.NoError(t, err)

		err = client.UpdateAdvertiserProfile(profile)
		assert.Error(t, err)
	})

	t.Run("error from api (status code)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		profile := newTestProfile()

		err = mockResponseData(http.MethodPut, endpoint, http.StatusBadRequest, profile)
		assert.NoError(t, err)

		err = client.UpdateAdvertiserProfile(profile)
		assert.Error(t, err)
	})

	t.Run("error from api (api error)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		profile := newTestProfile()

		apiError := &Error{
			Code:        123,
			Data:        "field_name",
			IPAddress:   "127.0.0.1",
			Message:     "some error message",
			Method:      http.MethodPut,
			RequestGUID: "7f3d97a8fd67ff57861904df6118dcc8",
			StatusCode:  http.StatusBadRequest,
			URL:         fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelAdvertiser),
		}

		err = mockResponseData(http.MethodPut, endpoint, http.StatusBadRequest, apiError)
		assert.NoError(t, err)

		err = client.UpdateAdvertiserProfile(profile)
		assert.Error(t, err)
		assert.Equal(t, apiError.Message, err.Error())
	})
}

// ExampleClient_UpdateAdvertiserProfile example using UpdateAdvertiserProfile()
//
// See more examples in /examples/
func ExampleClient_UpdateAdvertiserProfile() {

	// Load the client (using test client for example only)
	client, err := newTestClient()
	if err != nil {
		fmt.Printf("error loading client: %s", err.Error())
		return
	}

	// Start with an existing profile
	profile := newTestProfile()
	profile.Name = "TonicPow Test"

	// Mock response (for example only)
	_ = mockResponseData(
		http.MethodPut,
		fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelAdvertiser),
		http.StatusOK,
		profile,
	)

	// Update profile
	err = client.UpdateAdvertiserProfile(profile)
	if err != nil {
		fmt.Printf("error updating profile: " + err.Error())
		return
	}
	fmt.Printf("profile updated: %s", profile.Name)
	// Output:profile updated: TonicPow Test
}

// BenchmarkClient_UpdateAdvertiserProfile benchmarks the method UpdateAdvertiserProfile()
func BenchmarkClient_UpdateAdvertiserProfile(b *testing.B) {
	client, _ := newTestClient()
	profile := newTestProfile()
	_ = mockResponseData(
		http.MethodPut,
		fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelAdvertiser),
		http.StatusOK,
		profile,
	)
	for i := 0; i < b.N; i++ {
		_ = client.UpdateAdvertiserProfile(profile)
	}
}

// mockResponseData is used for mocking the response
func mockResponseData(method, endpoint string, statusCode int, model interface{}) error {
	httpmock.Reset()
	data, err := json.Marshal(model)
	if err != nil {
		return err
	}
	httpmock.RegisterResponder(method, endpoint, httpmock.NewStringResponder(statusCode, string(data)))
	return nil
}
