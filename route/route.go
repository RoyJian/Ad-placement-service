package route

import "Ad_Placement_Service/route/api/v1"
import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/health", v1.GetHealth)
		apiV1.POST("/ad", v1.CreateAd)

	}
	return router
}
