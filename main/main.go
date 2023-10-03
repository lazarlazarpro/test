package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sort"
	"strconv"
)

import (
	"io/ioutil"
)

type Config struct {
	PackSizes []int `json:"packSizes"`
	Max       int   `json:"Max"`
}

var config Config

func LoadConfig(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	if config.PackSizes == nil {
		return errors.New("PackSizes is nil after loading config")
	}

	return nil
}

func roundNumberOfItems(numberOfItems int) int {
	roundedNumberOfItems := numberOfItems
	for _, packSize := range config.PackSizes {
		if numberOfItems >= packSize {
			numberOfItems %= packSize
		}
	}
	if numberOfItems > 0 {
		roundedNumberOfItems += config.PackSizes[len(config.PackSizes)-1] - numberOfItems
	}

	return roundedNumberOfItems
}

func getPackSizes(numberOfItems int) *map[int]int {
	numberOfItems = roundNumberOfItems(numberOfItems)
	packCounts := make(map[int]int)

	for _, packSize := range config.PackSizes {
		if numberOfItems >= packSize {
			packCounts[packSize] = numberOfItems / packSize
			numberOfItems %= packSize
		}
	}

	return &packCounts
}

func packHandler(w http.ResponseWriter, r *http.Request) {
	itemsStr := r.URL.Query().Get("items")
	items, err := strconv.Atoi(itemsStr)
	if err != nil {
		http.Error(w, "Invalid item count", http.StatusBadRequest)
		return
	}

	if items <= 0 || items >= config.Max {
		http.Error(w, "Item count must be greater than 0 and less than "+strconv.Itoa(config.Max), http.StatusBadRequest)
		return
	}

	result := getPackSizes(items)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	err := LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(config.PackSizes)))
	http.HandleFunc("/pack", packHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
