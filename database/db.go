package database

import (
	"context"
	"errors"
	"os"
	"urlShortner/core"

	"github.com/go-redis/redis/v8"
)

type DbServiceAdapter struct {
	client      *redis.Client
	urlShortner core.UrlShortner
}

func NewDbServiceAdapter(client *redis.Client, urlShortner core.UrlShortner) *DbServiceAdapter {
	return &DbServiceAdapter{client: client, urlShortner: urlShortner}
}

//Save method takes a url and saves it in database along
func (db *DbServiceAdapter) Save(url string) (string, error) {
	ctx := context.Background()
	url, err := validateUrl(url)
	if err != nil {
		return url, err
	}
	value, err := db.client.Get(ctx, url).Result()
	if err == nil {
		return value, nil
	}
	code := db.urlShortner.Create()

	err = db.client.HSet(ctx, "urls", url, code).Err()
	if err != nil {
		return "", err
	}
	return os.Getenv("HOST_ADDRESS") + "/" + code, nil
}

//Get method will fetch actual associated url which was provided as input
//Here we are providing short url as input and it will fetch original url from database
func (db *DbServiceAdapter) Get(shortUrl string) (string, error) {
	var resultUrl string
	values, err := db.client.HGetAll(context.Background(), "urls").Result()
	if err != nil {
		return "", errors.New("url not found")
	}
	for k, v := range values {
		if v == shortUrl {
			resultUrl = k
			return k, nil
		}
	}

	return resultUrl, nil
}
