package tonicpow

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// newTestCampaign will return a dummy example for tests
func newTestCampaign() *Campaign {
	return &Campaign{
		Goals:               []*Goal{newTestGoal()},
		Images:              []*CampaignImage{newTestCampaignImages()},
		CreatedAt:           "2021-01-01 00:00:01",
		Currency:            "usd",
		Description:         "This is a test campaign",
		FundingAddress:      "124oW4xLDfay1BXmubUG9r64bGCCxnuf4g",
		ImageURL:            "https://res.cloudinary.com/tonicpow/image/upload/c_crop,x_0,y_0,w_2048,h_1024/w_600,h_300,c_limit,g_center/v1611266301/glfwmr0yhyjydeyfhyih.jpg",
		PublicGUID:          "b02e13a7001546b1b7efb9df40ab75e5",
		Slug:                "tonicpow",
		TargetURL:           "https://tonicpow.com",
		Title:               "TonicPow",
		AdvertiserProfileID: testAdvertiserID,
		AdvertiserProfile:   newTestAdvertiserProfile(),
		Balance:             13.37,
		PayPerClickRate:     1,
		BalanceSatoshis:     11333377,
		ID:                  testCampaignID,
		LinksCreated:        1,
		PaidClicks:          1,
		PaidConversions:     1,
		Requirements:        newTestCampaignRequirements(),
		BotProtection:       true,
		ContributeEnabled:   true,
		MatchDomain:         true,
	}
}

// newTestCampaignImages will return a dummy example for tests
func newTestCampaignImages() *CampaignImage {
	return &CampaignImage{
		Height:   300,
		MimeType: "image/jpeg",
		URL:      "https://res.cloudinary.com/tonicpow/image/upload/c_crop,x_0,y_0,w_2048,h_1024/w_600,h_300,c_limit,g_center/v1611266301/glfwmr0yhyjydeyfhyih.jpg",
		Width:    600,
	}
}

// newTestCampaignRequirements will return a dummy example for tests
func newTestCampaignRequirements() *CampaignRequirements {
	return &CampaignRequirements{
		HandCash:    true,
		MoneyButton: true,
	}
}

// newTestCampaignResults will return a dummy example for tests
func newTestCampaignResults(currentPage, resultsPerPage int) *CampaignResults {
	return &CampaignResults{
		Campaigns:      []*Campaign{newTestCampaign()},
		CurrentPage:    currentPage,
		Results:        1,
		ResultsPerPage: resultsPerPage,
	}
}

// TestClient_CreateCampaign will test the method CreateCampaign()
func TestClient_CreateCampaign(t *testing.T) {
	// t.Parallel() (Cannot run in parallel - issues with overriding the mock client)

	t.Run("create a campaign (success)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		campaign := newTestCampaign()

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelCampaign)

		err = mockResponseData(http.MethodPost, endpoint, http.StatusCreated, campaign)
		assert.NoError(t, err)

		err = client.CreateCampaign(campaign)
		assert.NoError(t, err)
		assert.NotNil(t, campaign)
		assert.Equal(t, testCampaignID, campaign.ID)
	})

	t.Run("missing advertiser profile id", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		campaign := newTestCampaign()
		campaign.AdvertiserProfileID = 0

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelCampaign)

		err = mockResponseData(http.MethodPost, endpoint, http.StatusCreated, campaign)
		assert.NoError(t, err)

		err = client.CreateCampaign(campaign)
		assert.Error(t, err)
	})

	t.Run("missing campaign title", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		campaign := newTestCampaign()
		campaign.Title = ""

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelCampaign)

		err = mockResponseData(http.MethodPost, endpoint, http.StatusCreated, campaign)
		assert.NoError(t, err)

		err = client.CreateCampaign(campaign)
		assert.Error(t, err)
	})

	t.Run("missing campaign description", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		campaign := newTestCampaign()
		campaign.Description = ""

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelCampaign)

		err = mockResponseData(http.MethodPost, endpoint, http.StatusCreated, campaign)
		assert.NoError(t, err)

		err = client.CreateCampaign(campaign)
		assert.Error(t, err)
	})

	t.Run("missing campaign target url", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		campaign := newTestCampaign()
		campaign.TargetURL = ""

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelCampaign)

		err = mockResponseData(http.MethodPost, endpoint, http.StatusCreated, campaign)
		assert.NoError(t, err)

		err = client.CreateCampaign(campaign)
		assert.Error(t, err)
	})

	t.Run("error from api (status code)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		campaign := newTestCampaign()

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelCampaign)

		err = mockResponseData(http.MethodPost, endpoint, http.StatusBadRequest, campaign)
		assert.NoError(t, err)

		err = client.CreateCampaign(campaign)
		assert.Error(t, err)
	})

	t.Run("error from api (api error)", func(t *testing.T) {
		client, err := newTestClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		campaign := newTestCampaign()

		endpoint := fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelCampaign)

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

		err = client.CreateCampaign(campaign)
		assert.Error(t, err)
		assert.Equal(t, apiError.Message, err.Error())
	})
}

// ExampleClient_CreateCampaign example using CreateCampaign()
//
// See more examples in /examples/
func ExampleClient_CreateCampaign() {

	// Load the client (using test client for example only)
	client, err := newTestClient()
	if err != nil {
		fmt.Printf("error loading client: %s", err.Error())
		return
	}

	// For mocking
	responseCampaign := newTestCampaign()

	// Mock response (for example only)
	_ = mockResponseData(
		http.MethodPost,
		fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelCampaign),
		http.StatusCreated,
		responseCampaign,
	)

	// Create campaign (using mocking response)
	if err = client.CreateCampaign(responseCampaign); err != nil {
		fmt.Printf("error creating campaign: " + err.Error())
		return
	}
	fmt.Printf("created campaign: %s", responseCampaign.Title)
	// Output:created campaign: TonicPow
}

// BenchmarkClient_CreateCampaign benchmarks the method CreateCampaign()
func BenchmarkClient_CreateCampaign(b *testing.B) {
	client, _ := newTestClient()
	campaign := newTestCampaign()
	_ = mockResponseData(
		http.MethodPost,
		fmt.Sprintf("%s/%s", EnvironmentDevelopment.apiURL, modelCampaign),
		http.StatusCreated,
		campaign,
	)
	for i := 0; i < b.N; i++ {
		_ = client.CreateCampaign(campaign)
	}
}
