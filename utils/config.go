package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Configuration is the base json object.
type Configuration struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
	Lastfm Lastfm `json:"lastfm"`
}

// Lastfm is the lastfm json object.
type Lastfm struct {
	APIKey string
	Secret string
}

// Config retrieves the app's configuration form config.json.
func Config() Configuration {
	file, err := os.Open("config.json")
	Check(err)
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	Check(err)

	var config Configuration
	err = json.Unmarshal(contents, &config)
	Check(err)

	return config
}
