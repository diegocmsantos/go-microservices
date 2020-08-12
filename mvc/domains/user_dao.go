package domains

import (
	"fmt"
	"go-microservices/mvc/utils"
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
)

// GetUser get user from the data source
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user, ok := users[userID]; ok {
		return user, nil
	}

	return nil, utils.CreateError(
		fmt.Sprintf("user with id [%d] not found", userID),
		http.StatusNotFound,
		"not_found")
}
