package main

import (
	"go-rest-api/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting API server...")
	router := mux.NewRouter()

	router.HandleFunc("/health-check", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/characters", handlers.GetCharacters).Methods("GET")
	router.HandleFunc("/characters/{name}", handlers.GetCharacterByName).Methods("GET")
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)

}
