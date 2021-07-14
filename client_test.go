package tonicpow

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// newTestClient will return a client for testing purposes
func newTestClient() (*Client, error) {
	// Create a Resty Client
	client := resty.New()

	// Get the underlying HTTP Client and set it to Mock
	httpmock.ActivateNonDefault(client.GetClient())

	// Add custom headers in request
	headers := make(map[string][]string)
	headers["custom_header_1"] = append(headers["custom_header_1"], "value_1")

	// Create a new client
	newClient, err := NewClient(
		WithRequestTracing(),
		WithAPIKey(testAPIKey),
		WithEnvironment(EnvironmentDevelopment),
		WithCustomHeaders(headers),
	)
	if err != nil {
		return nil, err
	}
	newClient.WithCustomHTTPClient(client)

	// Return the mocking client
	return newClient, nil
}

// TestNewClient will test the method NewClient()
func TestNewClient(t *testing.T) {
	t.Parallel()

	t.Run("default client", func(t *testing.T) {
		client, err := NewClient(WithAPIKey(testAPIKey))
		assert.NoError(t, err)
		assert.NotNil(t, client)
		assert.Equal(t, defaultHTTPTimeout, client.options.httpTimeout)
		assert.Equal(t, defaultRetryCount, client.options.retryCount)
		assert.Equal(t, defaultUserAgent, client.options.userAgent)
		assert.Equal(t, false, client.options.requestTracing)
		assert.Equal(t, EnvironmentLive.apiURL, client.options.apiURL)
		assert.Equal(t, EnvironmentLive.name, client.options.environment)
	})

	t.Run("missing api key", func(t *testing.T) {
		client, err := NewClient(WithAPIKey(""))
		assert.Error(t, err)
		assert.Nil(t, client)
	})

	t.Run("custom http client", func(t *testing.T) {
		customHTTPClient := resty.New()
		customHTTPClient.SetTimeout(defaultHTTPTimeout)
		client, err := NewClient(WithAPIKey(testAPIKey))
		assert.NoError(t, err)
		assert.NotNil(t, client)
		client.WithCustomHTTPClient(customHTTPClient)
	})

	t.Run("custom http timeout", func(t *testing.T) {
		client, err := NewClient(WithAPIKey(testAPIKey), WithHTTPTimeout(10*time.Second))
		assert.NoError(t, err)
		assert.NotNil(t, client)
		assert.Equal(t, 10*time.Second, client.options.httpTimeout)
	})

	t.Run("custom retry count", func(t *testing.T) {
		client, err := NewClient(WithAPIKey(testAPIKey), WithRetryCount(3))
		assert.NoError(t, err)
		assert.NotNil(t, client)
		assert.Equal(t, 3, client.options.retryCount)
	})

	t.Run("custom headers", func(t *testing.T) {
		headers := make(map[string][]string)
		headers["custom_header_1"] = append(headers["custom_header_1"], "value_1")
		headers["custom_header_2"] = append(headers["custom_header_2"], "value_1")
		headers["custom_header_2"] = append(headers["custom_header_2"], "value_2")
		client, err := NewClient(WithAPIKey(testAPIKey), WithCustomHeaders(headers))
		assert.NoError(t, err)
		assert.NotNil(t, client)
		assert.Equal(t, 2, len(client.options.customHeaders))
		assert.Equal(t, []string{"value_1"}, client.options.customHeaders["custom_header_1"])
	})

	t.Run("custom options", func(t *testing.T) {
		client, err := NewClient(WithAPIKey(testAPIKey), WithUserAgent("custom user agent"))
		assert.NotNil(t, client)
		assert.NoError(t, err)
		assert.Equal(t, "custom user agent", client.GetUserAgent())
	})

	t.Run("custom environment (staging)", func(t *testing.T) {
		client, err := NewClient(WithAPIKey(testAPIKey), WithEnvironment(EnvironmentStaging))
		assert.NotNil(t, client)
		assert.NoError(t, err)
		assert.Equal(t, client.options.apiURL, EnvironmentStaging.apiURL)
		assert.Equal(t, client.options.environment, environmentStagingName)

		client, err = NewClient(WithAPIKey(testAPIKey), WithEnvironmentString(environmentStagingName))
		assert.NotNil(t, client)
		assert.NoError(t, err)
		assert.Equal(t, client.options.apiURL, EnvironmentStaging.apiURL)
		assert.Equal(t, client.options.environment, environmentStagingName)
	})

	t.Run("custom environment (development)", func(t *testing.T) {
		client, err := NewClient(WithAPIKey(testAPIKey), WithEnvironment(EnvironmentDevelopment))
		assert.NotNil(t, client)
		assert.NoError(t, err)
		assert.Equal(t, client.options.apiURL, EnvironmentDevelopment.apiURL)
		assert.Equal(t, client.options.environment, environmentDevelopmentName)

		client, err = NewClient(WithAPIKey(testAPIKey), WithEnvironmentString(environmentDevelopmentName))
		assert.NotNil(t, client)
		assert.NoError(t, err)
		assert.Equal(t, client.options.apiURL, EnvironmentDevelopment.apiURL)
		assert.Equal(t, client.options.environment, environmentDevelopmentName)
	})

	t.Run("default no environment", func(t *testing.T) {
		client, err := NewClient(WithAPIKey(testAPIKey), WithEnvironmentString(""))
		assert.NotNil(t, client)
		assert.NoError(t, err)
		assert.Equal(t, client.options.apiURL, EnvironmentLive.apiURL)
		assert.Equal(t, client.options.environment, environmentLiveName)
	})
}

// TestClient_GetUserAgent will test the method GetUserAgent()
func TestClient_GetUserAgent(t *testing.T) {
	t.Parallel()

	t.Run("get user agent", func(t *testing.T) {
		client, err := NewClient(WithAPIKey(testAPIKey))
		assert.NoError(t, err)
		assert.NotNil(t, client)
		userAgent := client.GetUserAgent()
		assert.Equal(t, defaultUserAgent, userAgent)
	})
}

// ExampleVersion example using Version()
//
// See more examples in /examples/
func ExampleVersion() {
	fmt.Printf("version: %s", Version())
	// Output:version: v0.6.5
}

// ExampleUserAgent example using UserAgent()
//
// See more examples in /examples/
func ExampleUserAgent() {
	fmt.Printf("user agent: %s", UserAgent())
	// Output:user agent: go-tonicpow: v0.6.5
}

// TestClient_GetEnvironment will test the method GetEnvironment()
func TestClient_GetEnvironment(t *testing.T) {
	t.Parallel()

	t.Run("get client environment", func(t *testing.T) {
		client, err := NewClient(WithAPIKey(testAPIKey))
		assert.NoError(t, err)
		assert.NotNil(t, client)
		env := client.GetEnvironment()
		assert.Equal(t, environmentLiveName, env)
	})
}

// ExampleNewClient example using NewClient()
//
// See more examples in /examples/
func ExampleNewClient() {
	client, err := NewClient(WithAPIKey(testAPIKey))
	if err != nil {
		fmt.Printf("error loading client: %s", err.Error())
		return
	}
	fmt.Printf("loaded client: %s", client.options.userAgent)
	// Output:loaded client: go-tonicpow: v0.6.5
}

// BenchmarkNewClient benchmarks the method NewClient()
func BenchmarkNewClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewClient(WithAPIKey(testAPIKey))
	}
}

// TestDefaultClientOptions will test the method defaultClientOptions()
func TestDefaultClientOptions(t *testing.T) {
	t.Parallel()

	options := defaultClientOptions()
	assert.NotNil(t, options)

	assert.Equal(t, defaultHTTPTimeout, options.httpTimeout)
	assert.Equal(t, defaultRetryCount, options.retryCount)
	assert.Equal(t, defaultUserAgent, options.userAgent)
	assert.Equal(t, EnvironmentLive.apiURL, options.apiURL)
	assert.Equal(t, EnvironmentLive.name, options.environment)
	assert.Equal(t, false, options.requestTracing)
}

// BenchmarkDefaultClientOptions benchmarks the method defaultClientOptions()
func BenchmarkDefaultClientOptions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = defaultClientOptions()
	}
}
