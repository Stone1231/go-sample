package jsonex

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{
		Title:  "Casablanca",
		Year:   1942,
		Color:  false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{
		Title:  "Cool Hand Luke",
		Year:   1967,
		Color:  true,
		Actors: []string{"Paul Newman"}},
	{
		Title:  "Bullitt",
		Year:   1968,
		Color:  true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func Test_main(t *testing.T) {
	{
		//!+Marshal
		data, err := json.Marshal(movies)
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
		//!-Marshal
	}

	{
		//!+MarshalIndent
		data, err := json.MarshalIndent(movies, "", "    ")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		log.Printf("%s\n", data)
		//!-MarshalIndent

		//!+Unmarshal
		var _movies []Movie
		if err := json.Unmarshal(data, &_movies); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		log.Println(_movies)

		var titles []struct{ Title string }
		if err := json.Unmarshal(data, &titles); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		log.Println(titles)
		//!-Unmarshal
	}
}
