package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/satyy/go-microservice/src/api/client/rest_client"
	"github.com/satyy/go-microservice/src/api/domain/github"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}
func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GitHubErrorReponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	response, err := rest_client.Post(urlCreateRepo, request, headers)

	if err != nil {
		log.Println(fmt.Sprintf("GetError when trying to create a new Repo in github %s", err.Error()))
		return nil, &github.GitHubErrorReponse{
			StatusCode:      http.StatusInternalServerError,
			Message:          err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GitHubErrorReponse{StatusCode: http.StatusInternalServerError, Message: "Invalid Response Body"}
	}
	// *** Close response body!!! ***
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GitHubErrorReponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GitHubErrorReponse{StatusCode: http.StatusInternalServerError, Message: "Invalid JSON Response Body"}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var successResponse github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &successResponse); err != nil {
		log.Println(fmt.Sprintf("GetError when trying to unmarshal create repo success response: %s ", err.Error()))
		return nil, &github.GitHubErrorReponse{StatusCode: http.StatusInternalServerError, Message: "GetError when trying to unmarshal create repo success response"}
	}
	return &successResponse, nil
}
