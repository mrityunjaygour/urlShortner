package database

import (
	"context"
	"errors"
	"os"
	"time"
	"urlShortner/core"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DbServiceAdapter struct {
	database    *mongo.Database
	urlShortner core.UrlShortner
}

func NewDbServiceAdapter(db *mongo.Database, urlShortner core.UrlShortner) *DbServiceAdapter {
	return &DbServiceAdapter{database: db, urlShortner: urlShortner}
}

//Save method takes a url and saves it in database
func (dba *DbServiceAdapter) Save(url string) (string, error) {
	ctx := context.Background()
	url, err := validateUrl(url)
	if err != nil {
		return url, err
	}
	filter := bson.D{{"value", url}}
	count, err := dba.database.Collection(os.Getenv("COLLECTION_NAME")).CountDocuments(ctx, filter)
	if err == nil {
		if count != 0 {
			var urlModel core.Url
			result := dba.database.Collection(os.Getenv("COLLECTION_NAME")).FindOne(ctx, filter)
			err = result.Decode(&urlModel)
			if err != nil {
				return "", err
			}
			return os.Getenv("HOST_ADDRESS") + "/" + urlModel.ShortUrl, nil
		}
	}
	urlModel := dba.urlShortner.Create()
	urlModel.Value = url
	urlModel.Created = time.Now().Round(time.Second).String()
	result, err := dba.database.Collection(os.Getenv("COLLECTION_NAME")).InsertOne(ctx, *urlModel)
	if err != nil {
		return "", err
	}

	if result.InsertedID == nil {
		return "", errors.New("could not insert data")
	}

	return os.Getenv("HOST_ADDRESS") + "/" + urlModel.ShortUrl, nil
}

//Get method will fetch actual associated url which was provided as input
//Here we are providing short url as input and it will fetch original url from database
func (dba *DbServiceAdapter) Get(shortUrl string) (string, error) {
	var urlModel core.Url
	filter := bson.D{{"shorturl", shortUrl}}
	update := bson.D{{"$set", bson.D{{"updated", time.Now().Round(time.Second).String()}}}}
	result := dba.database.Collection(os.Getenv("COLLECTION_NAME")).FindOneAndUpdate(context.Background(), filter, update)
	err := result.Decode(&urlModel)
	if err != nil {
		return "", err
	}

	// updateReslt, err := dba.database.Collection(os.Getenv("COLLECTION_NAME")).UpdateOne(context.Background(), filter, update)
	// if err != nil {
	// 	return "", errors.New("couldn't updated data")
	// }
	// if updateReslt.MatchedCount == 0 {
	// 	return "", errors.New("no match found for update")
	// }
	return urlModel.Value, nil
}
