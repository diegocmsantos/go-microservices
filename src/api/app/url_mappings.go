package app

import (
	"go-microservices/src/api/controllers/health"
	"go-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/health", health.Health)
	router.POST("/repositories", repositories.CreateRepo)
}
