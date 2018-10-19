package gowinds

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
