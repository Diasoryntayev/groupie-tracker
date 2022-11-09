package grabjson

import (
	"groupie-tracker/server/model"
	"strconv"
	"strings"
)

func SearchArtist(searchWord string) (*model.AllArtist, error) {
	var result []model.Artist
	artists, err := GetData()
	if err != nil {
		return nil, err
	}

	for _, artist := range artists.Artist {
		if myContains(artist.Name, searchWord) {
			result = append(result, artist)
			continue
		} else if strings.Contains(strconv.Itoa(artist.CreationDate), searchWord) {
			result = append(result, artist)
			continue
		} else if strings.Contains(artist.FirstAlbum, searchWord) {
			result = append(result, artist)
			continue
		} else {
			for _, member := range artist.Members {
				if strings.Contains(member, searchWord) {
					result = append(result, artist)
					break
				}
			}
		}
		for location := range artist.DatesLocations {
			if myContains(location, searchWord) {
				result = append(result, artist)
				break
			}
		}
		for _, dates := range artist.DatesLocations {
			for _, date := range dates {
				if strings.Contains(date, searchWord) {
					result = append(result, artist)
					break
				}
			}
			break
		}
	}
	allArtists := &model.AllArtist{FoundArtist: result}
	return allArtists, nil
}

func myContains(s, sub string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(sub))
}
