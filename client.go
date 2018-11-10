package gowinds

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const (
	baseURL       = "https://striketracker.highwinds.com"
	basePath      = "api/v1/accounts"
	mediaType     = "application/json"
	applicationID = "gowinds"
)

// RequestOptions specifies global API parameters for every call
type RequestOptions struct {
	// AccountHash is required and variable
	AccountHash string `url:"account_hash,omitempty"`
}

// createURL concatenates the account hash at request time
func (r *RequestOptions) createURL() string {
	url := fmt.Sprintf("%s/%s", basePath, r.AccountHash)
	return url
}

// Response is the API call response
type Response struct {
	*http.Response
}

// Logger interface is the default logger
type logger interface {
	Printf(string, ...interface{})
}

// Client is our core universal API client
type Client struct {
	client                   *http.Client
	debug                    bool
	AuthorizationHeaderToken string
	BaseURL                  *url.URL
	logger
	// Services
}

// SetLogger sets the logger
func (c *Client) SetLogger(l logger) error {
	c.logger = l
	return nil
}

// SetDebug enables/disables debug logging after creating a client
func (c *Client) SetDebug(toggle bool) {
	c.debug = toggle
	return
}

// SetBaseURL Sets optional BaseURL just in case - mostly for tests
func (c *Client) SetBaseURL(u string) error {
	url, err := url.Parse(u)
	if err != nil {
		return err
	}
	c.BaseURL = url
	return nil
}

// NewClient returns a copy of the client
func NewClient(authorizationHeaderToken string) (*Client, error) {

	// Fetch auth data
	if authorizationHeaderToken == "" || len(authorizationHeaderToken) == 0 {
		return nil, fmt.Errorf("authorizationHeaderToken required")
	}

	// URL object parsed
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{client: &http.Client{}, AuthorizationHeaderToken: authorizationHeaderToken, BaseURL: u, debug: false}
	// Mount services
	// c.Analytics
	// c.Hosts
	return c, nil
}

// NewRequest packgs a new http request
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	// validate and have obj
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// build url and queries
	u := c.BaseURL.ResolveReference(rel)
	v, _ := query.Values(body)
	u.RawQuery = v.Encode()

	// create raw request, body is not used on striketracker
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	if c.debug {
		c.logger.Printf("Request: ", req)
	}

	req.Close = true

	// pack headers
	req.Header.Add("Authorization", c.AuthorizationHeaderToken)
	req.Header.Add("X-Application-Id", applicationID)
	req.Header.Add("Content-Type", mediaType)

	return req, nil
}

// Do fires the request
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := Response{Response: resp}

	if v != nil {
		// Just in case we use an io.Writer to decode elsewhere
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			// Decode in the struct injected in v
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return &response, err
			}
		}
	}

	return &response, err
}

// DoRequest creates and fires a request
func (c *Client) DoRequest(method, path string, body, v interface{}) (*Response, error) {
	req, err := c.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	return c.Do(req, v)
}
