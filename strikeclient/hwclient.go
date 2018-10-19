package strikeclient

// StrikeClient is the core wrapper for all functions in StrikeTracker
type StrikeClient struct {
	AuthorizationHeaderToken string
	AccountHash              string
	APIURL                   string
}

// New creates a new connection client
func New(authToken string, accountHash string) *StrikeClient {
	newStrike := &StrikeClient{
		AuthorizationHeaderToken: authToken,
		AccountHash:              accountHash, // Maybe lift AccountHash out of here in the future
		APIURL:                   "https://striketracker.highwinds.com",
	}

	return newStrike
}
