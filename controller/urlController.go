package controller

import (
	"encoding/json"
	"net/http"
	"urlShortner/database"

	"github.com/gorilla/mux"
)

type Request struct {
	Url string `json:"url"`
}
type UrlControllerAdapter struct {
	UrlDbAdapter database.UrlShortnerDbService
}

//Create handler to create short url
func (c *UrlControllerAdapter) Create(w http.ResponseWriter, r *http.Request) {
	req := Request{}
	w.Header().Set("content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	code, err := c.UrlDbAdapter.Save(req.Url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(code)
}

//Get handler fetches original url from short url
func (c *UrlControllerAdapter) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	values := mux.Vars(r)
	code := values["code"]
	url, err := c.UrlDbAdapter.Get(code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}
