package main

import (
	"log"
	"net/http"
)

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}

	// load data from config
	itemsLimit := config.Max
	packSizes := config.PackSizes

	http.HandleFunc("/pack", func(w http.ResponseWriter, r *http.Request) {
		packHandler(w, r, packSizes, itemsLimit)
	})

	// start the http server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
