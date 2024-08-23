package Groupie_tracker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// func to Get All data From json
func GetArtistsDataStruct() ([]JsonData, error) {
	var artistData []JsonData

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, fmt.Errorf("error fetching data from artist data: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching data from artist data: %d", response.StatusCode)
	}

	err = json.NewDecoder(response.Body).Decode(&artistData)
	if err != nil {
		return nil, fmt.Errorf("error fetching data from artist data: %v", err)
	}
	
	return artistData, nil
}

// / func to fetching data from any struct and return Struct Artist with Id user
func FetchDataRelationFromId(id string) (Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api"
	var artist Artist
	err := GetanyStruct(url + "/artists/" + id, &artist)
	if err != nil {
		return Artist{}, fmt.Errorf("error fetching data from artist data: %w", err)
	}

	if artist.ID == 0 {
		return artist, nil
	}

	var date Date
	errdate := GetanyStruct(url + "/dates/" + id, &date)
	if errdate != nil {
		return Artist{}, fmt.Errorf("error fetching data from artist data: %w", errdate)
	}

	var location Location
	errlocations := GetanyStruct(url + "/locations/" + id, &location)
	if errlocations != nil {
		return Artist{}, fmt.Errorf("error fetching data from locations data: %w", errlocations)
	}
	var relation Relation

	errrelation := GetanyStruct(url + "/relation/" + id, &relation)
	if errrelation != nil {
		return Artist{}, fmt.Errorf("error fetching data from locations data: %w", errrelation)
	}
	artist.Location = location.Location
	artist.Date = date.Date

	var maps []string
	artist.DatesLocations , maps = formatLocations(relation.DatesLocations)
	artist.maps = maps

	return artist, nil
}

// func to UnmarshalJSON from any struct with send url and any
// return error for has any error

func GetanyStruct(url string, result interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data from URL: %w", err)
	}
	defer response.Body.Close()
	// Decode the JSON response into the provided struct
	if err := json.NewDecoder(response.Body).Decode(result); err != nil {
		return fmt.Errorf("error decoding JSON data: %w", err)
	}
	return nil
}

// func To Format String To remove '_' or '-' and Capitaliz text
func formatLocations(locations map[string][]string) (map[string][]string , []string ){
	formatted := make(map[string][]string, len(locations))
	var str []string
	for location, dates := range locations {
		formattedLoc := strings.Title(strings.NewReplacer("-", " ", "_", " ").Replace(location))
		formatted[formattedLoc] = dates
	}
	return formatted , str
}
