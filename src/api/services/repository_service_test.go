package services

import (
	"github.com/stretchr/testify/assert"
	"go-microservices/src/api/clients/restclient"
	"go-microservices/src/api/domain/repositories"
	"go-microservices/src/api/providers/github_provider"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidInputName(t *testing.T) {
	// Arrange
	request := repositories.CreateRepoRequest{}

	// Act
	result, err := RepositoryService.CreateRepo(request)

	// Assert
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "Invalid repository name", err.Message())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	// Arrange
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        github_provider.UrlCreateRepo,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires Authentication", "documentation_url": "any string"}`)),
		},
	})
	request := repositories.CreateRepoRequest{
		Name: "testing",
	}

	// Act
	result, err := RepositoryService.CreateRepo(request)

	// Assert
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires Authentication", err.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	// Arrange
	restclient.FlushMockups()
	request := repositories.CreateRepoRequest{
		Name:        "testing",
		Description: "test description",
	}
	restclient.AddMockup(restclient.Mock{
		Url:        github_provider.UrlCreateRepo,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing"}`)),
		},
	})

	// Act
	result, err := RepositoryService.CreateRepo(request)

	// Assert
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "testing", result.Name)
}
