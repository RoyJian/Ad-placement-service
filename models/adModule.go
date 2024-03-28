package models

import (
	"Ad_Placement_Service/service/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"reflect"
	"time"
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

type AdQueryParams struct {
	Gender   string `form:"gender" binding:"omitempty,oneof=M F"`
	Age      int    `form:"age" binding:"omitempty,gte=1,lte=100"`
	Country  string `form:"country" binding:"omitempty,iso3166_1_alpha2"`
	Platform string `form:"platform" binding:"omitempty,oneof=android ios web"`
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
}

func (adQueryParams *AdQueryParams) Query() ([]Advertisement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := mongodb.GetCollection("advertisements")
	var agg = mongo.Pipeline{}

	// Filter expired ads
	agg = append(agg, bson.D{
		{"$match", bson.D{
			{"startAt", bson.D{{"$lte", time.Now().UTC()}}},
			{"endAt", bson.D{{"$gte", time.Now().UTC()}}},
		}},
	})
	// Filter Gender
	if !reflect.ValueOf(adQueryParams.Gender).IsZero() {
		agg = append(agg, bson.D{{"$match", bson.D{{"$or", bson.A{
			bson.D{{"conditions.gender", adQueryParams.Gender}},
			bson.D{{"conditions.gender", bson.D{{"$exists", false}}}},
		}}}}})
	}
	// Filter Age
	if !reflect.ValueOf(adQueryParams.Age).IsZero() {
		agg = append(agg, bson.D{
			{"$match", bson.D{
				{"conditions.ageStart", bson.D{{"$lte", adQueryParams.Age}}},
				{"conditions.ageEnd", bson.D{{"$gte", adQueryParams.Age}}},
			}},
		})
	}
	// Filter Country
	if !reflect.ValueOf(adQueryParams.Country).IsZero() {
		agg = append(agg, bson.D{
			{"$match", bson.D{{"$or", bson.A{
				bson.D{{"conditions.country", bson.D{{"$all", bson.A{adQueryParams.Country}}}}},
				bson.D{{"conditions.country", bson.D{{"$exists", false}}}},
			}}}},
		})
	}

	// Filter Platform
	if !reflect.ValueOf(adQueryParams.Platform).IsZero() {
		agg = append(agg, bson.D{
			{"$match", bson.D{{"$or", bson.A{
				bson.D{{"conditions.platform", bson.D{{"$all", bson.A{adQueryParams.Platform}}}}},
				bson.D{{"conditions.platform", bson.D{{"$exists", false}}}},
			}}}},
		})
	}

	cursor, err := collection.Aggregate(ctx, agg)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var res []Advertisement
	if err := cursor.All(context.TODO(), &res); err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return res, nil

}
