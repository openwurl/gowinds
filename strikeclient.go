package gowinds

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	authorizationHeaderToken = "AuthorizationHeaderToken"
	libVersion               = "0.1.0"
	baseURL                  = "https://striketracker.highwinds.com"
	basePath                 = "api/v1/accounts"
	mediaType                = "application/json"
	applicationID            = "gowinds/" + libVersion
)

// RequestOptions specifies global API parameters
type RequestOptions struct {
	// AccountHash is required and variable
	AccountHash string `url:"account_hash,omitempty"`
}

// Response is the API call response
type Response struct {
	*http.Response
}

func (r *RequestOptions) createURL() (url string) {
	url = fmt.Sprintf("%s/%s", basePath, r.AccountHash)
	return
}

// Client is the core wrapper for all functions in StrikeTracker
type Client struct {
	client *http.Client

	AuthorizationHeaderToken string
	BaseURL                  *url.URL

	// API Services
	Origin OriginService
}

// NewClient initializes and retursn a base client
func NewClient() (*Client, error) {
	// goconfig this?
	authorizationHeaderToken := os.Getenv(authorizationHeaderToken)
	if authorizationHeaderToken == "" {
		return nil, fmt.Errorf("you must export %s", authorizationHeaderToken)
	}

	/*
		if httpClient == nil {
			httpClient = &http.Client{}
		}
	*/

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{client: &http.Client{}, AuthorizationHeaderToken: authorizationHeaderToken, BaseURL: u}
	c.Origin = &OriginServiceOp{client: c}

	return c, nil
}

// NewRequest initializes a new http request with headers
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	// Process relative path
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	// json encode the request if exists
	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Close = true

	req.Header.Add("Authorization", c.AuthorizationHeaderToken)
	req.Header.Add("X-Application-Id", applicationID)
	req.Header.Add("Content-Type", mediaType)

	return req, nil
}

// Do executes an http request
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := Response{Response: resp}

	// check for !200-299 response
	// io.Copy to new io.Writer and check body
	// for api error payload before passing back

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return &response, err
			}
		}
	}

	return &response, err
}

// DoRequest for fast requesting bro
func (c *Client) DoRequest(method, path string, body, v interface{}) (*Response, error) {
	req, err := c.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	return c.Do(req, v)
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
