package main

import (
	"fmt"
	"log"
	"net/http"

	Groupie_tracker "groupie_tracker/handlers"
)

func main() {
	port := ":8091"
	http.HandleFunc("/", Groupie_tracker.GetDataFromJson)
	http.HandleFunc("/Artist/{id}", Groupie_tracker.HandlerShowRelation)
	http.HandleFunc("/geoMap", Groupie_tracker.GeoMap)
	http.HandleFunc("/styles/", Groupie_tracker.HandleStyle)
	http.HandleFunc("/search", Groupie_tracker.HandleSearch)
	fmt.Printf("http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
