package gowinds

import (
	"fmt"
)

// RequestOptions specifies global API parameters
type RequestOptions struct {
	// AccountHash is required and variable
	AccountHash string `url:"account_hash,omitempty"`
}

func (r *RequestOptions) createURL() (url string) {
	url = fmt.Sprintf("/api/v1/accounts/%s", r.AccountHash)
	return
}

// Client is the core wrapper for all functions in StrikeTracker
type Client struct {
	AuthorizationHeaderToken string
	APIURL                   string
	ApplicationID            string

	//AccountHash              string
	// API Services
	Origin OriginService
}

// New creates a new connection client
func New(authToken string, applicationID string) *Client {
	newStrike := &Client{
		AuthorizationHeaderToken: authToken,
		ApplicationID:            applicationID,
		APIURL:                   "https://striketracker.highwinds.com",
	}

	return newStrike
}
