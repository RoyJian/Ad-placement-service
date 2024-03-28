package v1

import (
	"Ad_Placement_Service/controllers"
	"Ad_Placement_Service/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func CreateAd(c *gin.Context) {
	var ad models.Advertisement
	ad.Conditions.Init()
	if err := c.ShouldBindWith(&ad, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := controllers.CreateAd(ad); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Create Ad success")
}

func GetPlacementAd(c *gin.Context) {
	var params models.AdQueryParams
	if err := c.BindQuery(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	res, queryErr := controllers.QueryAd(params)
	if queryErr != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, queryErr.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"items": res,
	})
}
