package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Configuration is the base json object.
type Configuration struct {
	Token    string   `json:"token"`
	Prefix   string   `json:"prefix"`
	Lastfm   Lastfm   `json:"lastfm"`
	Database Database `json:"database"`
}

// Lastfm is the lastfm json object.
type Lastfm struct {
	APIKey string `json:"apikey"`
	Secret string `json:"secret"`
}

// Database is the database json object.
type Database struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
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
