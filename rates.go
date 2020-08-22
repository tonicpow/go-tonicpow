package tonicpow

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetCurrentRate will get a current rate for the given currency (using default currency amount)
//
// For more information: https://docs.tonicpow.com/#71b8b7fc-317a-4e68-bd2a-5b0da012361c
func (c *Client) GetCurrentRate(currency string, customAmount float64) (rate *Rate, err error) {

	// Must have an currency
	if len(currency) == 0 {
		err = fmt.Errorf("missing field: %s", fieldCurrency)
		return
	}

	// Fire the request
	var response string
	if response, err = c.request(fmt.Sprintf("%s/%s?%s=%f", modelRates, currency, fieldAmount, customAmount), http.MethodGet, nil); err != nil {
		return
	}

	// Only a 200 is treated as a success
	if err = c.error(http.StatusOK, response); err != nil {
		return
	}

	// Convert model response
	err = json.Unmarshal([]byte(response), &rate)
	return
}
