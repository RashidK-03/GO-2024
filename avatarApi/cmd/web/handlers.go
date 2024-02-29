package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome!")
}

func GetCharacters(w http.ResponseWriter, r *http.Request) {
	//characters := api.Characters

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(api.Characters)
	// jsonResponse, err := json.Marshal(characters)
	// if err != nil {
	// 	return
	// }

	//w.Write(jsonResponse)
}

func GetCharacterByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	name := strings.ToLower(params["name"])

	for _, character := range api.Characters {
		if strings.ToLower(character.Name) == name {
			json.NewEncoder(w).Encode(character)
			return
		}
	}

	http.Error(w, "Wrong name", http.StatusNotFound)
}
