package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionAdvertisement = "advertisements"
)

type Advertisement struct {
	Id         primitive.ObjectID `json:"id" binding:"omitempty" bson:"_id,omitempty"`
	Title      string             `json:"title" binding:"required" bson:"title"`
	StartAt    *time.Time         `json:"startAt" binding:"required" bson:"startAt"`
	EndAt      *time.Time         `json:"endAt" binding:"required,gtfield=StartAt" bson:"endAt"`
	Conditions Condition          `json:"conditions" bson:"conditions"`
}

type Condition struct {
	AgeStart int      `json:"ageStart,omitempty" binding:"omitempty,gte=1,lte=100" bson:"ageStart,omitempty"`
	AgeEnd   int      `json:"ageEnd,omitempty"   binding:"omitempty,gte=1,gtfield=AgeStart,lte=100" bson:"ageEnd,omitempty"`
	Country  []string `json:"country,omitempty"  binding:"omitempty,dive,iso3166_1_alpha2" bson:"country,omitempty"`
	Platform []string `json:"platform,omitempty" binding:"omitempty,dive,oneof=android ios web" bson:"platform,omitempty"`
	Gender   string   `json:"gender,omitempty"   binding:"omitempty,oneof=M F" bson:"gender,omitempty"`
}

type AdvertisementRepository interface {
	Create(ctx context.Context, ad Advertisement) error
	Query(ctx context.Context, params PlacementRequest) ([]Advertisement, error)
}
