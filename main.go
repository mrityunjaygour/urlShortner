package main

import (
	"log"
	"net/http"
	"os"
	"urlShortner/controller"
	"urlShortner/core"
	"urlShortner/database"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var rdb *redis.Client
var adapter core.UrlShortnerAdapter

// initiallized database here
func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDRESS"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       1,
	})
}

func main() {
	initializeRoutes()
}

func initializeRoutes() {
	err := godotenv.Load() // this will load env variables
	if err != nil {
		log.Fatalln("env variables could not be loaded")
	}
	r := mux.NewRouter() // here we created mux router's instance
	urlController := controller.UrlControllerAdapter{database.NewDbServiceAdapter(rdb, &adapter)}
	r.HandleFunc("/url", urlController.Create).Methods("POST")
	r.HandleFunc("/{code}", urlController.Get).Methods("GET")

	http.ListenAndServe(os.Getenv("HOST_ADDRESS"), r)
}
