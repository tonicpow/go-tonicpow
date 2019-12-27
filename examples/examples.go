package main

import (
	"log"
	"os"

	"github.com/tonicpow/go-tonicpow"
)

var (
	// TonicPowAPI is the client we will load on start-up
	TonicPowAPI *tonicpow.Client
)

// Load the TonicPow API Client once when the application loads
func init() {

	// Get the API key (from env or your own config)
	apiKey := os.Getenv("TONICPOW_API_KEY")
	if len(apiKey) == 0 {
		log.Fatalf("api key is required: %s", "TONICPOW_API_KEY")
	}

	// Load the api client (creates a new session)
	// You can also set the environment or client options
	var err error
	TonicPowAPI, err = tonicpow.NewClient(apiKey, tonicpow.LocalEnvironment, nil)
	if err != nil {
		log.Fatalf("error in NewClient: %s", err.Error())
	}
}

func main() {

	// Example for ending the api session for the application
	// This is not needed, sessions will automatically expire
	defer func() {
		_ = TonicPowAPI.EndSession("")
	}()

	// Example: Prolong a session
	if err := TonicPowAPI.ProlongSession(""); err != nil {
		log.Fatalf("ProlongSession: %s", err.Error())
	}

	log.Println("Session created and prolonged! Ending...")
}
