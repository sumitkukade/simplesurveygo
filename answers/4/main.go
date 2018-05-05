package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 100 * time.Second}

// function to get json from URL
func getJsonFromUrl(target interface{}) error {
	url := "https://raw.githubusercontent.com/prust/wikipedia-movie-data/master/movies.json"
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {

}
