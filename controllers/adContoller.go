package controllers

import (
	"Ad_Placement_Service/models"
)

type AdQueryRes struct {
	item []models.Advertisement
}

func CreateAd(ad models.Advertisement) error {
	if err := ad.InsertDb(); err != nil {
		return err
	}
	return nil
}

func QueryAd(query models.AdQueryParams) (AdQueryRes, error) {
	var res AdQueryRes

	return res, nil
}
