package main

import (
	"persephone/commands"
	"persephone/database"
	"persephone/lib"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/aurora"
)

var migrate = false

// main entry point
func main() {
	if migrate {
		database.Migrate()
	} else {
		config := lib.Config()

		client := aurora.New(&aurora.Options{
			DisgordOptions: &disgord.Config{
				BotToken: config.Token,
				Logger:   disgord.DefaultLogger(false),
			},
			OwnerID: config.BotOwner,
		})

		client.Use(aurora.DefaultLogger())
		client.GetPrefix = func(m *disgord.Message) string {
			return config.Prefix
		}

		lib.Check(client.Init())
	}
}

// Initializes all commands (register them here)
func init() {
	aurora.Use(commands.InitBand().Register())
	aurora.Use(commands.InitHelp().Register())
	aurora.Use(commands.InitLogin().Register())
	aurora.Use(commands.InitLogout().Register())
	aurora.Use(commands.InitNowPlaying().Register())
	aurora.Use(commands.InitPing().Register())
	aurora.Use(commands.InitWhoknows().Register())
	aurora.Use(commands.InitYoutube().Register())
}
