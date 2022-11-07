package grabjson

import "groupie-tracker/server/model"

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
