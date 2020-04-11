package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/satyy/go-microservice/src/api/domain/repositories"
	"github.com/satyy/go-microservice/src/api/log/zap_logger"
	"github.com/satyy/go-microservice/src/api/services"
	"github.com/satyy/go-microservice/src/api/utils/errors"
	"net/http"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.BadRequestApiError("Invalid Json body")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	zap_logger.Info("Request to create repo", zap_logger.Field("repo_name", request.Name))
	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		zap_logger.Error("Error creating repo",
			zap_logger.Field("repo_name", request.Name),
			zap_logger.Field("status", err.GetStatus()),
			zap_logger.Field("error", err.GetMessage()))
		c.JSON(err.GetStatus(), err)
		return
	}

	zap_logger.Info("Repo Created", zap_logger.Field("repo_name", request.Name))
	c.JSON(http.StatusCreated, result)
}

func CreateRepos(c *gin.Context) {
	var request []repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.BadRequestApiError("Invalid Json body")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	zap_logger.Info("Request to create repo(s)")
	result, err := services.RepositoryService.CreateRepos(request)
	if err != nil {
		zap_logger.Error("Error creating repo(s)",
			zap_logger.Field("status", err.GetStatus()),
			zap_logger.Field("error", err.GetMessage()))
		c.JSON(err.GetStatus(), err)
		return
	}
	zap_logger.Info("Request completed for create repo(s)", zap_logger.Field("status", result.StatusCode))
	c.JSON(result.StatusCode, result)
}
