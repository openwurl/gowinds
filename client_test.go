package gowinds

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var coreAPIStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var resp string
	switch r.RequestURI {
	case "/api/v1/accounts/default/hosts/brokenhash":
		resp = `
		{
			"Error": "Does not exist"
		}
		`
	case "/api/v1/accounts/default/hosts/expectedhash":
		resp = `
		{
			"HashCode": "expectedhash",
			"Name": "Host01",
			"Error": nil
		}
		`
	}
	w.Write([]byte(resp))
}))

func TestNewClient(t *testing.T) {
	c, err := NewClient("")
	if err == nil {
		t.Fatal("expected new client instantiation to fail, instead got client: ", c)
	}

	testToken := "TestToken"
	c, err = NewClient(testToken)
	if err != nil {
		t.Fatal("failed to instantiate new client: ", err)
	}

	if c.AuthorizationHeaderToken != testToken {
		t.Fatalf("expected c.AuthorizationHeaderToken to be %s, got %s", testToken, c.AuthorizationHeaderToken)
	}

	var testLogger logger
	if reflect.TypeOf(c.logger) != reflect.TypeOf(testLogger) {
		t.Fatalf("expected c.logger to be %v, got %v", reflect.TypeOf(testLogger), c.logger)
	}

	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	c.SetLogger(logger)

	if !reflect.DeepEqual(logger, c.logger) {
		t.Fatalf("expected c.logger to be %v, got %v", logger, c.logger)
	}

	c.SetBaseURL(coreAPIStub.URL)

	if c.BaseURL.String() != coreAPIStub.URL {
		t.Fatalf("Expected c.BaseURL to be %v, got %v", coreAPIStub.URL, c.BaseURL)
	}

}

//func TestNewRequest(t *testing.T) {
//
//}

//func TestDo(t *testing.T) {
//
//}

func TestDoRequest(t *testing.T) {
	testToken := "TestToken"
	c, err := NewClient(testToken)
	if err != nil {
		t.Fatal("failed to instantiate new client: ", err)
	}

	// purposefully bad URL
	err = c.SetBaseURL("postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require")
	if err == nil {
		t.Fatalf("expected c.SetBaseURL to error, got: %v - %v", c.BaseURL, err)
	}

	err = c.SetBaseURL(coreAPIStub.URL)
	if err != nil {
		t.Fatalf("expected c.SetBaseURL to not error, got: %v", err)
	}

	reqOpt := RequestOptions{
		AccountHash: "default",
	}
	var hostResponse interface{}
	method := "GET"
	path := fmt.Sprintf("%s/default/hosts/brokenhash", reqOpt.createURL())
	resp, err := c.DoRequest(method, path, nil, hostResponse)
	if err != nil {
		t.Fatalf("expected err to be nil, got %v", err)
	}
	t.Log(resp)
	// This will be expanded once we have things to put it into and properly analyze
}
