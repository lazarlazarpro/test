package main

import (
	"challange/pack"
	"encoding/json"
	"net/http"
	"strconv"
)

// packHandler handles the packing request
func packHandler(w http.ResponseWriter, r *http.Request, packSizes []int, itemsLimit int) {
	// handle CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	itemsStr := r.URL.Query().Get("items")
	items, err := strconv.Atoi(itemsStr)
	if err != nil {
		http.Error(w, "Invalid item count", http.StatusBadRequest)
		return
	}

	if items <= 0 || items >= itemsLimit {
		http.Error(w, "Item count must be greater than 0 and less than "+strconv.Itoa(itemsLimit), http.StatusBadRequest)
		return
	}

	p := pack.NewPacker(packSizes)
	result := p.GetPackSizes(items)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
