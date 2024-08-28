package Groupie_tracker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var location Location

func GeoMap(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	indexStr := r.URL.Query().Get("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(location.Location) {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		return
	}
	result, err := geocodeAddress(location.Location[index])
	if err != nil {
		http.Error(w, "Error geocoding address", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func geocodeAddress(address string) (string, error) {
	var result string
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?format=json&q=%s", address)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: %s", resp.Status)
	}

	var results []GeocodeResponse
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return "", fmt.Errorf("no results the error is: %s", err)
	}
	if len(results) == 0 {
		return "", fmt.Errorf("no results found for address: %s", address)
	}
	result = fmt.Sprintf("https://www.google.com/maps/place/%v,%v/@%v,%v,3000m/", results[0].Lat, results[0].Lon, results[0].Lat, results[0].Lon)

	return result, nil
}
