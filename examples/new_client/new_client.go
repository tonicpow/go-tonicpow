package main

import (
	"log"
	"os"

	"github.com/tonicpow/go-tonicpow"
)

func main() {

	//
	// Load the api client
	// You can also set the environment or client options
	//
	client, err := tonicpow.NewClient(
		tonicpow.WithAPIKey(os.Getenv("TONICPOW_API_KEY")),
		tonicpow.WithEnvironmentString(os.Getenv("TONICPOW_ENVIRONMENT")),
	)
	if err != nil {
		log.Fatalf("error in NewClient: %s", err.Error())
	}

	log.Println(
		"client: ", client.GetUserAgent(),
		"environment: ", client.GetEnvironment(),
	)
}
