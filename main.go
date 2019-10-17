package main

import (
	"persephone/commands"
	"persephone/database"
	"persephone/lib"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
)

var (
	migrate = false // make true to migrate databases
	config  = lib.Config()
)

// main entry point
func main() {
	if migrate {
		database.Migrate()
	} else {
		client := atlas.New(&atlas.Options{
			DisgordOptions: &disgord.Config{
				BotToken: config.Token,
				// Logger:   disgord.DefaultLogger(true), // uncomment for disgord logging
			},
			OwnerID: config.BotOwner,
		})

		client.Use(atlas.DefaultLogger())
		client.GetPrefix = func(m *disgord.Message) string {
			return config.Prefix
		}

		lib.RegisterEvents(client)

		lib.Check(client.Init())
	}
}

// Initializes all commands (register them here)
func init() {
	atlas.Use(commands.InitAbout().Register())
	atlas.Use(commands.InitBandinfo().Register())
	// atlas.Use(commands.InitBand().Register()
	// atlas.Use(commands.InitChart().Register())
	atlas.Use(commands.InitCrownBoard().Register())
	atlas.Use(commands.InitCrowns().Register())
	atlas.Use(commands.InitHelp().Register())
	atlas.Use(commands.InitUnregister().Register())
	atlas.Use(commands.InitNowPlaying().Register())
	atlas.Use(commands.InitRecent().Register())
	// atlas.Use(commands.InitPing().Register())
	atlas.Use(commands.InitPlays().Register())
	atlas.Use(commands.InitRegister().Register())
	atlas.Use(commands.InitTaste().Register())
	atlas.Use(commands.InitWhoknows().Register())
	atlas.Use(commands.InitYoutube().Register())

	// atlas.Use(commands.InitNewtaste())

	// Bot Owner commands.
	atlas.Use(commands.InitDeleteUser().Register())
}
