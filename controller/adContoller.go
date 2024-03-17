package controller

import (
	"Ad_Placement_Service/models"
	"log"
)

func CreateAd(ad models.Advertisement) error {
	log.Println(ad)
	if err := ad.InsertDb(); err != nil {
		return err
	}
	return nil
}
