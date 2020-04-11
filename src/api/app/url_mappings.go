package app

import (
	"github.com/satyy/go-microservice/src/api/controllers/health"
	"github.com/satyy/go-microservice/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/health", health.Health)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
