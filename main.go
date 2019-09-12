package main

import (
	"persephone/commands"
	"persephone/database"
	"persephone/lib"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/aurora"
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
		client := aurora.New(&aurora.Options{
			DisgordOptions: &disgord.Config{
				BotToken: config.Token,
				Logger:   disgord.DefaultLogger(false),
				Presence: &disgord.UpdateStatusCommand{
					Game: &disgord.Activity{
						Name: "Metal",
						Type: 2,
					},
					Status: disgord.StatusOnline,
				},
			},
			OwnerID: config.BotOwner,
		})

		client.Use(aurora.DefaultLogger())
		client.GetPrefix = func(m *disgord.Message) string {
			return config.Prefix
		}

		// Handles the starboard
		client.On(disgord.EvtMessageReactionAdd, func(s disgord.Session, evt *disgord.MessageReactionAdd) {
			if evt.PartialEmoji.Name == "⭐" {
				currentChannel, _ := s.GetChannel(evt.ChannelID)
				message, _ := s.GetMessage(currentChannel.ID, evt.MessageID)
				guild, _ := s.GetGuild(currentChannel.GuildID)

				// only work if we've reached the activation count trigger limit
				if len(message.Reactions) >= config.Starboard.ActivationCount {
					for _, channel := range guild.Channels {
						if channel.Name == config.Starboard.Channel {
							f, t := lib.AddEmbedFooter(message)
							client.CreateMessage(channel.ID, &disgord.CreateMessageParams{
								Embed: &disgord.Embed{
									Title:       "Content",
									URL:         lib.GenerateMessageURL(guild.ID, message),
									Description: message.Content,
									Fields: []*disgord.EmbedField{
										{
											Name:   "Author",
											Value:  message.Author.Mention(),
											Inline: true,
										},
										{
											Name:   "Channe",
											Value:  currentChannel.Name,
											Inline: true,
										},
									}, Timestamp: t,
								},
							})
						}
					}
				}
			}
		})

		lib.Check(client.Init())
	}
}

// Initializes all commands (register them here)
func init() {
	aurora.Use(commands.InitAbout().Register())
	aurora.Use(commands.InitBandinfo().Register())
	// aurora.Use(commands.InitBand().Register())
	aurora.Use(commands.InitCrowns().Register())
	aurora.Use(commands.InitHelp().Register())
	aurora.Use(commands.InitLogout().Register())
	aurora.Use(commands.InitNowPlaying().Register())
	aurora.Use(commands.InitRecent().Register())
	aurora.Use(commands.InitPing().Register())
	aurora.Use(commands.InitPlays().Register())
	aurora.Use(commands.InitRegister().Register())
	aurora.Use(commands.InitWhoknows().Register())
	aurora.Use(commands.InitYoutube().Register())
}
