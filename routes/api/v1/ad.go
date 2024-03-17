package v1

import (
	"Ad_Placement_Service/controllers"
	"Ad_Placement_Service/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

func CreateAd(c *gin.Context) {
	var ad models.Advertisement
	if err := c.ShouldBindWith(&ad, binding.JSON); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := controllers.CreateAd(ad); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Create Ad success")
}
