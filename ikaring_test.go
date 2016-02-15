package ikaring

import (
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
