package models

import (
	"time"
)

type Advertisement struct {
	Title      string     `json:"title" binding:"required"`
	StartAt    *time.Time `json:"startAt" binding:"required"`
	EndAt      *time.Time `json:"endAt" binding:"required"`
	Conditions Condition  `json:"conditions"`
}
type Condition struct {
	AgeStart int      `json:"ageStart" binding:"omitempty,gte=1,lte=100"`
	AgeEnd   int      `json:"ageEnd"   binding:"omitempty,gte=1,gtfield=AgeStart,lte=100"`
	Country  []string `json:"country"  binding:"omitempty,dive,iso3166_1_alpha2"`
	Platform []string `json:"platform" binding:"omitempty,dive,oneof=android ios web"`
	Gender   string   `json:"gender"   binding:"omitempty,oneof=M F"`
}

func (ad *Advertisement) InsertDb() error {

	return nil
}
