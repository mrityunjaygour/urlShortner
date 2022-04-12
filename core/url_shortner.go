package core

import (
	"github.com/google/uuid"
)

type Url struct {
	Value    string `json:"value" bson:"value"`
	ShortUrl string `json:"shorturl" bson:"shorturl"`
	Created  string `json:"created" bson:"created"`
	Updated  string `json:"updated" bson:"updated"`
}

type UrlShortner interface {
	Create() *Url
}

//type UrlShortnerAdapter struct{}

func (u *Url) Create() *Url {
	var url Url
	code := uuid.NewString()
	url.ShortUrl = code[24:]
	return &url
}
