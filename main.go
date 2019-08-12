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
	ping := commands.InitPing()
	aurora.Use(ping.Register())

	np := commands.InitNowPlaying("np")
	aurora.Use(np.Register())

	help := commands.InitHelp("h")
	aurora.Use(help.Register())

	login := commands.InitLogin("li")
	aurora.Use(login.Register())

	logout := commands.InitLogout("lo")
	aurora.Use(logout.Register())

	wk := commands.InitWhoknows("wk")
	aurora.Use(wk.Register())

	bi := commands.InitBandinfo("bi")
	aurora.Use(bi.Register())
}
