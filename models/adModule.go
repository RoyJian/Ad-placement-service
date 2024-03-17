package models

import (
	"Ad_Placement_Service/enums"
)

type Advertisement struct {
	Title      string    `json:"title" binding:"required"`
	StartAt    string    `json:"startAt"`
	EndAt      string    `json:"endAt"`
	Conditions Condition `json:"conditions"`
}
type Condition struct {
	AgeStart int              `json:"ageStart"`
	AgeEnd   int              `json:"ageEnd"`
	Country  []string         `json:"country"`
	Platform []enums.Platform `json:"platform"`
	Gender   enums.Gender     `json:"gender"`
}
