package lib

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Configuration is the base yaml object.
type Configuration struct {
	Token        string `yaml:"token"`
	BotOwner     string `yaml:"bot_owner"`
	BotID        string `yaml:"bot_id"`
	GuildID      string `yaml:"guild_id"`
	LogChannelID string `yaml:"log_channel_id"`
	ElevatedRole string `yaml:"elevated_role"`
	Prefix       string `yaml:"prefix"`
	// Starboard string `json:"starboard"`
	Starboard struct {
		Channel         string `yaml:"channel"`
		ActivationCount int    `yaml:"activation_count"`
	} `yaml:"starboard"`
	Lastfm struct {
		APIKey string `yaml:"apikey"`
		Secret string `yaml:"secret"`
	} `yaml:"lastfm"`
	Database struct {
		Hostname string `yaml:"hostname"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	YouTube struct {
		APIKey string `yaml:"apikey"`
	} `yaml:"youtube"`
	Website struct {
		AppURL string `yaml:"app_url"`
		APIURL string `yaml:"api_url"`
	} `yaml:"website"`
}

// Config retrieves the app's configuration form config.json.
func Config() Configuration {
	file, err := os.Open(LocGet("config.yml"))
	Check(err)
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	Check(err)

	var config Configuration
	err = yaml.Unmarshal(contents, &config)
	Check(err)

	return config
}
