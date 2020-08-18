package services

import (
	"go-microservices/src/api/config"
	"go-microservices/src/api/domain/github"
	"go-microservices/src/api/domain/repositories"
	"go-microservices/src/api/providers/github_provider"
	"go-microservices/src/api/utils/errors"
	"strings"
)

type repositoryService struct{}

type repositoryServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService repositoryServiceInterface
)

func init() {
	RepositoryService = &repositoryService{}
}

func (s *repositoryService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("Invalid repository name")
	}

	request := github.CreateRepoRequestBody{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Owner: response.Owner.Login,
		Name:  response.Name,
	}

	return &result, nil
}
