package core

import (
	"github.com/google/uuid"
)

type UrlShortner interface {
	Create() string
}

type UrlShortnerAdapter struct{}

func (u *UrlShortnerAdapter) Create() string {
	code := uuid.NewString() // here we are generating a random uuid code
	return code[24:]         // here we are sending a short code out of complete uuid code
}
