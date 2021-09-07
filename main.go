package main

import (
	"github.com/gorilla/handlers"
	"log"
	"muelltermine/api"
	"net/http"
)

func main() {
	log.Println("Hello muelltermine!")

	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
