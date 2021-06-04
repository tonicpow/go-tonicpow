package tonicpow

import (
	"time"
)

const (
	// Config defaults
	apiVersion         string = "v1"
	defaultHTTPTimeout        = 10 * time.Second          // Default timeout for all GET requests in seconds
	defaultRetryCount  int    = 2                         // Default retry count for HTTP requests
	defaultUserAgent          = "go-tonicpow: " + version // Default user agent
	version            string = "v0.6.0"                  // go-tonicpow version

	// Field key names for various model requests
	fieldAdvertiserProfileID = "advertiser_profile_id"
	fieldAmount              = "amount"
	fieldAPIKey              = "api_key"
	fieldCampaignID          = "campaign_id"
	fieldCurrency            = "currency"
	fieldCurrentPage         = "current_page"
	fieldCustomDimensions    = "custom_dimensions"
	fieldDelayInMinutes      = "delay_in_minutes"
	fieldExpired             = "expired"
	fieldFeedType            = "feed_type"
	fieldGoalID              = "goal_id"
	fieldID                  = "id"
	fieldMinimumBalance      = "minimum_balance"
	fieldName                = "name"
	fieldReason              = "reason"
	fieldResultsPerPage      = "results_per_page"
	fieldSearchQuery         = "query"
	fieldSlug                = "slug"
	fieldSortBy              = "sort_by"
	fieldSortOrder           = "sort_order"
	fieldTargetURL           = "target_url"
	fieldUserID              = "user_id"
	fieldVisitorSessionGUID  = "tncpw_session"

	// Model names (used for Request endpoints)
	modelAdvertiser string = "advertisers"
	modelApp        string = "apps"
	modelCampaign   string = "campaigns"
	modelConversion string = "conversions"
	modelGoal       string = "goals"
	modelRates      string = "rates"

	// Environment names
	environmentDevelopmentAlias string = "local"
	environmentDevelopmentName  string = "development"
	environmentLiveAlias        string = "production"
	environmentLiveName         string = "live"
	environmentStagingAlias     string = "beta"
	environmentStagingName      string = "staging"

	// SortByFieldBalance is for sorting results by field: balance
	SortByFieldBalance string = "balance"

	// SortByFieldCreatedAt is for sorting results by field: created_at
	SortByFieldCreatedAt string = "created_at"

	// SortByFieldName is for sorting results by field: name
	SortByFieldName string = "name"

	// SortByFieldLinksCreated is for sorting results by field: links_created
	SortByFieldLinksCreated string = "links_created"

	// SortByFieldPaidClicks is for sorting results by field: paid_clicks
	SortByFieldPaidClicks string = "paid_clicks"

	// SortByFieldPayPerClick is for sorting results by field: pay_per_click_rate
	SortByFieldPayPerClick string = "pay_per_click_rate"

	// SortOrderAsc is for returning the results in ascending order
	SortOrderAsc string = "asc"

	// SortOrderDesc is for returning the results in descending order
	SortOrderDesc string = "desc"

	// FeedTypeAtom is for using the feed type: Atom
	FeedTypeAtom feedType = "atom"

	// FeedTypeJSON is for using the feed type: JSON
	FeedTypeJSON feedType = "json"

	// FeedTypeRSS is for using the feed type: RSS
	FeedTypeRSS feedType = "rss"
)

var (

	// appSortFields is used for allowing specific fields for sorting
	appSortFields = []string{
		SortByFieldCreatedAt,
		SortByFieldName,
	}

	// campaignSortFields is used for allowing specific fields for sorting
	campaignSortFields = []string{
		SortByFieldBalance,
		SortByFieldCreatedAt,
		SortByFieldLinksCreated,
		SortByFieldPaidClicks,
		SortByFieldPayPerClick,
	}
)

// feedType is used for the campaign feeds (rss, atom, json)
type feedType string

// environment is used for changing the environment for running client requests
type environment struct {
	alias  string
	apiURL string
	name   string
}

// Current environments available
var (
	EnvironmentLive = environment{
		apiURL: "https://api.tonicpow.com/" + apiVersion,
		name:   environmentLiveName,
		alias:  environmentLiveAlias,
	}
	EnvironmentStaging = environment{
		apiURL: "https://api.staging.tonicpow.com/" + apiVersion,
		name:   environmentStagingName,
		alias:  environmentStagingAlias,
	}
	EnvironmentDevelopment = environment{
		apiURL: "http://localhost:3000/" + apiVersion,
		name:   environmentDevelopmentName,
		alias:  environmentDevelopmentAlias,
	}
)

// Error is the universal Error response from the API
//
// For more information: https://docs.tonicpow.com/#d7fe13a3-2b6d-4399-8d0f-1d6b8ad6ebd9
type Error struct {
	Code        int         `json:"code"`
	Data        interface{} `json:"data"`
	IPAddress   string      `json:"ip_address"`
	Message     string      `json:"message"`
	Method      string      `json:"method"`
	RequestGUID string      `json:"request_guid"`
	StatusCode  int         `json:"status_code"`
	URL         string      `json:"url"`
}
