package v1

import (
	"Ad_Placement_Service/controller"
	"Ad_Placement_Service/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAd(c *gin.Context) {
	var ad models.Advertisement
	if err := c.BindJSON(&ad); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	if err := controller.CreateAd(ad); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "ad create success")
}
