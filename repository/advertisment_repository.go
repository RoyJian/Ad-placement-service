package repository

import (
	"Ad_Placement_Service/bootstrap/cache"
	"Ad_Placement_Service/bootstrap/db"
	"Ad_Placement_Service/domain"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/sync/singleflight"
	"log"
	"reflect"
	"time"
)

type AdvertisementRepository struct {
	database   db.Db
	collection string
	cache      cache.Cache
	engine     singleflight.Group
}

func NewAdvertisementRepository(db *db.MongoDb, cache *cache.Redis) *AdvertisementRepository {
	var engine singleflight.Group
	return &AdvertisementRepository{
		database:   db,
		collection: domain.CollectionAdvertisement,
		cache:      cache,
		engine:     engine,
	}
}

func (ar *AdvertisementRepository) Create(ctx context.Context, ad *domain.Advertisement) error {
	collection := ar.database.GetCollection(ar.collection)
	_, err := collection.InsertOne(ctx, ad)
	if err != nil {
		return err
	}
	return nil
}
func (ar *AdvertisementRepository) Query(ctx context.Context, params domain.PlacementRequest) ([]domain.Advertisement, error) {
	// cache
	key := fmt.Sprintf("%+v", params)
	res, err := ar.queryFromCache(ctx, key)
	if err == nil {
		log.Println("Cache hit")
		return res, nil
	}
	// Cache miss, uses singleflight to avoid Hotspot Invalid
	v, err, _ := ar.engine.Do(key, func() (interface{}, error) {
		res, err := ar.queryFromDB(ctx, params)
		// Write to cache
		defer ar.writeToCache(ctx, key, res) // write query result to cache
		return res, err
	})
	if err != nil {
		return nil, err
	}
	return v.([]domain.Advertisement), nil
}

func (ar *AdvertisementRepository) queryFromCache(ctx context.Context, key string) ([]domain.Advertisement, error) {
	var res []domain.Advertisement
	var bytes []byte
	err := ar.cache.Get(ctx, key).Scan(&bytes)
	if err != nil {
		log.Println("Cache miss")
		return nil, err
	}
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		log.Println("Parse json error", err)
		return nil, err
	}
	return res, nil
}

func (ar *AdvertisementRepository) queryFromDB(ctx context.Context, params domain.PlacementRequest) ([]domain.Advertisement, error) {
	var res []domain.Advertisement
	var agg []bson.D
	collection := ar.database.GetCollection(ar.collection)
	// Filter expired ads
	agg = append(agg, bson.D{
		{"$match", bson.D{
			{"startAt", bson.D{{"$lte", time.Now().UTC()}}},
			{"endAt", bson.D{{"$gte", time.Now().UTC()}}},
		}},
	})
	// Filter Gender
	if !reflect.ValueOf(params.Gender).IsZero() {
		agg = append(agg, bson.D{{"$match", bson.D{{"$or", bson.A{
			bson.D{{"conditions.gender", params.Gender}},
			bson.D{{"conditions.gender", bson.D{{"$exists", false}}}},
		}}}}})
	}
	// Filter Age
	if !reflect.ValueOf(params.Age).IsZero() {
		agg = append(agg, bson.D{
			{"$match", bson.D{
				{"conditions.ageStart", bson.D{{"$lte", params.Age}}},
				{"conditions.ageEnd", bson.D{{"$gte", params.Age}}},
			}},
		})
	}
	// Filter Country
	if !reflect.ValueOf(params.Country).IsZero() {
		agg = append(agg, bson.D{
			{"$match", bson.D{{"$or", bson.A{
				bson.D{{"conditions.country", bson.D{{"$all", bson.A{params.Country}}}}},
				bson.D{{"conditions.country", bson.D{{"$exists", false}}}},
			}}}},
		})
	}
	// Filter Platform
	if !reflect.ValueOf(params.Platform).IsZero() {
		agg = append(agg, bson.D{
			{"$match", bson.D{{"$or", bson.A{
				bson.D{{"conditions.platform", bson.D{{"$all", bson.A{params.Platform}}}}},
				bson.D{{"conditions.platform", bson.D{{"$exists", false}}}},
			}}}},
		})
	}

	agg = append(agg, bson.D{{"$skip", params.Offset}})        // set query offset
	agg = append(agg, bson.D{{"$limit", params.Limit}})        // set query limit
	agg = append(agg, bson.D{{"$sort", bson.D{{"endAt", 1}}}}) // sort query result

	cursor, err := collection.Aggregate(ctx, agg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err := cursor.All(context.TODO(), &res); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return res, nil
}

func (ar *AdvertisementRepository) writeToCache(ctx context.Context, key string, value []domain.Advertisement) {
	encodeValue, err := json.Marshal(value)
	if err != nil {
		log.Println(err)

	}
	if err := ar.cache.Set(ctx, key, encodeValue); err != nil {
		log.Println(err)
	}
}
