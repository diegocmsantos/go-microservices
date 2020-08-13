package domains

import (
	"fmt"
	"go-microservices/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		1: {
			Id:        1,
			FirstName: "Diego",
			LastName:  "Maia",
			Email:     "diegocmsantos@gmail.com",
		},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

// GetUser get user from the data source
func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {

	log.Println("This function was called.")

	if user, ok := users[userID]; ok {
		return user, nil
	}

	return nil, utils.CreateError(
		fmt.Sprintf("user with id [%d] not found", userID),
		http.StatusNotFound,
		"not_found")
}
