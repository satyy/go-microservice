package services

import (
	"github.com/satyy/go-microservice/src/api/config"
	"github.com/satyy/go-microservice/src/api/domain/github"
	"github.com/satyy/go-microservice/src/api/domain/repositories"
	"github.com/satyy/go-microservice/src/api/providers/github_provider"
	"github.com/satyy/go-microservice/src/api/utils/errors"
	"net/http"
	"sync"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToke(), request)
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

func (s *reposService) CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	inputChannel := make(chan repositories.CreateRepositoriesResult)
	outputChannel := make(chan repositories.CreateReposResponse)
	buffer :=  make(chan bool, config.GetMaxRoutinesPerRequest())

	defer close(outputChannel)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(request))

	go s.handleRepoResult(&waitGroup, buffer, inputChannel, outputChannel)

	for _, currentRequest := range request {
		buffer <- true
		go s.createRepoConcurrent(currentRequest, inputChannel)
	}

	waitGroup.Wait()
	close(inputChannel)

	result := <-outputChannel

	successCreations := 0
	for _, currentResponse := range result.Results {
		if currentResponse.Response != nil {
			successCreations++
		}
	}

	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.GetStatus()
	} else if successCreations == len(request) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil
}

func (s *reposService) handleRepoResult(waitGroup *sync.WaitGroup, buffer chan bool, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse

	for result := range input {
		results.Results = append(results.Results, result)
		waitGroup.Done()
		<-buffer
	}
	output <- results
}

func (s *reposService) createRepoConcurrent(request repositories.CreateRepoRequest, inputChan chan repositories.CreateRepositoriesResult) {
	if err := request.Validate(); err != nil {
		inputChan <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	result, err := s.CreateRepo(request)

	if err != nil {
		inputChan <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	inputChan <- repositories.CreateRepositoriesResult{Response: result}
}
