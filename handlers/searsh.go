package Groupie_tracker

import (
	"strings"
)

func Search_data(search string, artists []JsonData) []SearchResult {
	var result []SearchResult
	var artisData Artist
	searchLower := strings.ToLower(search)
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), searchLower) {
			result = append(result, SearchResult{
				Id:   artist.Id,
				Text: artist.Name + " - artist/band",
				Type: "artist/band",
			})
		}

		if strings.Contains(strings.ToLower(artist.FirstAlbum), searchLower) {
			result = append(result, SearchResult{
				Id:   artist.Id,
				Text: artist.FirstAlbum + " - firstalbum",
				Type: "firstalbum",
			})
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), searchLower) {
				result = append(result, SearchResult{
					Id:   artist.Id,
					Text: member + " - member",
					Type: "member",
				})
			}
		}
		for _, location := range artisData.Location {
			if strings.Contains(strings.ToLower(location), searchLower) {
				result = append(result, SearchResult{
					Id:   artist.Id,
					Text: location + " - location",
					Type: "location",
				})
			}
		}
	}
	return result
}

// strings.Contains(strings.ToLower(item.Members[]), strings.ToLower(search))
