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
	theRouter.HandleFunc("/user", handlers.UserHandler).Methods(http.MethodGet, http.MethodPost)
	theRouter.HandleFunc("/user/{id}", handlers.UserHandler).Methods(http.MethodGet)
	theRouter.HandleFunc("/user/{id}/loan/{loanId}", handlers.LoanHandler).Methods(http.MethodPost)
	theRouter.HandleFunc("/user/{id}/amortization/{loanId}", handlers.AmortizationHandler).Methods(http.MethodGet)

	log.Println("THe API is listining")

	http.ListenAndServe(":8083", theRouter)

}
