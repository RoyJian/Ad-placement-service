package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHealth(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
