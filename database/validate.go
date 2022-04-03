package database

import (
	"errors"
	"net/url"
	"strings"
)

//validate fundtion will validate input url string
func validateUrl(urlString string) (string, error) {

	if urlString == "" {
		return urlString, errors.New("input url can not be empty")
	}
	if strings.Contains(urlString, "http://.") {
		return urlString, errors.New("invalid input url")
	}
	if !strings.Contains(urlString, "http://") {
		urlString = "http://" + urlString
	}
	u, err := url.ParseRequestURI(urlString)
	if err != nil {
		return urlString, err
	}
	if !strings.Contains(u.Host, ".") {
		return urlString, errors.New("invalid input url")
	}
	return urlString, nil
}
