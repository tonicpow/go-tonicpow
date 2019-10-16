package tonicpow

// ConversionResponse is the structure response from a conversion
type ConversionResponse struct {
	Error
	AdditionalData     string `json:"additional_data"`
	ConversionGoalID   uint64 `json:"conversion_goal_id"`
	ConversionGoalName string `json:"conversion_goal_name"`
	ID                 uint64 `json:"id"`
	UserID             uint64 `json:"user_id"`
	ConversionTxID     string `json:"conversion_tx_id"`
	PayoutTxID         string `json:"payout_tx_id"`
}

// Error is the response from the request
type Error struct {
	Code        int    `json:"code"`
	Data        string `json:"data"`
	IPAddress   string `json:"ip_address"`
	Method      string `json:"method"`
	Message     string `json:"message"`
	RequestGUID string `json:"request_guid"`
	URL         string `json:"url"`
}
