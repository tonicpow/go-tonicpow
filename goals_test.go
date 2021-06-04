package tonicpow

// newTestGoal will return a dummy example for tests
func newTestGoal() *Goal {
	return &Goal{
		CampaignID:     23,
		Description:    "This is an example goal",
		ID:             1,
		MaxPerPromoter: 1,
		Name:           "example_goal",
		PayoutRate:     1,
		PayoutType:     "flat",
		Title:          "Example Goal",
	}
}
