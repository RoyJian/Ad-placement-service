package models

import (
	"Ad_Placement_Service/service/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Advertisement struct {
	Id         primitive.ObjectID `json:"id" binding:"omitempty" bson:"_id omitempty"`
	Title      string             `json:"title" binding:"required" bson:"title"`
	StartAt    *time.Time         `json:"startAt" binding:"required" bson:"startAt"`
	EndAt      *time.Time         `json:"endAt" binding:"required,gtfield=StartAt" bson:"endAt"`
	Conditions Condition          `json:"conditions" bson:"conditions"`
}
type Condition struct {
	AgeStart int      `json:"ageStart" binding:"omitempty,gte=1,lte=100" bson:"ageStart"`
	AgeEnd   int      `json:"ageEnd"   binding:"omitempty,gte=1,gtfield=AgeStart,lte=100" bson:"ageEnd"`
	Country  []string `json:"country"  binding:"omitempty,dive,iso3166_1_alpha2" bson:"country"`
	Platform []string `json:"platform" binding:"omitempty,dive,oneof=android ios web" bson:"platform"`
	Gender   string   `json:"gender"   binding:"omitempty,oneof=M F FM" bson:"gender"`
}
type QueryParams struct {
}

func (ad *Advertisement) InsertDb() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := mongodb.GetCollection("advertisements")
	_, err := collection.InsertOne(ctx, ad)
	if err != nil {
		return err
	}
	return nil
}

func (condition *Condition) Init() {
	condition.AgeStart = 1
	condition.AgeEnd = 100
	condition.Gender = "FM"
}
