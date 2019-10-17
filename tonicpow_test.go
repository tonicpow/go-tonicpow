package tonicpow

import (
	"fmt"
	"testing"
)

// TestNewClient test new client
func TestNewClient(t *testing.T) {
	client, err := NewClient("test123")
	if err != nil {
		t.Fatal(err)
	}

	if client.Parameters.AdvertiserSecretKey != "test123" {
		t.Fatal("expected value to be test123")
	}
}

// ExampleNewClient example using NewClient()
func ExampleNewClient() {
	client, _ := NewClient("test123")
	fmt.Println(client.Parameters.AdvertiserSecretKey)
	// Output:test123
}

// BenchmarkNewClient benchmarks the NewClient method
func BenchmarkNewClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewClient("test123")
	}
}

// TestClient_ConvertGoal tests the method ConvertGoal()
func TestClient_ConvertGoal(t *testing.T) {
	// Skip tis test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	privateGUID := "c2512ba9073341ae8c4bc8d153915cd8"
	client, err := NewClient(privateGUID)
	if err != nil {
		t.Fatal(err)
	}
	client.Parameters.Environment = LocalEnvironment

	if client.Parameters.AdvertiserSecretKey != privateGUID {
		t.Fatal("expected value to be " + privateGUID)
	}

	goalName := "signupgoal"
	sessionTxID := "f773c231ee9383125fe7932d6dbdb5447577c39cae8cc28210d19f6471294485"
	userID := "123"
	additionalData := "test data"

	var resp *ConversionResponse
	resp, err = client.ConvertGoal(goalName, sessionTxID, userID, additionalData)
	if err != nil {
		t.Fatal("error from ConvertGoal: " + err.Error())
	}

	t.Log("error:", resp.Message, "last_code", client.LastRequest.StatusCode, "payout: ", resp.PayoutTxID)
}
