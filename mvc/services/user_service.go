package services

import (
	"go-microservices/mvc/domains"
	"go-microservices/mvc/utils"
)

type userService struct{}

var UserService userService

//GetUser will look for the userID into the data source
func (us *userService) GetUser(userID int64) (*domains.User, *utils.ApplicationError) {
	return domains.UserDao.GetUser(userID)
}
