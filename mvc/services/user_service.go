package services

import (
	"go-microservices/mvc/domains"
	"go-microservices/mvc/utils"
)

//GetUser will look for the userID into the data source
func GetUser(userID int64) (*domains.User, *utils.ApplicationError) {
	return domains.GetUser(userID)
}
