package tonicpow

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
