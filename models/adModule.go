package models

import (
	"Ad_Placement_Service/service/cache"
	"Ad_Placement_Service/service/mongodb"
	"context"
	"encoding/json"
	"fmt"
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
	Limit    int    `form:"limit,default=5" binding:"gte=1,lte=100"`
	Offset   int    `form:"offset,default=0"`
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

func (adQueryParams *AdQueryParams) generateKey() string {
	s := fmt.Sprintf("%+v", adQueryParams)
	return s
}

func (adQueryParams *AdQueryParams) Query() ([]Advertisement, error) {
	var res []Advertisement
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// cache
	adQueryParams.queryFromCache(ctx, &res)
	if !reflect.ValueOf(res).IsZero() {
		log.Println("Cache hit")
		return res, nil
	}
	// cache miss, query DB
	err := adQueryParams.queryFromDB(ctx, &res)
	if err != nil {
		return nil, err
	}
	defer adQueryParams.writeToCache(res)
	return res, nil
}

func (adQueryParams *AdQueryParams) queryFromCache(ctx context.Context, res *[]Advertisement) {
	var bytes []byte
	key := adQueryParams.generateKey()
	err := cache.Get(ctx, key).Scan(&bytes)
	if err != nil {
		log.Println("Cache miss")
		return
	}
	err = json.Unmarshal(bytes, res)
	if err != nil {
		log.Println("Parse json error", err)
	}
}

func (adQueryParams *AdQueryParams) queryFromDB(ctx context.Context, res *[]Advertisement) error {
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

	agg = append(agg, bson.D{{"$skip", adQueryParams.Offset}}) // set query offset
	agg = append(agg, bson.D{{"$limit", adQueryParams.Limit}}) // set query limit
	agg = append(agg, bson.D{{"$sort", bson.D{{"endAt", 1}}}}) // sort query result

	cursor, err := collection.Aggregate(ctx, agg)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if err := cursor.All(context.TODO(), res); err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func (adQueryParams *AdQueryParams) writeToCache(value []Advertisement) {
	encodeValue, err := json.Marshal(value)
	if err != nil {
		log.Fatal(err)
	}
	if err := cache.Set(adQueryParams.generateKey(), encodeValue); err != nil {
		log.Fatal(err)
	}
}
