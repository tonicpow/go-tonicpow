package tonicpow

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	// keyFromOS is the test api key for most tests (set in your env)
	keyFromOS = os.Getenv("TONICPOW_API_KEY")

	// testAPIKey is for examples
	testAPIKey = "3ez9d6z7a6549c3f5gf9g2cc8911achz"
)

func init() {
	if len(keyFromOS) > 0 {
		testAPIKey = keyFromOS
	}
}

// TestNewClient test new client
func TestNewClient(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	client, err := NewClient(testAPIKey, LocalEnvironment, nil)
	if err != nil {
		t.Fatal(err)
	}

	if client.Parameters.apiKey != testAPIKey {
		t.Fatalf("expected value to be %s", testAPIKey)
	}

	if client.Parameters.environment != LocalEnvironment {
		t.Fatalf("expected value to be %s, got %s", LocalEnvironment, client.Parameters.environment)
	}

	if client.Parameters.UserAgent != defaultUserAgent {
		t.Fatalf("expected value to be %s, got %s", defaultUserAgent, client.Parameters.UserAgent)
	}

	if client.LastRequest.StatusCode != 0 {
		t.Fatalf("expected value to be %d, got %d", 0, client.LastRequest.StatusCode)
	}

	if client.LastRequest.Method != "" {
		t.Fatalf("expected value to be %s, got %s", http.MethodPost, client.LastRequest.Method)
	}

	if len(client.LastRequest.URL) != 0 {
		t.Fatalf("expected value to be set, was empty/nil")
	}
}

// ExampleNewClient example using NewClient()
func ExampleNewClient() {
	client, _ := NewClient(testAPIKey, LocalEnvironment, nil)
	fmt.Println(client.Parameters.apiKey)
	// Output:3ez9d6z7a6549c3f5gf9g2cc8911achz
}

// BenchmarkNewClient benchmarks the NewClient method
func BenchmarkNewClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewClient(testAPIKey, LocalEnvironment, nil)
	}
}

// TestDefaultOptions tests setting ClientDefaultOptions()
func TestDefaultOptions(t *testing.T) {

	options := ClientDefaultOptions()

	if options.UserAgent != defaultUserAgent {
		t.Fatalf("expected value: %s got: %s", defaultUserAgent, options.UserAgent)
	}

	if options.BackOffExponentFactor != 2.0 {
		t.Fatalf("expected value: %f got: %f", 2.0, options.BackOffExponentFactor)
	}

	if options.BackOffInitialTimeout != 2*time.Millisecond {
		t.Fatalf("expected value: %v got: %v", 2*time.Millisecond, options.BackOffInitialTimeout)
	}

	if options.BackOffMaximumJitterInterval != 2*time.Millisecond {
		t.Fatalf("expected value: %v got: %v", 2*time.Millisecond, options.BackOffMaximumJitterInterval)
	}

	if options.BackOffMaxTimeout != 10*time.Millisecond {
		t.Fatalf("expected value: %v got: %v", 10*time.Millisecond, options.BackOffMaxTimeout)
	}

	if options.DialerKeepAlive != 20*time.Second {
		t.Fatalf("expected value: %v got: %v", 20*time.Second, options.DialerKeepAlive)
	}

	if options.DialerTimeout != 5*time.Second {
		t.Fatalf("expected value: %v got: %v", 5*time.Second, options.DialerTimeout)
	}

	if options.RequestRetryCount != 2 {
		t.Fatalf("expected value: %v got: %v", 2, options.RequestRetryCount)
	}

	if options.RequestTimeout != 10*time.Second {
		t.Fatalf("expected value: %v got: %v", 10*time.Second, options.RequestTimeout)
	}

	if options.TransportExpectContinueTimeout != 3*time.Second {
		t.Fatalf("expected value: %v got: %v", 3*time.Second, options.TransportExpectContinueTimeout)
	}

	if options.TransportIdleTimeout != 20*time.Second {
		t.Fatalf("expected value: %v got: %v", 20*time.Second, options.TransportIdleTimeout)
	}

	if options.TransportMaxIdleConnections != 10 {
		t.Fatalf("expected value: %v got: %v", 10, options.TransportMaxIdleConnections)
	}

	if options.TransportTLSHandshakeTimeout != 5*time.Second {
		t.Fatalf("expected value: %v got: %v", 5*time.Second, options.TransportTLSHandshakeTimeout)
	}
}

// TestClient_Request tests the method Request()
func TestClient_Request(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Start a new client
	_, err := NewClient(testAPIKey, LocalEnvironment, nil)
	if err != nil {
		t.Fatal(err)
	}
}

// TestClient_CreateUser tests the CreateUser() method
/*func TestClient_CreateUser(t *testing.T) {

	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Start a new client
	client, err := NewClient(testAPIKey, LocalEnvironment, nil)
	if err != nil {
		t.Fatal(err)
	}

	user := &User{
		Email:     fmt.Sprintf("Testing%d@TonicPow.com", rand.Intn(100000)),
		FirstName: "Austin",
	}
	if _, err = client.CreateUser(user); err != nil {
		t.Fatalf("%s", err.Error())
	}
}*/

// TestClient_UpdateUser tests the UpdateUser() method
/*func TestClient_UpdateUser(t *testing.T) {

	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Start a new client
	client, err := NewClient(testAPIKey, LocalEnvironment, nil)
	if err != nil {
		t.Fatal(err)
	}

	user := &User{
		Email:     fmt.Sprintf("Testing%d@TonicPow.com", rand.Intn(100000)),
		FirstName: "Austin",
	}
	if user, err = client.CreateUser(user); err != nil {
		t.Fatalf("%s", err.Error())
	}

	user.MiddleName = "Danger"
	if _, err = client.UpdateUser(user); err != nil {
		t.Fatalf("%s", err.Error())
	}
}*/

// TestClient_GetUser tests the GetUser() method
/*func TestClient_GetUser(t *testing.T) {

	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Start a new client
	client, err := NewClient(testAPIKey, LocalEnvironment, nil)
	if err != nil {
		t.Fatal(err)
	}

	user := &User{
		Email:     fmt.Sprintf("Testing%d@TonicPow.com", rand.Intn(100000)),
		FirstName: "Austin",
	}
	if user, err = client.CreateUser(user); err != nil {
		t.Fatalf("%s", err.Error())
	}

	if _, err = client.GetUser(user.ID, user.Email); err != nil {
		t.Fatalf("%s", err.Error())
	}
}*/

// todo: add more tests (covering all requests)

// todo: add examples for each Request that can be viewed in godocs
