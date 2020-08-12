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
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, err.Message, "user with id [0] not found")
	assert.EqualValues(t, err.Code, "not_found")
}

func TestGetUserUserFound(t *testing.T) {
	userID := int64(1)
	user, err := GetUser(userID)

	assert.NotNil(t, user, fmt.Sprintf("User with id [%d] is expected", userID))
	assert.Nil(t, err)

	assert.EqualValues(t, user.Id, 1)
	assert.EqualValues(t, user.FirstName, "Diego")
	assert.EqualValues(t, user.LastName, "Maia")
	assert.EqualValues(t, user.Email, "diegocmsantos@gmail.com")
}
