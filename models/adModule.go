package models

import (
	"Ad_Placement_Service/service/mongodb"
	"context"
	"log"
	"time"
)

type Advertisement struct {
	Title      string     `json:"title" binding:"required" bson:"title"`
	StartAt    *time.Time `json:"startAt" binding:"required" bson:"startAt"`
	EndAt      *time.Time `json:"endAt" binding:"required,gtfield=StartAt" bson:"endAt"`
	Conditions Condition  `json:"conditions" bson:"conditions"`
}
type Condition struct {
	AgeStart int      `json:"ageStart" binding:"omitempty,gte=1,lte=100" bson:"ageStart"`
	AgeEnd   int      `json:"ageEnd"   binding:"omitempty,gte=1,gtfield=AgeStart,lte=100" bson:"ageEnd"`
	Country  []string `json:"country"  binding:"omitempty,dive,iso3166_1_alpha2" bson:"country"`
	Platform []string `json:"platform" binding:"omitempty,dive,oneof=android ios web" bson:"platform"`
	Gender   string   `json:"gender"   binding:"omitempty,oneof=M F" bson:"gender"`
}

func (ad *Advertisement) InsertDb() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := mongodb.GetCollection("advertisements")
	res, err := collection.InsertOne(ctx, ad)
	log.Println(res)
	if err != nil {
		return err
	}
	return nil
}
