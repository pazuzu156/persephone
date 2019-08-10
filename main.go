package main

import (
	"persephone/commands"
	"persephone/database"
	"persephone/utils"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/aurora"
)

var migrate = false

// main entry point
func main() {
	if migrate {
		database.Migrate()
	} else {
		config := utils.Config()

		client := aurora.New(&aurora.Options{
			DisgordOptions: &disgord.Config{
				BotToken: config.Token,
				Logger:   disgord.DefaultLogger(false),
			},
		})

		client.Use(aurora.DefaultLogger())
		client.GetPrefix = func(m *disgord.Message) string {
			return config.Prefix
		}

		utils.Check(client.Init())
	}
}

// Initializes all commands (register them here)
func init() {
	ping := commands.InitPing()
	aurora.Use(ping.Register())

	np := commands.InitNowPlaying("np")
	aurora.Use(np.Register())

	help := commands.InitHelp()
	aurora.Use(help.Register())

	login := commands.InitLogin()
	aurora.Use(login.Register())

	logout := commands.InitLogout()
	aurora.Use(logout.Register())
}
