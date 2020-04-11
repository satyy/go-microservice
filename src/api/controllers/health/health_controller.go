package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	healthCheck = "Application is Up!"
)

func Health(c *gin.Context) {
	c.String(http.StatusOK, healthCheck)
}