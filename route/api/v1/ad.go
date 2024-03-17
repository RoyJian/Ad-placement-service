package v1

import (
	"Ad_Placement_Service/controller"
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
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	if err := controller.CreateAd(ad); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "ad create success")
}
