package grabjson

import (
	"encoding/json"
	"groupie-tracker/server/model"
	"io"
	"net/http"
)

var (
	artists          model.AllArtist
	relation         model.Relations
	jsonlinkArtists  = "https://groupietrackers.herokuapp.com/api/artists"
	jsonlinkRelation = "https://groupietrackers.herokuapp.com/api/relation"
)

func GetData() (model.AllArtist, error) {
	if err := getJson(jsonlinkArtists, &artists.Artist); err != nil {
		return artists, err
	}

	if err := getJson(jsonlinkRelation, &relation); err != nil {
		return artists, err
	}
	for i, w := range relation.Location {
		artists.Artist[i].DatesLocations = w.DatesLocations
	}
	return artists, nil
}

func getJson(url string, object interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bb, &object); err != nil {
		return err
	}
	return nil
}
