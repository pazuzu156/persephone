package fm

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"persephone/lib"
)

// MaArtist is an artist struct for metal-archives artist
type MaArtist struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// GetMaArtistList returns a list of artists defined in
// artists.json for metal-archives
func GetMaArtistList() []MaArtist {
	file, err := os.Open("artists.json")
	lib.Check(err)
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	lib.Check(err)

	var artists []MaArtist
	err = json.Unmarshal(contents, &artists)
	lib.Check(err)

	return artists
}

// GetMaArtist returns an artist listed in
// artists.json for metal-archives
func GetMaArtist(artist string) MaArtist {
	list := GetMaArtistList()

	for _, a := range list {
		if a.Name == artist {
			return a
		}
	}

	return MaArtist{}
}
