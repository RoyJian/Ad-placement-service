package routes

import (
	"Ad_Placement_Service/bootstrap/http"
	"Ad_Placement_Service/controllers"
)

type AdvertisementRoute struct {
	adminController     *controllers.AdminController
	placementController *controllers.PlacementController
	handler             *http.Gin
}

func NewAdvertisementRoute(
	adminController *controllers.AdminController,
	placementController *controllers.PlacementController,
	handler *http.Gin,
) *AdvertisementRoute {
	return &AdvertisementRoute{
		adminController:     adminController,
		placementController: placementController,
		handler:             handler,
	}
}

func (ar AdvertisementRoute) Setup() {
	api := ar.handler.Gin.Group("/api/v1")
	{
		api.POST("/ad", ar.adminController.Create)
		api.GET("/ad", ar.placementController.Query)
	}

}
