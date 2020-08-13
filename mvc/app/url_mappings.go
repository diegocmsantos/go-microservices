package app

import (
	"go-microservices/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
