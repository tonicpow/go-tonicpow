package tonicpow

// APIEnvironment is used to differentiate the environment when making requests
type APIEnvironment string

const (

	// APIKeyName is the key name for requests
	APIKeyName = "api_key"

	// APIVersion current version for all endpoints
	APIVersion = "v1"

	// DefaultUserAgent is the default user agent for all requests
	DefaultUserAgent string = "go-tonicpow: " + APIVersion

	// LiveEnvironment is the live production environment
	LiveEnvironment APIEnvironment = "https://api.tonicpow.com/" + APIVersion + "/"

	// LocalEnvironment is for testing locally using your own api instance
	LocalEnvironment APIEnvironment = "http://localhost:3000/" + APIVersion + "/"

	// StagingEnvironment is used for production-like testing
	StagingEnvironment APIEnvironment = "https://apistaging.tonicpow.com/" + APIVersion + "/"

	// SessionCookie is the cookie name for session tokens
	SessionCookie = "session_token"

	// TestEnvironment is a test-only environment
	//TestEnvironment APIEnvironment = "https://test.tonicpow.com/"+APIVersion+"/"
)

// Error is the universal error response from the API
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
