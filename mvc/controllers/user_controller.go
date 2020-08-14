package controllers

import (
	"fmt"
	"go-microservices/mvc/services"
	"go-microservices/mvc/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUser is controller to get a single User based on its user ID
func GetUser(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		apiError := utils.CreateError(
			fmt.Sprintf("%s must be a number", userIDStr),
			http.StatusBadRequest,
			"bad_request")

		utils.RespondError(c, apiError)
		return
	}

	user, apiError := services.UserService.GetUser(userID)
	if apiError != nil {
		utils.RespondError(c, apiError)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}
