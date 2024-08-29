package Groupie_tracker

import (
	"strconv"
	"strings"
)

func Search_data(search string, artists []JsonData) []JsonData {
	var result []JsonData
	for _, item := range artists {
		dates := strconv.Itoa(item.CreationDate)
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(item.FirstAlbum), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(item.Locations), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(dates), strings.ToLower(search)) {
			result = append(result, item)
		}
	}
	return result
}

// strings.Contains(strings.ToLower(item.Members[]), strings.ToLower(search))
