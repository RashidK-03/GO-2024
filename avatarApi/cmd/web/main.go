package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting API server...")
	router := mux.NewRouter()

	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/characters", GetCharacters).Methods("GET")
	router.HandleFunc("/characters/{name}", GetCharacterByName).Methods("GET")
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}
