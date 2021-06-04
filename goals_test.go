package tonicpow

// newTestGoal will return a dummy example for tests
func newTestGoal() *Goal {
	return &Goal{
		CampaignID:     testCampaignID,
		Description:    "This is an example goal",
		ID:             testGoalID,
		MaxPerPromoter: 1,
		Name:           "example_goal",
		PayoutRate:     1,
		PayoutType:     "flat",
		Title:          "Example Goal",
	}
}
