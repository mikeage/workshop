package main

import (
	"log"
	"net/http"
	"workshop/handlers"

	"github.com/gorilla/mux"
)

func main() {
	theRouter := mux.NewRouter()
	theRouter.HandleFunc("/route", handlers.BasicHandler).Methods(http.MethodGet)
	log.Println("The API is listening")
	http.ListenAndServe(":8083", theRouter)
}
