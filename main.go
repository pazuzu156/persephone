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

		client.Use(atlas.DefaultLogger())
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
							_, t := lib.AddEmbedFooter(message)
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
											Name:   "Channel",
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
	atlas.Use(commands.InitAbout().Register())
	atlas.Use(commands.InitBandinfo().Register())
	// atlas.Use(commands.InitBand().Register())
	atlas.Use(commands.InitCrowns().Register())
	atlas.Use(commands.InitHelp().Register())
	atlas.Use(commands.InitUnregister().Register())
	atlas.Use(commands.InitNowPlaying().Register())
	atlas.Use(commands.InitRecent().Register())
	// atlas.Use(commands.InitPing().Register())
	atlas.Use(commands.InitPlays().Register())
	atlas.Use(commands.InitRegister().Register())
	atlas.Use(commands.InitTaste().Register())
	atlas.Use(commands.InitTopCrowns().Register())
	atlas.Use(commands.InitWhoknows().Register())
	atlas.Use(commands.InitYoutube().Register())

	// atlas.Use(commands.InitNewtaste().Register())
}
