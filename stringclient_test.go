package gowinds

import "testing"

func TestNew(t *testing.T) {
	c := New("XXXX", "APPNAME")

	if c.AuthorizationHeaderToken != "XXXX" {
		t.Fatalf("Expected XXXX as AuthorizationHeaderToken, got %v", c.AuthorizationHeaderToken)
	}

	if c.ApplicationID != "APPNAME" {
		t.Fatalf("Expected APPNAME as ApplicationID, got %v", c.ApplicationID)
	}
}
