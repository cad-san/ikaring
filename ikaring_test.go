package ikaring

import (
	"net/http"
	"testing"
)

func TestIkaClientCreate(t *testing.T) {
	c, err := CreateClient()
	if err != nil {
		t.Errorf("ikaring.CreateClient() has error %v", err)
	}
	if c.Authorized() == true {
		t.Errorf("ikaring.Authorized() should be false")
	}
}

func TestIkaClientSetSession(t *testing.T) {
	c, _ := CreateClient()
	c.SetSession("dummy_session")
	if c.Authorized() != true {
		t.Errorf("ikaring.Authorized() should be true")
	}
}

func TestDecodeJSONError(t *testing.T) {
	jsonStr := `{"error":"error details"}`
	err := checkJSONError([]byte(jsonStr))
	if err == nil || err.Error() != "error details" {
		t.Errorf("decoded error info is invalid")
	}
}

func TestDecodeJSONErrorEmpty(t *testing.T) {
	jsonStr := `{}`
	err := checkJSONError([]byte(jsonStr))
	if err != nil {
		t.Errorf("decoded error should empty")
	}
}

func TestGetOarthQuery(t *testing.T) {
	splatoonOauthURL := "https://splatoon.nintendo.net/users/auth/nintendo"

	resp, err := http.DefaultClient.Get(splatoonOauthURL)
	if err != nil {
		t.Errorf("GET request is failed. URL : %s", splatoonOauthURL)
	}

	query, err := getOauthQuery(resp, "name", "password")
	if err != nil {
		t.Errorf("getOauthQuery() has error :%v", err)
	}
	if len(query) != 10 {
		t.Errorf("query should have 10 entities")
	}
	if _, ok := query["client_id"]; !ok {
		t.Errorf("query should have client_id")
	}
	if _, ok := query["state"]; !ok {
		t.Errorf("query should have state")
	}
}
