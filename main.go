package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"urlShortner/controller"
	"urlShortner/core"
	"urlShortner/database"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var urlModel *core.Url

// initiallized database here

func main() {
	startRoutes()
}

func init() {
	err := godotenv.Load() // this will load env variables
	if err != nil {
		log.Fatalln("env variables could not be loaded")
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln("error in database connection", err)
	}
	db = client.Database(os.Getenv("DB_NAME"))
	log.Println("database connected successfully")
}

func startRoutes() {
	r := mux.NewRouter() // here we created mux router's instance
	urlController := controller.UrlControllerAdapter{database.NewDbServiceAdapter(db, urlModel)}
	r.HandleFunc("/url", urlController.Create).Methods("POST")
	r.HandleFunc("/{code}", urlController.Get).Methods("GET")

	http.ListenAndServe(os.Getenv("HOST_ADDRESS"), r)
}
