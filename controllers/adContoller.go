package controllers

import (
	"Ad_Placement_Service/models"
)

func CreateAd(ad models.Advertisement) error {
	if err := ad.InsertDb(); err != nil {
		return err
	}
	return nil
}
