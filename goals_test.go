package tonicpow

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// newTestGoal will return a dummy example for tests
func newTestGoal() *Goal {
	return &Goal{
		CampaignID:     testCampaignID,
		Description:    "This is an example goal",
		ID:             testGoalID,
		MaxPerPromoter: 1,
		Name:           "example_goal",
		PayoutRate:     0.01,
		PayoutType:     "flat",
		Title:          "Example Goal",
	}
}

// TestClient_CreateGoal will test the method CreateGoal()
func TestClient_CreateGoal(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	t.Run("create a goal (success)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		goal := newTestGoal()

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelGoal)

		err = mockResponseData(http.MethodPost, endpoint, http.StatusCreated, goal)
		assert.NoError(t, err)

		err = client.CreateGoal(goal)
		assert.NoError(t, err)
		assert.NotNil(t, goal)
		assert.Equal(t, testGoalID, goal.ID)
	})

	t.Run("missing campaign id", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		goal := newTestGoal()
		goal.CampaignID = 0

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelGoal)

		err = mockResponseData(http.MethodPost, endpoint, http.StatusCreated, goal)
		assert.NoError(t, err)

		err = client.CreateGoal(goal)
		assert.Error(t, err)
	})

	t.Run("missing goal name", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		goal := newTestGoal()
		goal.Name = ""

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelGoal)

		err = mockResponseData(http.MethodPost, endpoint, http.StatusCreated, goal)
		assert.NoError(t, err)

		err = client.CreateGoal(goal)
		assert.Error(t, err)
	})

	t.Run("error from api (status code)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		goal := newTestGoal()

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelGoal)

		err = mockResponseData(http.MethodPost, endpoint, http.StatusBadRequest, goal)
		assert.NoError(t, err)

		err = client.CreateGoal(goal)
		assert.Error(t, err)
	})

	t.Run("error from api (api error)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		goal := newTestGoal()

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelGoal)

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

		err = client.CreateGoal(goal)
		assert.Error(t, err)
		assert.Equal(t, apiError.Message, err.Error())
	})
}

// ExampleClient_CreateGoal example using CreateGoal()
//
// See more examples in /examples/
func ExampleClient_CreateGoal() {

	// Load the client (using test client for example only)
	client, err := newTestClient()
	if err != nil {
		fmt.Printf("error loading client: %s", err.Error())
		return
	}

	// Mock response (for example only)
	responseGoal := newTestGoal()
	_ = mockResponseData(
		http.MethodPost,
		fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelGoal),
		http.StatusCreated,
		responseGoal,
	)

	// Create goal (using mocking response)
	if err = client.CreateGoal(responseGoal); err != nil {
		fmt.Printf("error creating goal: " + err.Error())
		return
	}
	fmt.Printf("created goal: %s", responseGoal.Name)
	// Output:created goal: example_goal
}

// BenchmarkClient_CreateGoal benchmarks the method CreateGoal()
func BenchmarkClient_CreateGoal(b *testing.B) {
	client, _ := newTestClient()
	goal := newTestGoal()
	_ = mockResponseData(
		http.MethodPost,
		fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelGoal),
		http.StatusCreated,
		goal,
	)
	for i := 0; i < b.N; i++ {
		_ = client.CreateGoal(goal)
	}
}

// TestClient_GetGoal will test the method GetGoal()
func TestClient_GetGoal(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	t.Run("get a goal (success)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		goal := newTestGoal()

		endpoint := fmt.Sprintf(
			"%s/%s/details/%d", EnvironmentDevelopment.apiURL,
			modelGoal, goal.ID,
		)

		err = mockResponseData(http.MethodGet, endpoint, http.StatusOK, goal)
		assert.NoError(t, err)

		var newGoal *Goal
		newGoal, err = client.GetGoal(goal.ID)
		assert.NoError(t, err)
		assert.NotNil(t, newGoal)
		assert.Equal(t, testGoalID, goal.ID)
	})

	t.Run("missing goal id", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		goal := newTestGoal()
		goal.ID = 0

		endpoint := fmt.Sprintf(
			"%s/%s/details/%d", EnvironmentDevelopment.apiURL,
			modelGoal, goal.ID,
		)

		err = mockResponseData(http.MethodGet, endpoint, http.StatusOK, goal)
		assert.NoError(t, err)

		var newGoal *Goal
		newGoal, err = client.GetGoal(goal.ID)
		assert.Error(t, err)
		assert.Nil(t, newGoal)
	})

	t.Run("error from api (status code)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		goal := newTestGoal()

		endpoint := fmt.Sprintf(
			"%s/%s/details/%d", EnvironmentDevelopment.apiURL,
			modelGoal, goal.ID,
		)
		err = mockResponseData(http.MethodGet, endpoint, http.StatusBadRequest, goal)
		assert.NoError(t, err)

		var newGoal *Goal
		newGoal, err = client.GetGoal(goal.ID)
		assert.Error(t, err)
		assert.Nil(t, newGoal)
	})

	t.Run("error from api (api error)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		goal := newTestGoal()

		endpoint := fmt.Sprintf(
			"%s/%s/details/%d", EnvironmentDevelopment.apiURL,
			modelGoal, goal.ID,
		)

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

		err = mockResponseData(http.MethodGet, endpoint, http.StatusBadRequest, apiError)
		assert.NoError(t, err)

		var newGoal *Goal
		newGoal, err = client.GetGoal(goal.ID)
		assert.Error(t, err)
		assert.Nil(t, newGoal)
		assert.Equal(t, apiError.Message, err.Error())
	})
}

// ExampleClient_GetGoal example using GetGoal()
//
// See more examples in /examples/
func ExampleClient_GetGoal() {

	// Load the client (using test client for example only)
	client, err := newTestClient()
	if err != nil {
		fmt.Printf("error loading client: %s", err.Error())
		return
	}

	// Mock response (for example only)
	responseGoal := newTestGoal()
	_ = mockResponseData(
		http.MethodGet,
		fmt.Sprintf(
			"%s/%s/details/%d", EnvironmentDevelopment.apiURL,
			modelGoal, responseGoal.ID,
		),
		http.StatusOK,
		responseGoal,
	)

	// Get goal (using mocking response)
	if responseGoal, err = client.GetGoal(responseGoal.ID); err != nil {
		fmt.Printf("error getting goal: " + err.Error())
		return
	}
	fmt.Printf("goal: %s", responseGoal.Name)
	// Output:goal: example_goal
}

// BenchmarkClient_GetGoal benchmarks the method GetGoal()
func BenchmarkClient_GetGoal(b *testing.B) {
	client, _ := newTestClient()
	goal := newTestGoal()
	_ = mockResponseData(
		http.MethodGet,
		fmt.Sprintf(
			"%s/%s/details/%d", EnvironmentDevelopment.apiURL,
			modelGoal, goal.ID,
		),
		http.StatusOK,
		goal,
	)
	for i := 0; i < b.N; i++ {
		_, _ = client.GetGoal(goal.ID)
	}
}
