package v1

import (
	"Ad_Placement_Service/controllers"
	"Ad_Placement_Service/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"strconv"
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
	//now := time.Now()
	gender := c.Query("gender")
	age := c.Query("age")
	country := c.Query("country")
	platform := c.Query("platform")
	rowOffset := c.DefaultQuery("offset", "0")
	rowLimit := c.DefaultQuery("limit", "5")
	offset, offsetErr := strconv.ParseInt(rowOffset, 10, 32)
	if offsetErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "offset must be number")
	}
	limit, limitErr := strconv.ParseInt(rowLimit, 10, 32)
	if limitErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "limit must be number")
	}
	log.Printf("offset:%d, limit:%d", offset, limit)
	log.Printf("gender:%s, age:%s, country:%s, platform:%s \n", gender, age, country, platform)

	c.JSON(http.StatusOK, "res")

}
