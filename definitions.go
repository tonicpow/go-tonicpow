package tonicpow

// APIEnvironment is used to differentiate the environment when making requests
type APIEnvironment string

const (

	// defaultUserAgent is the default user agent for all requests
	defaultUserAgent string = "go-tonicpow: v0.5.4"

	// Field key names for various model requests
	fieldAdvertiserProfileID = "advertiser_profile_id"
	fieldAmount              = "amount"
	fieldAPIKey              = "api_key"
	fieldCampaignID          = "campaign_id"
	fieldCurrency            = "currency"
	fieldCurrentPage         = "current_page"
	fieldCustomDimensions    = "custom_dimensions"
	fieldDelayInMinutes      = "delay_in_minutes"
	fieldEmail               = "email"
	fieldExpired             = "expired"
	fieldFeedType            = "feed_type"
	fieldGoalID              = "goal_id"
	fieldID                  = "id"
	fieldLastBalance         = "last_balance"
	fieldMinimumBalance      = "minimum_balance"
	fieldName                = "name"
	fieldPassword            = "password"
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
	modelAdvertiser = "advertisers"
	modelCampaign   = "campaigns"
	modelConversion = "conversions"
	modelGoal       = "goals"
	modelRates      = "rates"
	modelUser       = "users"

	// SortByFieldBalance is for sorting results by field: balance
	SortByFieldBalance = "balance"

	// SortByFieldCreatedAt is for sorting results by field: created_at
	SortByFieldCreatedAt = "created_at"

	// SortByFieldName is for sorting results by field: name
	SortByFieldName = "name"

	// SortByFieldLinksCreated is for sorting results by field: links_created
	SortByFieldLinksCreated = "links_created"

	// SortByFieldPaidClicks is for sorting results by field: paid_clicks
	SortByFieldPaidClicks = "paid_clicks"

	// SortByFieldPayPerClick is for sorting results by field: pay_per_click_rate
	SortByFieldPayPerClick = "pay_per_click_rate"

	// SortOrderAsc is for returning the results in ascending order
	SortOrderAsc = "asc"

	// SortOrderDesc is for returning the results in descending order
	SortOrderDesc = "desc"

	// FeedTypeAtom is for using the feed type: Atom
	FeedTypeAtom = "atom"

	// FeedTypeJSON is for using the feed type: JSON
	FeedTypeJSON = "json"

	// FeedTypeRSS is for using the feed type: RSS
	FeedTypeRSS = "rss"

	// apiVersion current version for all endpoints
	apiVersion = "v1"

	// LiveEnvironment is the live production environment
	LiveEnvironment APIEnvironment = "https://api.tonicpow.com/" + apiVersion + "/"

	// LocalEnvironment is for testing locally using your own api instance
	LocalEnvironment APIEnvironment = "http://localhost:3000/" + apiVersion + "/"

	// StagingEnvironment is used for production-like testing
	StagingEnvironment APIEnvironment = "https://api.staging.tonicpow.com/" + apiVersion + "/"
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

// AdvertiserProfile is the advertiser_profile model (child of User)
//
// For more information: https://docs.tonicpow.com/#2f9ec542-0f88-4671-b47c-d0ee390af5ea
type AdvertiserProfile struct {
	HomepageURL         string `json:"homepage_url"`
	IconURL             string `json:"icon_url"`
	PublicGUID          string `json:"public_guid"`
	Name                string `json:"name"`
	ID                  uint64 `json:"id,omitempty"`
	LinkServiceDomainID uint64 `json:"link_service_domain_id"`
	UserID              uint64 `json:"user_id"`
	DomainVerified      bool   `json:"domain_verified"`
	Unlisted            bool   `json:"unlisted"`
}

// AdvertiserResults is the page response for advertiser profile results from listing
type AdvertiserResults struct {
	Advertisers    []*AdvertiserProfile `json:"advertisers"`
	CurrentPage    int                  `json:"current_page"`
	Results        int                  `json:"results"`
	ResultsPerPage int                  `json:"results_per_page"`
}

// App is the app model (child of advertiser_profile)
//
// For more information: (todo)
type App struct {
	AdvertiserProfileID uint64 `json:"advertiser_profile_id"`
	ID                  uint64 `json:"id"`
	Name                string `json:"name"`
	UserID              uint64 `json:"user_id"`
	WebhookURL          string `json:"webhook_url"`
}

// AppResults is the page response for app results from listing
type AppResults struct {
	Apps           []*App `json:"apps"`
	CurrentPage    int    `json:"current_page"`
	Results        int    `json:"results"`
	ResultsPerPage int    `json:"results_per_page"`
}

// Campaign is the campaign model (child of AdvertiserProfile)
//
// For more information: https://docs.tonicpow.com/#5aca2fc7-b3c8-445b-aa88-f62a681f8e0c
type Campaign struct {
	Goals                 []*Goal               `json:"goals"`
	Images                []*CampaignImage      `json:"images"`
	CreatedAt             string                `json:"created_at"`
	LastEventAt           string                `json:"last_event_at"`
	Currency              string                `json:"currency"`
	Description           string                `json:"description"`
	ExpiresAt             string                `json:"expires_at"`
	FundingAddress        string                `json:"funding_address"`
	ImageURL              string                `json:"image_url"`
	PublicGUID            string                `json:"public_guid"`
	Slug                  string                `json:"slug"`
	TargetURL             string                `json:"target_url"`
	Title                 string                `json:"title"`
	TxID                  string                `json:"-"`
	AdvertiserProfile     *AdvertiserProfile    `json:"advertiser_profile"`
	Balance               float64               `json:"balance"`
	BalanceAlertThreshold float64               `json:"balance_alert_threshold"`
	PayPerClickRate       float64               `json:"pay_per_click_rate"`
	AdvertiserProfileID   uint64                `json:"advertiser_profile_id"`
	BalanceSatoshis       uint64                `json:"balance_satoshis"`
	ID                    uint64                `json:"id,omitempty"`
	LinksCreated          uint64                `json:"links_created"`
	LinkServiceDomainID   uint64                `json:"link_service_domain_id"`
	PaidClicks            uint64                `json:"paid_clicks"`
	PaidConversions       uint64                `json:"paid_conversions"`
	Requirements          *CampaignRequirements `json:"requirements"`
	BotProtection         bool                  `json:"bot_protection"`
	ContributeEnabled     bool                  `json:"contribute_enabled"`
	DomainVerified        bool                  `json:"domain_verified"`
	Unlisted              bool                  `json:"unlisted"`
	MatchDomain           bool                  `json:"match_domain"`
}

// CampaignImage is the structure of the image meta data
type CampaignImage struct {
	Height   int    `json:"height"`
	MimeType string `json:"mime_type"`
	URL      string `json:"url"`
	Width    int    `json:"width"`
}

// CampaignRequirements is the structure for "requirements"
//
// DO NOT CHANGE ORDER - malign
//
type CampaignRequirements struct {
	VisitorCountries    []string `json:"visitor_countries"`
	Application         bool     `json:"application"`
	Facebook            bool     `json:"facebook"`
	Google              bool     `json:"google"`
	HandCash            bool     `json:"handcash"`
	KYC                 bool     `json:"kyc"`
	MoneyButton         bool     `json:"moneybutton"`
	Relay               bool     `json:"relay"`
	Twitter             bool     `json:"twitter"`
	VisitorRestrictions bool     `json:"visitor_restrictions"`
}

// CampaignResults is the page response for campaign results from listing
type CampaignResults struct {
	Campaigns      []*Campaign `json:"campaigns"`
	CurrentPage    int         `json:"current_page"`
	Results        int         `json:"results"`
	ResultsPerPage int         `json:"results_per_page"`
}

// Conversion is the response of getting a conversion
//
// For more information: https://docs.tonicpow.com/#75c837d5-3336-4d87-a686-d80c6f8938b9
type Conversion struct {
	Amount           float64 `json:"amount,omitempty"`
	CampaignID       uint64  `json:"campaign_id"`
	CustomDimensions string  `json:"custom_dimensions"`
	GoalID           uint64  `json:"goal_id"`
	GoalName         string  `json:"goal_name,omitempty"`
	ID               uint64  `json:"id,omitempty"`
	PayoutAfter      string  `json:"payout_after,omitempty"`
	Status           string  `json:"status"`
	StatusData       string  `json:"status_data"`
	TxID             string  `json:"tx_id"`
	UserID           uint64  `json:"user_id"`
}

// Features is the structure for JSON field "features" under user
type Features struct {
	CustomDomains       bool `json:"custom_domains"`
	UnlistedAdvertisers bool `json:"unlisted_advertisers"`
}

// Goal is the goal model (child of Campaign)
//
// For more information: https://docs.tonicpow.com/#316b77ab-4900-4f3d-96a7-e67c00af10ca
type Goal struct {
	CampaignID      uint64  `json:"campaign_id"`
	Description     string  `json:"description"`
	ID              uint64  `json:"id,omitempty"`
	LastConvertedAt string  `json:"last_converted_at"`
	MaxPerPromoter  int16   `json:"max_per_promoter"`
	MaxPerVisitor   int16   `json:"max_per_visitor"`
	Name            string  `json:"name"`
	PayoutRate      float64 `json:"payout_rate"`
	Payouts         int     `json:"payouts"`
	PayoutType      string  `json:"payout_type"`
	Title           string  `json:"title"`
}

// Rate is the rate results
//
// For more information: https://docs.tonicpow.com/#fb00736e-61b9-4ec9-acaf-e3f9bb046c89
type Rate struct {
	Currency            string  `json:"currency"`
	CurrencyAmount      float64 `json:"currency_amount"`
	CurrencyLastUpdated string  `json:"currency_last_updated,omitempty"`
	CurrencyName        string  `json:"currency_name,omitempty"`
	Price               float64 `json:"price"`
	PriceInSatoshis     int64   `json:"price_in_satoshis"`
	RateLastUpdated     string  `json:"rate_last_updated,omitempty"`
}

// User is the user model
//
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// For more information: https://docs.tonicpow.com/#50b3c130-7254-4a05-b312-b14647736e38
type User struct {
	Features           *Features `json:"features"`
	AvatarURL          string    `json:"avatar_url"`
	Bio                string    `json:"bio"`
	Country            string    `json:"country"`
	CreatedAt          string    `json:"created_at"`
	DefaultWallet      string    `json:"default_wallet"`
	Email              string    `json:"email"`
	FirstName          string    `json:"first_name"`
	InternalAddress    string    `json:"internal_address"`
	LastName           string    `json:"last_name"`
	MiddleName         string    `json:"middle_name"`
	NewPassword        string    `json:"new_password,omitempty"`
	NewPasswordConfirm string    `json:"new_password_confirm,omitempty"`
	Password           string    `json:"password,omitempty"`
	PayoutAddress      string    `json:"payout_address"`
	Phone              string    `json:"phone"`
	ReadTosAt          string    `json:"read_tos_at"`
	ReferralURL        string    `json:"referral_url"`
	Status             string    `json:"status"`
	TncpwSession       string    `json:"tncpw_session,omitempty"`
	Username           string    `json:"username"`
	Balance            uint64    `json:"balance"`
	Earned             uint64    `json:"earned"`
	ExperiencePoints   uint64    `json:"experience_points"`
	ID                 uint64    `json:"id,omitempty"`
	ReferralLinkID     uint64    `json:"referral_link_id"`
	ReferredByUserID   uint64    `json:"referred_by_user_id"`
	Referrals          uint      `json:"referrals"`
	ReferralsAccepted  uint      `json:"referrals_accepted"`
	EmailVerified      bool      `json:"email_verified"`
	FacebookAuth       bool      `json:"facebook_auth"`
	GoogleAuth         bool      `json:"google_auth"`
	HandCashAuth       bool      `json:"handcash_auth"`
	HasAlpha           bool      `json:"has_alpha"`
	HasBeta            bool      `json:"has_beta"`
	MoneyButtonAuth    bool      `json:"moneybutton_auth"`
	PhoneVerified      bool      `json:"phone_verified"`
	RelayAuth          bool      `json:"relay_auth"`
	Sounds             bool      `json:"sounds"`
	StarfishAuth       bool      `json:"starfish_auth"`
	TwitterAuth        bool      `json:"twitter_auth"`
}
