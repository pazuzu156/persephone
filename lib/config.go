package lib

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Configuration is the base json object.
type Configuration struct {
	Token    string `json:"token"`
	BotOwner string `json:"bot_owner"`
	BotID    string `json:"bot_id"`
	Prefix   string `json:"prefix"`
	// Starboard string `json:"starboard"`
	Starboard struct {
		Channel         string `json:"channel"`
		ActivationCount int    `json:"activation_count"`
	} `json:"starboard"`
	Lastfm struct {
		APIKey string `json:"apikey"`
		Secret string `json:"secret"`
	} `json:"lastfm"`
	Database struct {
		Hostname string `json:"hostname"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
	} `json:"database"`
	YouTube struct {
		APIKey string `json:"apikey"`
	} `json:"youtube"`
	Website struct {
		AppURL string `json:"app_url"`
		APIURL string `json:"api_url"`
	} `json:"website"`
}

// Config retrieves the app's configuration form config.json.
func Config() Configuration {
	file, err := os.Open(LocGet("config.json"))
	Check(err)
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	Check(err)

	var config Configuration
	err = json.Unmarshal(contents, &config)
	Check(err)

	return config
}
