package gowinds

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	libVersion    = "0.1.0"
	baseURL       = "https://striketracker.highwinds.com"
	basePath      = "api/v1/accounts"
	mediaType     = "application/json"
	applicationID = "gowinds/" + libVersion
)

// RequestOptions specifies global API parameters
type RequestOptions struct {
	// AccountHash is required and variable
	AccountHash string `url:"account_hash,omitempty"`
}

func (r *RequestOptions) createURL() (url string) {
	url = fmt.Sprintf("%s/%s", basePath, r.AccountHash)
	return
}

// Client is the core wrapper for all functions in StrikeTracker
type Client struct {
	AuthorizationHeaderToken string
	APIURL                   string
	BaseURL                  *url.URL
	//ApplicationID            string

	// API Services
	Origin OriginService
}

// NewRequest initializes a new http request with headers
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	// Process relative path
	//rel, err := url.Parse(path)
	//if err != nil {
	//	return nil, err
	//}

	//u := c.BaseURL.ResolveReference(rel)

	// json encode the request if exists
	return nil, nil
}

// New creates a new connection client
//func New(authToken string, applicationID string) *Client {
//	newStrike := &Client{
//		AuthorizationHeaderToken: authToken,
//		ApplicationID:            applicationID,
//		APIURL:                   baseURL,
//	}
//
//	return newStrike
//}
