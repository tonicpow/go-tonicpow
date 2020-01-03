package main

import (
	"fmt"
	"log"
	"math/rand"
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

	//
	// Load the api client (creates a new session)
	// You can also set the environment or client options
	//
	var err error
	TonicPowAPI, err = tonicpow.NewClient(apiKey, tonicpow.LocalEnvironment, nil)
	if err != nil {
		log.Fatalf("error in NewClient: %s", err.Error())
	}
}

func main() {
	//
	// Example for ending the api session for the application
	// This is not needed, sessions will automatically expire
	//
	defer func() {
		_ = TonicPowAPI.EndSession("")
	}()

	// Example vars
	var err error
	var userSessionToken string
	testPassword := "ExamplePassForNow0!"

	//
	// Example: Prolong a session
	//
	if err = TonicPowAPI.ProlongSession(""); err != nil {
		log.Fatalf("ProlongSession: %s", err.Error())
	} else {
		log.Println("session created and prolonged...")
	}

	//
	// Example: Create a user
	//
	user := &tonicpow.User{
		Email:    fmt.Sprintf("Tes_ti-ng+%d@TonicPow.com", rand.Intn(100000)),
		Password: testPassword,
	}
	if user, err = TonicPowAPI.CreateUser(user); err != nil {
		log.Fatalf("create user failed - api error: %s data: %s", TonicPowAPI.LastRequest.Error.Message, TonicPowAPI.LastRequest.Error.Data)
	} else {
		log.Printf("user %d created", user.ID)
	}

	//
	// Example: Get a user (id)
	//
	if user, err = TonicPowAPI.GetUser(user.ID, ""); err != nil {
		log.Fatalf("get user failed - api error: %s", TonicPowAPI.LastRequest.Error.Message)
	} else {
		log.Printf("got user by id %d", user.ID)
	}

	//
	// Example: Get a user (email)
	//
	if user, err = TonicPowAPI.GetUser(0, user.Email); err != nil {
		log.Fatalf("get user failed - api error: %s", TonicPowAPI.LastRequest.Error.Message)
	} else {
		log.Printf("got user by email %s", user.Email)
	}

	//
	// Example: Update a user
	//
	user.FirstName = "Austin"
	if user, err = TonicPowAPI.UpdateUser(user, ""); err != nil {
		log.Fatalf("update user failed - api error: %s data: %s", TonicPowAPI.LastRequest.Error.Message, TonicPowAPI.LastRequest.Error.Data)
	} else {
		log.Printf("user %d updated - first_name: %s", user.ID, user.FirstName)
	}

	//
	// Example: Get new updated balance for user
	//
	if user, err = TonicPowAPI.GetUserBalance(user.ID); err != nil {
		log.Fatalf("get user failed - api error: %s", TonicPowAPI.LastRequest.Error.Message)
	} else {
		log.Printf("user balance: %d", user.Balance)
	}

	//
	// Example: Forgot password
	//
	if err = TonicPowAPI.ForgotPassword(user.Email); err != nil {
		log.Fatalf("forgot password failed - api error: %s", TonicPowAPI.LastRequest.Error.Message)
	} else {
		log.Printf("sent forgot password: %s", user.Email)
	}

	//
	// Example: Create an advertiser
	//
	advertiser := &tonicpow.AdvertiserProfile{
		UserID:      user.ID,
		Name:        "Acme Advertising",
		HomepageURL: "https://tonicpow.com",
		IconURL:     "https://tonicpow.com/images/logos/apple-touch-icon.png",
	}
	if advertiser, err = TonicPowAPI.CreateAdvertiserProfile(advertiser, ""); err != nil {
		log.Fatalf("create advertiser failed - api error: %s data: %s", TonicPowAPI.LastRequest.Error.Message, TonicPowAPI.LastRequest.Error.Data)
	} else {
		log.Printf("advertiser profile %s id %d created", advertiser.Name, advertiser.ID)
	}

	//
	// Example: Get an advertiser profile
	//
	if advertiser, err = TonicPowAPI.GetAdvertiserProfile(advertiser.ID, ""); err != nil {
		log.Fatalf("get advertiser profile failed - api error: %s", TonicPowAPI.LastRequest.Error.Message)
	} else {
		log.Printf("got advertiser profile by id %d", advertiser.ID)
	}

	//
	// Example: Login for a user
	//
	user.Password = testPassword
	userSessionToken, err = TonicPowAPI.LoginUser(user)
	if err != nil {
		log.Fatalf("user login failed - api error: %s data: %s", TonicPowAPI.LastRequest.Error.Message, TonicPowAPI.LastRequest.Error.Data)
	} else {
		log.Printf("user login: %s token: %s", user.Email, userSessionToken)
	}

	//
	// Example: Logout (just for our example)
	//
	defer func(token string) {
		_ = TonicPowAPI.LogoutUser(token)
		log.Println("user logout complete")
	}(userSessionToken)

	//
	// Example: Current user details
	//
	user, err = TonicPowAPI.CurrentUser(userSessionToken)
	if err != nil {
		log.Fatalf("current user failed - api error: %s data: %s", TonicPowAPI.LastRequest.Error.Message, TonicPowAPI.LastRequest.Error.Data)
	} else {
		log.Printf("current user: %s", user.Email)
	}

	//
	// Example: Create an advertiser as a user
	//
	advertiser = &tonicpow.AdvertiserProfile{
		UserID:      user.ID,
		Name:        "Acme User Advertising",
		HomepageURL: "https://tonicpow.com",
		IconURL:     "https://tonicpow.com/images/logos/apple-touch-icon.png",
	}
	if advertiser, err = TonicPowAPI.CreateAdvertiserProfile(advertiser, userSessionToken); err != nil {
		log.Fatalf("create advertiser failed - api error: %s data: %s", TonicPowAPI.LastRequest.Error.Message, TonicPowAPI.LastRequest.Error.Data)
	} else {
		log.Printf("advertiser profile %s id %d created", advertiser.Name, advertiser.ID)
	}

	//
	// Example: Get Advertiser Profile
	//
	if advertiser, err = TonicPowAPI.GetAdvertiserProfile(advertiser.ID, userSessionToken); err != nil {
		log.Fatalf("get advertiser profile failed - api error: %s", TonicPowAPI.LastRequest.Error.Message)
	} else {
		log.Printf("got advertiser profile by id %d", advertiser.ID)
	}

	//
	// Example: Update advertising profile
	//
	advertiser.Name = "Acme New User Advertising"
	if advertiser, err = TonicPowAPI.UpdateAdvertiserProfile(advertiser, userSessionToken); err != nil {
		log.Fatalf("update advertiser failed - api error: %s data: %s", TonicPowAPI.LastRequest.Error.Message, TonicPowAPI.LastRequest.Error.Data)
	} else {
		log.Printf("advertiser profile %s id %d updated", advertiser.Name, advertiser.ID)
	}

	//
	// Example: Create a campaign
	//
	campaign := &tonicpow.Campaign{
		AdvertiserProfileID: advertiser.ID,
		Currency:            "USD",
		Description:         "Earn BSV for sharing things you like.",
		ImageURL:            "https://i.imgur.com/TbRFiaR.png",
		PayPerClickRate:     0.01,
		TargetURL:           "https://offers.tonicpow.com",
		Title:               "TonicPow Offers",
	}
	if campaign, err = TonicPowAPI.CreateCampaign(campaign, userSessionToken); err != nil {
		log.Fatalf("create campaign failed - api error: %s data: %s", TonicPowAPI.LastRequest.Error.Message, TonicPowAPI.LastRequest.Error.Data)
	} else {
		log.Printf("campaign %s id %d created", campaign.Title, campaign.ID)
	}

	//
	// Example: Get Campaign
	//
	if campaign, err = TonicPowAPI.GetCampaign(campaign.ID, userSessionToken); err != nil {
		log.Fatalf("get campaign failed - api error: %s", TonicPowAPI.LastRequest.Error.Message)
	} else {
		log.Printf("got campaign by id %d", campaign.ID)
	}

	//
	// Example: Create a Goal
	//
	goal := &tonicpow.Goal{
		CampaignID:  campaign.ID,
		Description: "Bring leads and get paid!",
		Name:        "new-lead-landing-page",
		PayoutRate:  0.50,
		PayoutType:  "flat",
		Title:       "Landing Page Leads",
	}
	if goal, err = TonicPowAPI.CreateGoal(goal, userSessionToken); err != nil {
		log.Fatalf("create goal failed - api error: %s data: %s", TonicPowAPI.LastRequest.Error.Message, TonicPowAPI.LastRequest.Error.Data)
	} else {
		log.Printf("goal %s id %d created", goal.Title, goal.ID)
	}

	//
	// Example: Get Goal
	//
	if goal, err = TonicPowAPI.GetGoal(goal.ID, userSessionToken); err != nil {
		log.Fatalf("get goal failed - api error: %s", TonicPowAPI.LastRequest.Error.Message)
	} else {
		log.Printf("got goal by id %d", goal.ID)
	}

	//
	// Example: Create a Link (making a custom short code using the user's name)
	//
	link := &tonicpow.Link{
		CampaignID:      campaign.ID,
		UserID:          user.ID,
		CustomShortCode: fmt.Sprintf("%s%d", user.FirstName, rand.Intn(100000)),
	}
	if link, err = TonicPowAPI.CreateLink(link, userSessionToken); err != nil {
		log.Fatalf("create link failed - api error: %s data: %s", TonicPowAPI.LastRequest.Error.Message, TonicPowAPI.LastRequest.Error.Data)
	} else {
		log.Printf("link %s id %d created", link.ShortCode, link.ID)
	}

	//
	// Example: Get Link
	//
	if link, err = TonicPowAPI.GetLink(link.ID, userSessionToken); err != nil {
		log.Fatalf("get link failed - api error: %s", TonicPowAPI.LastRequest.Error.Message)
	} else {
		log.Printf("got link by id %d", goal.ID)
	}

	//
	// Example: List active campaigns
	//
	var campaigns []*tonicpow.Campaign
	if campaigns, err = TonicPowAPI.ListCampaigns(""); err != nil {
		log.Fatalf("list campaign failed - api error: %s", TonicPowAPI.LastRequest.Error.Message)
	} else {
		log.Printf("campaigns found: %d", len(campaigns))
	}

	//
	// Example: Activate User
	//
	//if err = TonicPowAPI.ActivateUser(user.ID); err != nil {
	//	log.Fatalf("activate user failed - api error: %s", TonicPowAPI.LastRequest.Error.Message)
	//}
}
