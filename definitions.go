package tonicpow

// APIEnvironment is used to differentiate the environment when making requests
type APIEnvironment string

const (

	// Field key names for various model requests
	fieldAdvertiserProfileID = "advertiser_profile_id"
	fieldApiKey              = "api_key"
	fieldEmail               = "email"
	fieldID                  = "id"
	fieldPassword            = "password"
	fieldPasswordConfirm     = "password_confirm"
	fieldPhone               = "phone"
	fieldPhoneCode           = "phone_code"
	fieldToken               = "token"
	fieldUserID              = "user_id"

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

// AdvertiserProfile is the advertiser_profile model
//
// For more information: https://docs.tonicpow.com/#2f9ec542-0f88-4671-b47c-d0ee390af5ea
type AdvertiserProfile struct {
	HomepageURL string `json:"homepage_url,omitempty"`
	IconURL     string `json:"icon_url,omitempty"`
	ID          uint64 `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	UserID      uint64 `json:"user_id,omitempty"`
}

// Campaign is the campaign model
//
// For more information: https://docs.tonicpow.com/#5aca2fc7-b3c8-445b-aa88-f62a681f8e0c
type Campaign struct {
	AdvertiserProfileID uint64  `json:"advertiser_profile_id,omitempty"`
	Balance             float64 `json:"balance,omitempty"`
	BalanceSatoshis     int64   `json:"balance_satoshis,omitempty"`
	Currency            string  `json:"currency,omitempty"`
	Description         string  `json:"description,omitempty"`
	FundingAddress      string  `json:"funding_address,omitempty"`
	ID                  uint64  `json:"id,omitempty"`
	ImageURL            string  `json:"image_url,omitempty"`
	PayPerClickRate     float64 `json:"pay_per_click_rate,omitempty"`
	PublicGUID          string  `json:"public_guid,omitempty"`
	TargetURL           string  `json:"target_url,omitempty"`
	Title               string  `json:"title,omitempty"`
}
