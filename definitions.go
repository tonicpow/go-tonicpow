package tonicpow

// APIEnvironment is used to differentiate the environment when making requests
type APIEnvironment string

const (

	// Field key names for various model requests
	fieldAdditionalData      = "additional_data"
	fieldAdvertiserProfileID = "advertiser_profile_id"
	fieldApiKey              = "api_key"
	fieldEmail               = "email"
	fieldID                  = "id"
	fieldName                = "name"
	fieldPassword            = "password"
	fieldPasswordConfirm     = "password_confirm"
	fieldPhone               = "phone"
	fieldPhoneCode           = "phone_code"
	fieldToken               = "token"
	fieldUserID              = "user_id"
	fieldCampaignID          = "campaign_id"
	fieldVisitorSessionID    = "visitor_session_id"
	fieldShortCode           = "short_code"

	// Model names (used for request endpoints)
	modelAdvertiser = "advertisers"
	modelCampaign   = "campaigns"
	modelGoal       = "goals"
	modelLink       = "links"
	modelUser       = "users"

	// apiVersion current version for all endpoints
	apiVersion = "v1"

	// defaultUserAgent is the default user agent for all requests
	defaultUserAgent string = "go-tonicpow: " + apiVersion

	// LiveEnvironment is the live production environment
	LiveEnvironment APIEnvironment = "https://api.tonicpow.com/" + apiVersion + "/"

	// LocalEnvironment is for testing locally using your own api instance
	LocalEnvironment APIEnvironment = "http://localhost:3000/" + apiVersion + "/"

	// StagingEnvironment is used for production-like testing
	StagingEnvironment APIEnvironment = "https://apistaging.tonicpow.com/" + apiVersion + "/"

	// sessionCookie is the cookie name for session tokens
	sessionCookie = "session_token"

	// TestEnvironment is a test-only environment
	//TestEnvironment APIEnvironment = "https://test.tonicpow.com/"+apiVersion+"/"
)

// Error is the universal error response from the API
//
// For more information: https://docs.tonicpow.com/#d7fe13a3-2b6d-4399-8d0f-1d6b8ad6ebd9
type Error struct {
	Code        int    `json:"code"`
	Data        string `json:"data"`
	IPAddress   string `json:"ip_address"`
	Method      string `json:"method"`
	Message     string `json:"message"`
	RequestGUID string `json:"request_guid"`
	URL         string `json:"url"`
}

// User is the user model
//
// For more information: https://docs.tonicpow.com/#50b3c130-7254-4a05-b312-b14647736e38
type User struct {
	Balance            uint64 `json:"balance,omitempty"`
	Email              string `json:"email,omitempty"`
	FirstName          string `json:"first_name,omitempty"`
	ID                 uint64 `json:"id,omitempty"`
	InternalAddress    string `json:"internal_address,omitempty"`
	LastName           string `json:"last_name,omitempty"`
	MiddleName         string `json:"middle_name,omitempty"`
	NewPassword        string `json:"new_password,omitempty"`
	NewPasswordConfirm string `json:"new_password_confirm,omitempty"`
	Password           string `json:"password,omitempty"`
	PayoutAddress      string `json:"payout_address,omitempty"`
	Phone              string `json:"phone,omitempty"`
	Status             string `json:"status,omitempty"`
}

// AdvertiserProfile is the advertiser_profile model (child of User)
//
// For more information: https://docs.tonicpow.com/#2f9ec542-0f88-4671-b47c-d0ee390af5ea
type AdvertiserProfile struct {
	HomepageURL string `json:"homepage_url,omitempty"`
	IconURL     string `json:"icon_url,omitempty"`
	ID          uint64 `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	UserID      uint64 `json:"user_id,omitempty"`
}

// Campaign is the campaign model (child of AdvertiserProfile)
//
// For more information: https://docs.tonicpow.com/#5aca2fc7-b3c8-445b-aa88-f62a681f8e0c
type Campaign struct {
	AdvertiserProfileID uint64  `json:"advertiser_profile_id,omitempty"`
	Balance             float64 `json:"balance,omitempty"`
	BalanceSatoshis     int64   `json:"balance_satoshis,omitempty"`
	Currency            string  `json:"currency,omitempty"`
	Description         string  `json:"description,omitempty"`
	FundingAddress      string  `json:"funding_address,omitempty"`
	Goals               []*Goal `json:"goals,omitempty"`
	ID                  uint64  `json:"id,omitempty"`
	ImageURL            string  `json:"image_url,omitempty"`
	PayPerClickRate     float64 `json:"pay_per_click_rate,omitempty"`
	PublicGUID          string  `json:"public_guid,omitempty"`
	TargetURL           string  `json:"target_url,omitempty"`
	Title               string  `json:"title,omitempty"`
}

// Goal is the goal model (child of Campaign)
//
// For more information: https://docs.tonicpow.com/#316b77ab-4900-4f3d-96a7-e67c00af10ca
type Goal struct {
	CampaignID  uint64  `json:"campaign_id,omitempty"`
	Description string  `json:"description,omitempty"`
	ID          uint64  `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	PayoutRate  float64 `json:"payout_rate,omitempty"`
	Payouts     int     `json:"payouts,omitempty"`
	PayoutType  string  `json:"payout_type,omitempty"`
	Title       string  `json:"title,omitempty"`
}

// Conversion is the result of goal.Convert()
//
// For more information: https://docs.tonicpow.com/#caeffdd5-eaad-4fc8-ac01-8288b50e8e27
type Conversion struct {
	AdditionalData string `json:"additional_data,omitempty"`
	ConversionTxID string `json:"conversion_tx_id,omitempty"`
	GoalID         uint64 `json:"goal_id,omitempty"`
	GoalName       string `json:"goal_name,omitempty"`
	UserID         string `json:"user_id,omitempty"`
}

// Link is the link model (child of User) (relates Campaign to User)
// Use the CustomShortCode on create for using your own short code
//
// For more information: https://docs.tonicpow.com/#ee74c3ce-b4df-4d57-abf2-ccf3a80e4e1e
type Link struct {
	CampaignID      uint64 `json:"campaign_id,omitempty"`
	CustomShortCode string `json:"custom_short_code,omitempty"`
	ID              uint64 `json:"id,omitempty"`
	ShortCode       string `json:"short_code,omitempty"`
	ShortCodeURL    string `json:"short_code_url,omitempty"`
	UserID          uint64 `json:"user_id,omitempty"`
}
