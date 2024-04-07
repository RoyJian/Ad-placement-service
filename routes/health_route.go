package routes

import (
	web "Ad_Placement_Service/bootstrap/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthRoute struct {
	handler *web.Gin
}

func NewHealthRoute(handler *web.Gin) *HealthRoute {
	return &HealthRoute{handler: handler}
}

func (h *HealthRoute) Setup() {
	h.handler.Gin.GET("/health", GetHealth)
}

func GetHealth(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
