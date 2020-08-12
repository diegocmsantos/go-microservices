package domains

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	userID := int64(0)
	user, err := GetUser(userID)

	assert.Nil(t, user, fmt.Sprintf("User with id [%d] is not expected", userID))
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.StatusCode)
}

func TestGetUserUserFound(t *testing.T) {
	userID := int64(1)
	user, err := GetUser(userID)

	assert.NotNil(t, user, fmt.Sprintf("User with id [%d] is expected", userID))
	assert.Nil(t, err)
}
