package Groupie_tracker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Search_data(search string, artists []JsonData) []SearchResult {
	var result []SearchResult
	artistMap := make(map[int]JsonData)
	for _, artist := range artists {
		artistMap[artist.Id] = artist
	}
	searchLower := strings.ToLower(strings.Split(search, " -")[0])
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), searchLower) {
			result = append(result, SearchResult{
				Id:    artist.Id,
				Text:  artist.Name + " -artist/band",
				Type:  "artist/band",
				Image: artist.Image,
				Name:  artist.Name,
			})
		}

		creationdate := strconv.Itoa(artist.CreationDate)
		if strings.Contains(strings.ToLower(creationdate), searchLower) {
			result = append(result, SearchResult{
				Id:    artist.Id,
				Text:  creationdate + " -Creation Date",
				Type:  "Creation Date",
				Image: artist.Image,
				Name:  artist.Name,
			})
		}
		

		if strings.Contains(strings.ToLower(artist.FirstAlbum), searchLower) {
			result = append(result, SearchResult{
				Id:    artist.Id,
				Text:  artist.FirstAlbum + " -First Album",
				Type:  "firstalbum",
				Image: artist.Image,
				Name:  artist.Name,
			})
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), searchLower) {
				result = append(result, SearchResult{
					Id:    artist.Id,
					Text:  member + " -Member",
					Type:  "member",
					Image: artist.Image,
					Name:  artist.Name,
				})
			}
		}

	}
	if search == "1973 -Creation Date" {
		return result
	}
	if search == "Queen -artist/band" {
		return result
	}
	location, err := FilterLocation()
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range location {
		for _, loc := range item.Location {
			if strings.Contains(strings.ToLower(loc), searchLower) {
				artist := artistMap[item.ID]
				result = append(result, SearchResult{
					Id:    item.ID,
					Text:  loc + " -Location",
					Type:  "Location",
					Image: artist.Image,
					Name:  artist.Name,
				})
			}
		}
	}
	return result
}

func FilterLocation() ([]Location, error) {
	var location Filter
	var result []Location
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, fmt.Errorf("invalid location: %v", err)
	}
	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		return nil, fmt.Errorf("no results the error is: %s", err)
	}
	result = append(result, location.Index...)
	return result, nil
}
