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

	var err error

	// Example: Prolong a session
	if err = TonicPowAPI.ProlongSession(""); err != nil {
		log.Fatalf("ProlongSession: %s", err.Error())
	}

	log.Println("session created and prolonged...")

	// Example: Create a user
	user := &tonicpow.User{
		Email:     "Austin+Testing4@TonicPow.com",
		FirstName: "Austin",
	}
	if user, err = TonicPowAPI.CreateUser(user); err != nil {
		log.Fatalf("create user failed - api error: %s data: %s", TonicPowAPI.LastRequest.Error.Message, TonicPowAPI.LastRequest.Error.Data)
	}

	log.Printf("user %d created", user.ID)

	// Example: Update a user
	user.MiddleName = "Danger"
	if user, err = TonicPowAPI.UpdateUser(user, ""); err != nil {
		log.Fatalf("update user failed - api error: %s data: %s", TonicPowAPI.LastRequest.Error.Message, TonicPowAPI.LastRequest.Error.Data)
	}

	log.Printf("user %d updated - middle_name: %s", user.ID, user.MiddleName)

}
