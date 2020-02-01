package lib

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// MaArtist is an artist struct for metal-archives artist
type MaArtist struct {
	Name string `yaml:"name"`
	ID   int    `yaml:"id"`
}

// GetMaArtistList returns a list of artists defined in
// artists.json for metal-archives
func GetMaArtistList() []MaArtist {
	file, err := os.Open(LocGet("artists.yml"))
	Check(err)
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	Check(err)

	var artists []MaArtist
	err = yaml.Unmarshal(contents, &artists)
	Check(err)

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
