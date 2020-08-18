package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	healthy = "healthy"
)

func Health(c *gin.Context) {
	c.String(http.StatusOK, healthy)
}
