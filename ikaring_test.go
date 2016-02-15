package ikaring

import (
	"testing"
)

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
