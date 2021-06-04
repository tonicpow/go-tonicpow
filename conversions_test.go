package tonicpow

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestConversion() *Conversion {
	return &Conversion{
		CampaignID:       testCampaignID,
		CustomDimensions: `{"some_field":"some_value"}`,
		GoalID:           testGoalID,
		GoalName:         testGoalName,
		ID:               testConversionID,
		Status:           "pending",
		UserID:           testUserID,
	}
}

// TestClient_CreateConversion will test the method CreateConversion()
func TestClient_CreateConversion(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	t.Run("create a conversion by goal id (success)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		conversion := newTestConversion()

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelConversion)

		err = mockResponseData(http.MethodPost, endpoint, http.StatusCreated, conversion)
		assert.NoError(t, err)

		var newConversion *Conversion
		newConversion, err = client.CreateConversion(
			WithGoalID(testGoalID),
			WithTncpwSession(testTncpwSession),
		)

		assert.NoError(t, err)
		assert.NotNil(t, newConversion)
		assert.Equal(t, testConversionID, newConversion.ID)
		assert.Equal(t, testGoalID, newConversion.GoalID)
	})

	t.Run("create a conversion by goal name (success)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		conversion := newTestConversion()

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelConversion)

		err = mockResponseData(http.MethodPost, endpoint, http.StatusCreated, conversion)
		assert.NoError(t, err)

		var newConversion *Conversion
		newConversion, err = client.CreateConversion(
			WithGoalName(testGoalName),
			WithTncpwSession(testTncpwSession),
		)

		assert.NoError(t, err)
		assert.NotNil(t, newConversion)
		assert.Equal(t, testConversionID, newConversion.ID)
		assert.Equal(t, testGoalName, newConversion.GoalName)
	})

	t.Run("create an e-commerce conversion (success)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		conversion := newTestConversion()

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelConversion)

		err = mockResponseData(http.MethodPost, endpoint, http.StatusCreated, conversion)
		assert.NoError(t, err)

		var newConversion *Conversion
		newConversion, err = client.CreateConversion(
			WithGoalID(testGoalID),
			WithTncpwSession(testTncpwSession),
			WithPurchaseAmount(120.00),
			WithCustomDimensions(`{"some_field":"some_value"`),
			WithDelay(30),
		)

		assert.NoError(t, err)
		assert.NotNil(t, newConversion)
		assert.Equal(t, testConversionID, newConversion.ID)
	})

	t.Run("missing goal id and session/user_id", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		conversion := newTestConversion()

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelConversion)

		err = mockResponseData(http.MethodPost, endpoint, http.StatusCreated, conversion)
		assert.NoError(t, err)

		var newConversion *Conversion
		newConversion, err = client.CreateConversion(
			WithDelay(1),
		)

		assert.Error(t, err)
		assert.Nil(t, newConversion)
	})

	t.Run("missing tncpw_session", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		conversion := newTestConversion()

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelConversion)

		err = mockResponseData(http.MethodPost, endpoint, http.StatusCreated, conversion)
		assert.NoError(t, err)

		var newConversion *Conversion
		newConversion, err = client.CreateConversion(
			WithGoalID(testGoalID),
		)

		assert.Error(t, err)
		assert.Nil(t, newConversion)
	})

	t.Run("missing goal", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		conversion := newTestConversion()

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelConversion)

		err = mockResponseData(http.MethodPost, endpoint, http.StatusCreated, conversion)
		assert.NoError(t, err)

		var newConversion *Conversion
		newConversion, err = client.CreateConversion(
			WithUserID(testUserID),
		)

		assert.Error(t, err)
		assert.Nil(t, newConversion)
	})

	t.Run("error from api (status code)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		conversion := newTestConversion()

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelConversion)

		err = mockResponseData(http.MethodPost, endpoint, http.StatusBadRequest, conversion)
		assert.NoError(t, err)

		var newConversion *Conversion
		newConversion, err = client.CreateConversion(
			WithGoalID(testGoalID),
			WithTncpwSession(testTncpwSession),
		)

		assert.Error(t, err)
		assert.Nil(t, newConversion)
	})

	t.Run("error from api (status code)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelConversion)

		apiError := &Error{
			Code:        400,
			Data:        "field_name",
			IPAddress:   "127.0.0.1",
			Message:     "some error message",
			Method:      http.MethodPut,
			RequestGUID: "7f3d97a8fd67ff57861904df6118dcc8",
			StatusCode:  http.StatusBadRequest,
			URL:         endpoint,
		}

		err = mockResponseData(http.MethodPost, endpoint, http.StatusBadRequest, apiError)
		assert.NoError(t, err)

		var newConversion *Conversion
		newConversion, err = client.CreateConversion(
			WithGoalID(testGoalID),
			WithTncpwSession(testTncpwSession),
		)

		assert.Error(t, err)
		assert.Nil(t, newConversion)
		assert.Equal(t, apiError.Message, err.Error())
	})
}

// ExampleClient_CreateConversion example using CreateConversion()
//
// See more examples in /examples/
func ExampleClient_CreateConversion() {

	// Load the client (using test client for example only)
	client, err := newTestClient()
	if err != nil {
		fmt.Printf("error loading client: %s", err.Error())
		return
	}

	// Mock response (for example only)
	responseConversion := newTestConversion()
	_ = mockResponseData(
		http.MethodPost,
		fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelConversion),
		http.StatusCreated,
		responseConversion,
	)

	// Create conversion (using mocking response)
	var newConversion *Conversion
	newConversion, err = client.CreateConversion(
		WithGoalID(testGoalID),
		WithTncpwSession(testTncpwSession),
	)

	if err != nil {
		fmt.Printf("error creating conversion: " + err.Error())
		return
	}
	fmt.Printf("conversion created: %d", newConversion.ID)
	// Output:conversion created: 99
}

// BenchmarkClient_CreateConversion benchmarks the method CreateConversion()
func BenchmarkClient_CreateConversion(b *testing.B) {
	client, _ := newTestClient()
	conversion := newTestConversion()
	_ = mockResponseData(
		http.MethodPost,
		fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelConversion),
		http.StatusCreated,
		conversion,
	)
	for i := 0; i < b.N; i++ {
		_, _ = client.CreateConversion(
			WithGoalID(testGoalID),
			WithTncpwSession(testTncpwSession),
		)
	}
}
