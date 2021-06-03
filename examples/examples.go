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

	//
	// Get the API key (from env or your own config)
	//
	apiKey := os.Getenv("TONICPOW_API_KEY")
	if len(apiKey) == 0 {
		log.Fatalf("api key is required: %s", "TONICPOW_API_KEY")
	}

	// Set the environment
	environmentString := os.Getenv("TONICPOW_ENVIRONMENT")
	var environment tonicpow.APIEnvironment
	if environmentString == "staging" {
		environment = tonicpow.StagingEnvironment
	} else if environmentString == "live" {
		environment = tonicpow.LiveEnvironment
	} else {
		environment = tonicpow.LocalEnvironment
	}

	//
	// Load the api client (creates a new session)
	// You can also set the environment or client options
	//
	var err error
	TonicPowAPI, err = tonicpow.NewClient(apiKey, environment, nil)
	if err != nil {
		log.Fatalf("error in NewClient: %s", err.Error())
	}
}

func main() {
	log.Println("finish examples! user agent: ", TonicPowAPI.Parameters.UserAgent)
}
