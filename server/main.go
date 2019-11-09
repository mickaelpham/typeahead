package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	words := []string{"mango", "manchester united", "mandalorian"}
	err := json.NewEncoder(w).Encode(words)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/search", searchHandler)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
