package controllers

import (
	"encoding/json"
	"go-microservices/mvc/services"
	"go-microservices/mvc/utils"
	"net/http"
	"strconv"
)

// GetUser is controller to get a single User based on its user ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiError := utils.CreateError(
			"userID must be a number",
			http.StatusBadRequest,
			"bad_request")
		jsonValue, _ := json.Marshal(apiError)
		w.WriteHeader(apiError.StatusCode)
		w.Write(jsonValue)
		return
	}

	user, apiError := services.GetUser(userID)
	if apiError != nil {
		jsonValue, _ := json.Marshal(apiError)
		w.WriteHeader(apiError.StatusCode)
		w.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	w.Write([]byte(jsonValue))
}
