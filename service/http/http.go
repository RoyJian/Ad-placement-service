package http

import (
	"Ad_Placement_Service/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func Init() error {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	route.SetRouter(router)
	port := fmt.Sprintf(":%s", os.Getenv("GIN_PORT"))
	if err := router.Run(port); err != nil {
		return err
	}
	return nil
}
