package gowinds

import "fmt"

// /api/v1/accounts/{account_hash}/origins/ | {origin_id}

// OriginService interface defines available origin methods
type OriginService interface {
	Create()
	List()
	Delete()
	Get()
	Update()
}

// OriginOptions specifies origin specific parameters
type OriginOptions struct {
	OriginID string `url:"origin_id,omitempty"`
}

func (o *OriginOptions) createPath() (url string) {
	if o.OriginID != "" {
		url = fmt.Sprintf("%s", o.OriginID)
	}

	return
}

// Origin represents an origin object
type Origin struct {
	AuthenticationType           string `json:"authenticationType,omitempty"`
	CertificateCN                string `json:"certificateCN,omitempty"`
	CreatedDate                  string `json:"createdDate,omitempty"`
	ErrorCacheTTLSeconds         int    `json:"errorCacheTTLSeconds,omitempty"`
	Hostname                     string `json:"hostname"`
	MaxConnectionsPerEdge        int    `json:"maxConnectionsPerEdge,omitempty"`
	MaxConnectionsPerEdgeEnabled bool   `json:"maxConnectionsPerEdgeEnabled,omitempty"`
	MaxRetryCount                int    `json:"maxRetryCount,omitempty"`
	MaximumOriginPullSeconds     int    `json:"maximumOriginPullSeconds,omitempty"`
	Name                         string `json:"name"`
	OriginCacheHeaders           string `json:"originCacheHeaders,omitempty"`
	OriginDefaultKeepAlive       string `json:"originDefaultKeepAlive,omitempty"`
	OriginPullHeaders            string `json:"originPullHeaders,omitempty"`
	OriginPullNegLinger          string `json:"originPullNegLinger,omitempty"`
	Path                         string `json:"path,omitempty"`
	Port                         int    `json:"port"`
	RequestTimeoutSeconds        int    `json:"requestTimeoutSeconds,omitempty"`
	SecurePort                   int    `json:"securePort,omitempty"`
	Type                         string `json:"type,omitempty"`
	UpdatedDate                  string `json:"updatedDate,omitempty"`
	VerifyCertificate            bool   `json:"verifyCertificate,omitempty"`
}

// OriginList represents a list of origins
type OriginList struct {
	List []Origin `json:"list"`
}

// OriginServiceOp implements OriginServie
type OriginServiceOp struct {
	client *Client
}

// Create creates a new origin
func (s *OriginServiceOp) Create() {

}

// List returns the account's origins
func (s *OriginServiceOp) List(reqOpt *RequestOptions, origOpt *OriginOptions) (*OriginList, error) {
	if reqOpt == nil {
		return nil, fmt.Errorf("no account hash defined")
	}

	var requestPath string
	if origOpt != nil {
		requestPath = origOpt.createPath()
	}

	path := fmt.Sprintf("%s?%s", reqOpt.createURL(), requestPath)

	fmt.Println(path)
	return nil, nil
}

// Delete terminates an origin
func (s *OriginServiceOp) Delete() {

}

// Get returns an individual origin
func (s *OriginServiceOp) Get() {

}

// Update modifies an origin
func (s *OriginServiceOp) Update() {

}
