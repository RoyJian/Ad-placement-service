package controllers

import (
	"Ad_Placement_Service/domain"
	"Ad_Placement_Service/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PlacementController struct {
	PlacementUseCase *usecase.PlacementUseCase
}

func NewPlacementController(PlacementUseCase *usecase.PlacementUseCase) *PlacementController {
	return &PlacementController{PlacementUseCase: PlacementUseCase}
}

func (pc *PlacementController) Query(ctx *gin.Context) {
	var params domain.PlacementRequest
	if err := ctx.BindQuery(&params); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	res, err := pc.PlacementUseCase.Query(ctx, params)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, domain.PlacementResponse{Item: res})
}
