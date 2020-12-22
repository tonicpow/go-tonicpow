package tonicpow

// APIEnvironment is used to differentiate the environment when making requests
type APIEnvironment string

const (

	// Field key names for various model requests
	fieldAdvertiserProfileID = "advertiser_profile_id"
	fieldAmount              = "amount"
	fieldAPIKey              = "api_key"
	fieldAppID               = "app_id"
	fieldCampaignID          = "campaign_id"
	fieldCurrency            = "currency"
	fieldCurrentPage         = "current_page"
	fieldCustomDimensions    = "custom_dimensions"
	fieldDelayInMinutes      = "delay_in_minutes"
	fieldEmail               = "email"
	fieldFeedType            = "feed_type"
	fieldGoalID              = "goal_id"
	fieldID                  = "id"
	fieldLabel               = "label"
	fieldLastBalance         = "last_balance"
	fieldLinkID              = "link_id"
	fieldName                = "name"
	fieldPassword            = "password"
	fieldPasswordConfirm     = "password_confirm"
	fieldReason              = "reason"
	fieldResultsPerPage      = "results_per_page"
	fieldSearchQuery         = "query"
	fieldShortCode           = "short_code"
	fieldSortBy              = "sort_by"
	fieldSortOrder           = "sort_order"
	fieldTargetURL           = "target_url"
	fieldToken               = "token"
	fieldUserID              = "user_id"
	fieldVisitorSessionGUID  = "tncpw_session"
	fieldWidgetID            = "widget_id"

	// Model names (used for Request endpoints)
	modelAdvertiser     = "advertisers"
	modelApp            = "apps"
	modelCampaign       = "campaigns"
	modelConversion     = "conversions"
	modelGoal           = "goals"
	modelLink           = "links"
	modelRates          = "rates"
	modelUser           = "users"
	modelVisitors       = "visitors"
	modelVisitorSession = "sessions"
	modelWidget         = "widgets"

	// SortByFieldBalance is for sorting results by field: balance
	SortByFieldBalance = "balance"

	// SortByFieldCreatedAt is for sorting results by field: created_at
	SortByFieldCreatedAt = "created_at"

	// SortByFieldEarned is for sorting results by field: earned
	SortByFieldEarned = "earned"

	// SortByFieldName is for sorting results by field: name
	SortByFieldName = "name"

	// SortByFieldLinksCreated is for sorting results by field: links_created
	SortByFieldLinksCreated = "links_created"

	// SortByFieldPaidClicks is for sorting results by field: paid_clicks
	SortByFieldPaidClicks = "paid_clicks"

	// SortByFieldPayPerClick is for sorting results by field: pay_per_click_rate
	SortByFieldPayPerClick = "pay_per_click_rate"

	// SortByFieldReferrals is for sorting results by field: referrals
	SortByFieldReferrals = "referrals"

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

	// defaultUserAgent is the default user agent for all requests
	defaultUserAgent string = "go-tonicpow: v0.4.49"

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

	// referralSortFields is used for allowing specific fields for sorting
	referralSortFields = []string{
		SortByFieldCreatedAt,
		SortByFieldEarned,
		SortByFieldReferrals,
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

// ActivityItem is the item for a recent activity request
type ActivityItem struct {
	Action       string `json:"action"`
	Amount       uint64 `json:"amount"`
	CampaignID   uint64 `json:"campaign_id"`
	CampaignSlug string `json:"campaign_slug"`
	Date         string `json:"date"`
	TxID         string `json:"tx_id"`
}

// AdvertiserProfile is the advertiser_profile model (child of User)
//
// For more information: https://docs.tonicpow.com/#2f9ec542-0f88-4671-b47c-d0ee390af5ea
type AdvertiserProfile struct {
	DomainVerified bool   `json:"domain_verified"`
	HomepageURL    string `json:"homepage_url"`
	IconURL        string `json:"icon_url"`
	ID             uint64 `json:"id,omitempty"`
	Name           string `json:"name"`
	UserID         uint64 `json:"user_id"`
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

// APIKey is the api_key model (child of app)
//
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// For more information: (todo)
type APIKey struct {
	CreatedAt   string `json:"created_at"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Preview     string `json:"preview"`
	ScopeModel  string `json:"scope_model"`
	AppID       uint64 `json:"app_id"`
	ID          uint64 `json:"id"`
	ScopeID     uint64 `json:"scope_id"`
	UserID      uint64 `json:"user_id"`
	Active      bool   `json:"active"`
	WriteAccess bool   `json:"write_access"`
}

// APIKeyResults is the page response for listing api_keys
type APIKeyResults struct {
	APIKeys        []*APIKey `json:"api_keys"`
	CurrentPage    int       `json:"current_page"`
	Results        int       `json:"results"`
	ResultsPerPage int       `json:"results_per_page"`
}

// Blocked is the blocked_user model
//
// For more information: (todo)
type Blocked struct {
	AdvertiserProfileID uint64 `json:"advertiser_profile_id"`
	BlockedUserID       uint64 `json:"blocked_user_id"`
	CampaignID          uint64 `json:"campaign_id"`
	ID                  uint64 `json:"id"`
	Reason              string `json:"reason"`
	UserID              uint64 `json:"user_id"`
}

// BlockedResults is the page response for listing blocks
type BlockedResults struct {
	Blocks         []*Blocked `json:"blocks"`
	CurrentPage    int        `json:"current_page"`
	Results        int        `json:"results"`
	ResultsPerPage int        `json:"results_per_page"`
}

// BasicStatistics is the statistics about a given campaign or advertiser
type BasicStatistics struct {
	Clicks           []*DataPoint `json:"clicks"`
	Conversions      []*DataPoint `json:"conversions"`
	TotalClicks      int          `json:"total_clicks"`
	TotalConversions int          `json:"total_conversions"`
	TotalSatoshis    int          `json:"total_satoshis"`
}

// DataPoint a data point to be plotted on the chart
type DataPoint struct {
	Count    uint64 `json:"count"`
	Date     string `json:"date"`
	Satoshis uint64 `json:"satoshis"`
}

// Campaign is the campaign model (child of AdvertiserProfile)
//
// For more information: https://docs.tonicpow.com/#5aca2fc7-b3c8-445b-aa88-f62a681f8e0c
type Campaign struct {
	Goals               []*Goal               `json:"goals"`
	Images              []*CampaignImage      `json:"images"`
	CreatedAt           string                `json:"created_at"`
	Currency            string                `json:"currency"`
	Description         string                `json:"description"`
	ExpiresAt           string                `json:"expires_at"`
	FundingAddress      string                `json:"funding_address"`
	ImageURL            string                `json:"image_url"`
	PublicGUID          string                `json:"public_guid"`
	Slug                string                `json:"slug"`
	TargetURL           string                `json:"target_url"`
	Title               string                `json:"title"`
	TxID                string                `json:"-"`
	AdvertiserProfile   *AdvertiserProfile    `json:"advertiser_profile"`
	AdvertiserProfileID uint64                `json:"advertiser_profile_id"`
	PayPerClickRate     float64               `json:"pay_per_click_rate"`
	Balance             float64               `json:"balance"`
	PaidClicks          uint64                `json:"paid_clicks"`
	LinksCreated        uint64                `json:"links_created"`
	ID                  uint64                `json:"id,omitempty"`
	BalanceSatoshis     uint64                `json:"balance_satoshis"`
	Requirements        *CampaignRequirements `json:"requirements"`
	DomainVerified      bool                  `json:"domain_verified"`
	MatchDomain         bool                  `json:"match_domain"`
	BotProtection       bool                  `json:"bot_protection"`
	ContributeEnabled   bool                  `json:"contribute_enabled"`
}

// CampaignApplication is the structure of the campaign application data
type CampaignApplication struct {
	AdvertiserProfileID uint64 `json:"advertiser_profile_id"`
	CampaignID          uint64 `json:"campaign_id"`
	ID                  uint64 `json:"id"`
	Response            string `json:"response"`
	Status              string `json:"status"`
	Submission          string `json:"submission"`
	UserID              uint64 `json:"user_id"`
}

// CampaignImage is the structure of the image meta data
type CampaignImage struct {
	Height   int    `json:"height"`
	MimeType string `json:"mime_type"`
	URL      string `json:"url"`
	Width    int    `json:"width"`
}

// CampaignRequirements is the structure for "requirements"
type CampaignRequirements struct {
	Application bool `json:"application"`
	HandCash    bool `json:"handcash"`
	KYC         bool `json:"kyc"`
	MoneyButton bool `json:"moneybutton"`
	Relay       bool `json:"relay"`
	Twitter     bool `json:"twitter"`
}

// CampaignResults is the page response for campaign results from listing
type CampaignResults struct {
	Campaigns      []*Campaign `json:"campaigns"`
	CurrentPage    int         `json:"current_page"`
	Results        int         `json:"results"`
	ResultsPerPage int         `json:"results_per_page"`
}

// CampaignStatistics is what we cache for redis for basic stats
type CampaignStatistics struct {
	Active              uint64  `json:"active"`
	Balance             float64 `json:"balance"`
	BalanceSatoshis     uint64  `json:"balance_satoshis"`
	Currency            string  `json:"currency"`
	Total               uint64  `json:"total"`
	TotalEarned         float64 `json:"total_earned"`
	TotalEarnedSatoshis uint64  `json:"total_earned_satoshis"`
}

// Conversion is the response of getting a conversion
//
// For more information: https://docs.tonicpow.com/#75c837d5-3336-4d87-a686-d80c6f8938b9
type Conversion struct {
	Amount           float64 `json:"amount"`
	CampaignID       uint64  `json:"campaign_id"`
	CustomDimensions string  `json:"custom_dimensions"`
	GoalID           uint64  `json:"goal_id"`
	GoalName         string  `json:"goal_name"`
	ID               uint64  `json:"ID,omitempty"`
	PayoutAfter      string  `json:"payout_after"`
	Status           string  `json:"status"`
	TxID             string  `json:"tx_id"`
	UserID           uint64  `json:"user_id"`
}

// Domain is returned after creating a domain
type Domain struct {
	CnameName  string `json:"cname_name"`
	CnameValue string `json:"cname_value"`
	DomainName string `json:"domain_name"`
	ID         uint64 `json:"id"`
	Label      string `json:"label"`
	UserID     uint64 `json:"user_id"`
	Verified   bool   `json:"verified"`
}

// DomainResults is the page response for listing domains
type DomainResults struct {
	CurrentPage    int       `json:"current_page"`
	Domains        []*Domain `json:"domains"`
	Results        int       `json:"results"`
	ResultsPerPage int       `json:"results_per_page"`
}

// Goal is the goal model (child of Campaign)
//
// For more information: https://docs.tonicpow.com/#316b77ab-4900-4f3d-96a7-e67c00af10ca
type Goal struct {
	CampaignID     uint64  `json:"campaign_id"`
	Description    string  `json:"description"`
	ID             uint64  `json:"id,omitempty"`
	MaxPerPromoter int16   `json:"max_per_promoter"`
	MaxPerVisitor  int16   `json:"max_per_visitor"`
	Name           string  `json:"name"`
	PayoutRate     float64 `json:"payout_rate"`
	Payouts        int     `json:"payouts"`
	PayoutType     string  `json:"payout_type"`
	Title          string  `json:"title"`
}

// Link is the link model (child of User) (relates Campaign to User)
// Use the CustomShortCode on create for using your own short code
//
// For more information: https://docs.tonicpow.com/#ee74c3ce-b4df-4d57-abf2-ccf3a80e4e1e
type Link struct {
	AvatarURL          string `json:"avatar_url,omitempty"`
	CampaignID         uint64 `json:"campaign_id"`
	CampaignImageURL   string `json:"image_url"`
	CampaignSlug       string `json:"slug"`
	CampaignTitle      string `json:"title"`
	ClickSatoshis      int32  `json:"click_satoshis"`
	ConversionSatoshis int32  `json:"conversion_satoshis"`
	CustomShortCode    string `json:"custom_short_code"`
	ID                 uint64 `json:"id,omitempty"`
	Label              string `json:"label"`
	ShortCode          string `json:"short_code"`
	ShortLinkURL       string `json:"short_link_url"`
	TargetURL          string `json:"target_url"`
	TotalClicks        int32  `json:"total_clicks"`
	TotalConversions   int32  `json:"total_conversions"`
	UserID             uint64 `json:"user_id"`
	Username           string `json:"username,omitempty"`
}

// LinkResults is the page response for link results from listing
type LinkResults struct {
	CurrentPage    int     `json:"current_page"`
	Links          []*Link `json:"links"`
	Results        int     `json:"results"`
	ResultsPerPage int     `json:"results_per_page"`
}

// Promoter is the public promoter response
type Promoter struct {
	AvatarURL       string `json:"avatar_url"`
	Bio             string `json:"bio"`
	Country         string `json:"country"`
	EmailVerified   bool   `json:"email_verified"`
	HandCashAuth    bool   `json:"handcash_auth"`
	MoneyButtonAuth bool   `json:"moneybutton_auth"`
	PhoneVerified   bool   `json:"phone_verified"`
	RelayAuth       bool   `json:"relay_auth"`
	TwitterAuth     bool   `json:"twitter_auth"`
	Username        string `json:"username"`
}

// PromoterResults is the page response for promoter results from listing
type PromoterResults struct {
	Promoters      []*Promoter `json:"promoters"`
	CurrentPage    int         `json:"current_page"`
	Results        int         `json:"results"`
	ResultsPerPage int         `json:"results_per_page"`
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

// RecentActivityResults is the page response for listing recent activity
type RecentActivityResults struct {
	Activities     []*ActivityItem `json:"activities"`
	CurrentPage    int             `json:"current_page"`
	Results        int             `json:"results"`
	ResultsPerPage int             `json:"results_per_page"`
}

// ReferralResults is the page response for referral results from listing
type ReferralResults struct {
	CurrentPage    int             `json:"current_page"`
	Referrals      []*UserReferral `json:"referrals"`
	Results        int             `json:"results"`
	ResultsPerPage int             `json:"results_per_page"`
}

// User is the user model
//
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// For more information: https://docs.tonicpow.com/#50b3c130-7254-4a05-b312-b14647736e38
type User struct {
	AvatarURL          string `json:"avatar_url"`
	Bio                string `json:"bio"`
	Email              string `json:"email"`
	Country            string `json:"country"`
	DefaultWallet      string `json:"default_wallet"`
	FirstName          string `json:"first_name"`
	InternalAddress    string `json:"internal_address"`
	LastName           string `json:"last_name"`
	MiddleName         string `json:"middle_name"`
	NewPassword        string `json:"new_password,omitempty"`
	NewPasswordConfirm string `json:"new_password_confirm,omitempty"`
	Password           string `json:"password,omitempty"`
	PayoutAddress      string `json:"payout_address"`
	Phone              string `json:"phone"`
	ReferralURL        string `json:"referral_url"`
	Status             string `json:"status"`
	TncpwSession       string `json:"tncpw_session,omitempty"`
	Username           string `json:"username"`
	Balance            uint64 `json:"balance"`
	Earned             uint64 `json:"earned"`
	ExperiencePoints   uint64 `json:"experience_points"`
	ID                 uint64 `json:"id,omitempty"`
	ReferralLinkID     uint64 `json:"referral_link_id"`
	ReferredByUserID   uint64 `json:"referred_by_user_id"`
	Referrals          uint   `json:"referrals"`
	ReferralsAccepted  uint   `json:"referrals_accepted"`
	EmailVerified      bool   `json:"email_verified"`
	HandCashAuth       bool   `json:"handcash_auth"`
	MoneyButtonAuth    bool   `json:"moneybutton_auth"`
	PhoneVerified      bool   `json:"phone_verified"`
	RelayAuth          bool   `json:"relay_auth"`
	TwitterAuth        bool   `json:"twitter_auth"`
}

// UserExists is a slim record of the User model
//
// For more information: https://docs.tonicpow.com/#50b3c130-7254-4a05-b312-b14647736e38
type UserExists struct {
	ID          uint64 `json:"id"`
	ReferralURL string `json:"referral_url"`
	Status      string `json:"status"`
}

// UserReferral is a slim record of the User model
//
// For more information: https://docs.tonicpow.com/#50b3c130-7254-4a05-b312-b14647736e38
type UserReferral struct {
	Email            string `json:"email"`
	ID               uint64 `json:"id"`
	PayoutAddress    string `json:"payout_address"`
	Referrals        uint   `json:"referrals"`
	ReferredByUserID uint64 `json:"referred_by_user_id"`
	Status           string `json:"status"`
}

// VisitorSession is the session for any visitor or user (related to link and campaign)
//
// For more information: https://docs.tonicpow.com/#d0d9055a-0c92-4f55-a370-762d44acf801
type VisitorSession struct {
	CampaignID       uint64 `json:"campaign_id"`
	CustomDimensions string `json:"custom_dimensions"`
	IPAddress        string `json:"ip_address"`
	LinkID           uint64 `json:"link_id"`
	LinkUserID       uint64 `json:"link_user_id"`
	Provider         string `json:"provider"`
	Referer          string `json:"referer"`
	TncpwSession     string `json:"tncpw_session,omitempty"`
	UserAgent        string `json:"user_agent"`
}

// Widget is returned after creating a widget
type Widget struct {
	AcceptedDomains string `json:"accepted_domains"`
	Height          int    `json:"height"`
	ID              uint64 `json:"id"`
	Label           string `json:"label"`
	TxID            string `json:"tx_id"`
	UserID          uint64 `json:"user_id"`
	Width           int    `json:"width"`
}

// WidgetResults is the page response for listing widgets
type WidgetResults struct {
	CurrentPage    int       `json:"current_page"`
	Results        int       `json:"results"`
	ResultsPerPage int       `json:"results_per_page"`
	Widgets        []*Widget `json:"widgets"`
}

// Withdrawal is returned after creating a withdrawal
type Withdrawal struct {
	Amount               uint64 `json:"amount"`
	AmountRejected       uint64 `json:"amount_rejected"`
	Fee                  uint64 `json:"fee"`
	ID                   uint64 `json:"id"`
	PayoutAddress        string `json:"payout_address"`
	UserID               uint64 `json:"user_id"`
	ProvidedInformation  string `json:"provided_information"`
	Response             string `json:"response"`
	Status               string `json:"status"`
	Transactions         uint32 `json:"transactions"`
	TransactionsRejected uint32 `json:"transactions_rejected"`
}

// WithdrawalResults is the page response for listing withdrawals
type WithdrawalResults struct {
	CurrentPage    int           `json:"current_page"`
	Results        int           `json:"results"`
	ResultsPerPage int           `json:"results_per_page"`
	Withdrawals    []*Withdrawal `json:"withdrawals"`
}
