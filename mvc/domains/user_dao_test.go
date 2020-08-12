package domains

import (
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	userID := int64(0)
	user, err := GetUser(userID)
	if user != nil {
		t.Error("User not expected")
	}

	if err == nil {
		t.Errorf("Expecting an error when passing a invalid userID like %d", userID)
	}

	if err.StatusCode != http.StatusNotFound {
		t.Errorf("Status Code error must be: %d", err.StatusCode)
	}
}
