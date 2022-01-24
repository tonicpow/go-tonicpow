package tonicpow

import "github.com/go-resty/resty/v2"

// AdvertiserService is the advertiser requests
type AdvertiserService interface {
	GetAdvertiserProfile(profileID uint64) (profile *AdvertiserProfile, response *StandardResponse, err error)
	ListAppsByAdvertiserProfile(profileID uint64, page, resultsPerPage int, sortBy, sortOrder string) (apps *AppResults, response *StandardResponse, err error)
	ListCampaignsByAdvertiserProfile(profileID uint64, page, resultsPerPage int, sortBy, sortOrder string) (campaigns *CampaignResults, response *StandardResponse, err error)
	UpdateAdvertiserProfile(profile *AdvertiserProfile) (*StandardResponse, error)
}

// CampaignService is the campaign requests
type CampaignService interface {
	CampaignsFeed(feedType FeedType) (feed string, response *StandardResponse, err error)
	CreateCampaign(campaign *Campaign) (*StandardResponse, error)
	GetCampaign(campaignID uint64) (campaign *Campaign, response *StandardResponse, err error)
	GetCampaignBySlug(slug string) (campaign *Campaign, response *StandardResponse, err error)
	ListCampaigns(page, resultsPerPage int, sortBy, sortOrder, searchQuery string, minimumBalance uint64, includeExpired bool) (results *CampaignResults, response *StandardResponse, err error)
	ListCampaignsByURL(targetURL string, page, resultsPerPage int, sortBy, sortOrder string) (results *CampaignResults, response *StandardResponse, err error)
	UpdateCampaign(campaign *Campaign) (response *StandardResponse, err error)
}

// ConversionService is the conversion requests
type ConversionService interface {
	CancelConversion(conversionID uint64, cancelReason string) (conversion *Conversion, response *StandardResponse, err error)
	CreateConversion(opts ...ConversionOps) (conversion *Conversion, response *StandardResponse, err error)
	GetConversion(conversionID uint64) (conversion *Conversion, response *StandardResponse, err error)
}

// GoalService is the goal requests
type GoalService interface {
	CreateGoal(goal *Goal) (*StandardResponse, error)
	DeleteGoal(goalID uint64) (bool, *StandardResponse, error)
	GetGoal(goalID uint64) (goal *Goal, response *StandardResponse, err error)
	UpdateGoal(goal *Goal) (*StandardResponse, error)
}

// RateService is the rate requests
type RateService interface {
	GetCurrentRate(currency string, customAmount float64) (rate *Rate, response *StandardResponse, err error)
}

// ClientInterface is the Tonicpow client interface
type ClientInterface interface {
	AdvertiserService
	CampaignService
	ConversionService
	GoalService
	RateService
	GetEnvironment() Environment
	GetUserAgent() string
	Options() *ClientOptions
	Request(httpMethod string, requestEndpoint string, data interface{}, expectedCode int) (response *StandardResponse, err error)
	WithCustomHTTPClient(client *resty.Client) *Client
}
