package tonicpow

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	testAdvertiserID      uint64 = 23
	testAdvertiserName    string = "TonicPow Test"
	testAPIKey            string = "TestAPIKey12345678987654321"
	testAppID             uint64 = 10
	testCampaignID        uint64 = 23
	testCampaignTargetURL string = "https://tonicpow.com"
	testConversionID      uint64 = 99
	testGoalID            uint64 = 13
	testRateCurrency      string = "usd"
	testTncpwSession      string = "TestSessionKey12345678987654321"
	testGoalName          string = "example_goal"
	testUserID            uint64 = 43
)

// TestVersion will test the method Version()
func TestVersion(t *testing.T) {
	t.Parallel()

	t.Run("get version", func(t *testing.T) {
		ver := Version()
		assert.Equal(t, version, ver)
	})
}

// TestUserAgent will test the method UserAgent()
func TestUserAgent(t *testing.T) {
	t.Parallel()

	t.Run("get user agent", func(t *testing.T) {
		agent := UserAgent()
		assert.Equal(t, defaultUserAgent, agent)
	})
}

// mockResponseData is used for mocking the response
func mockResponseData(method, endpoint string, statusCode int, model interface{}) error {
	httpmock.Reset()
	if model != nil && model != "" {
		data, err := json.Marshal(model)
		if err != nil {
			return err
		}
		httpmock.RegisterResponder(method, endpoint, httpmock.NewStringResponder(statusCode, string(data)))
	} else {
		httpmock.RegisterResponder(method, endpoint, httpmock.NewStringResponder(statusCode, ""))
	}

	return nil
}

// mockResponseFeed is used for mocking the response
func mockResponseFeed(endpoint string, statusCode int, feedResults string) {
	httpmock.Reset()
	httpmock.RegisterResponder(http.MethodGet, endpoint, httpmock.NewStringResponder(statusCode, feedResults))
}
