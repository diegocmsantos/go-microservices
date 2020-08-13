package services

import (
	"go-microservices/mvc/domains"
	"go-microservices/mvc/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock     userDaoMockStruct
	getUserMockFunc func(int64) (*domains.User, *utils.ApplicationError)
)

func init() {
	domains.UserDao = &userDaoMockStruct{}
}

type userDaoMockStruct struct{}

func (u *userDaoMockStruct) GetUser(userID int64) (*domains.User, *utils.ApplicationError) {
	return getUserMockFunc(userID)
}

func TestGetUserNotFound(t *testing.T) {
	getUserMockFunc = func(userID int64) (*domains.User, *utils.ApplicationError) {
		return nil, utils.CreateError("User not found", http.StatusNotFound, "not_found")
	}

	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestGetUserFound(t *testing.T) {
	getUserMockFunc = func(userID int64) (*domains.User, *utils.ApplicationError) {
		return &domains.User{
			Id:        1,
			FirstName: "John",
			LastName:  "Nobody",
			Email:     "nobody@nobodycorp.com",
		}, nil
	}

	user, err := UserService.GetUser(0)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "John", user.FirstName)
	assert.EqualValues(t, "Nobody", user.LastName)
	assert.EqualValues(t, "nobody@nobodycorp.com", user.Email)
}
