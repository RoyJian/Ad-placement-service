package controllers

import (
	"Ad_Placement_Service/domain"
	"Ad_Placement_Service/usecase"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type AdminController struct {
	AdminUseCase *usecase.AdminUseCase
}

func NewAdminController(AdminUseCase *usecase.AdminUseCase) *AdminController {
	return &AdminController{AdminUseCase: AdminUseCase}
}

func (ad *AdminController) Create(ctx *gin.Context) {
	advertisement := domain.Advertisement{
		Conditions: domain.Condition{
			AgeStart: 1,
			AgeEnd:   100,
		},
	}
	admin := domain.Admin{Advertisement: &advertisement}
	if err := ctx.ShouldBindWith(&advertisement, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := ad.AdminUseCase.Create(ctx, &admin); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "Create Advertisement success")
}
