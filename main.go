package main

import (
	"persephone/commands"
	"persephone/utils"

	"github.com/andersfylling/disgord"
	"github.com/polaron/aurora"
)

var migrate = false

// main entry point
func main() {
	if migrate {
		utils.Migrate()
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

	help := commands.InitHelp("help")
	aurora.Use(help.Register())

	login := commands.InitLogin("login")
	aurora.Use(login.Register())
}
