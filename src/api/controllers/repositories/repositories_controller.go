package repositories

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-microservices/src/api/domain/repositories"
	"go-microservices/src/api/services"
	"go-microservices/src/api/utils/errors"
	"net/http"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError(fmt.Sprintf("invalid json request body: %q", err.Error()))
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}
