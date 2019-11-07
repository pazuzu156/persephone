package lib

import (
	"fmt"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
)

var (
	_client *atlas.Atlas
)

// RegisterEvents registers any client events.
func RegisterEvents(client *atlas.Atlas) {
	_client = client

	client.On(disgord.EvtGuildMemberRemove, func(s disgord.Session, evt *disgord.GuildMemberRemove) {
		user := GetUser(evt.User)
		ur, cr := user.Delete()
		guild := GetServer(evt.GuildID)

		if cr && ur {
			s.CreateMessage(UInt64ToSnowflake(guild.GuildID), &disgord.CreateMessageParams{
				Content: fmt.Sprintf("User %s and there crowns were removed", user.Username),
			})
		} else if cr && !ur {
			s.CreateMessage(UInt64ToSnowflake(guild.GuildID), &disgord.CreateMessageParams{
				Content: fmt.Sprintf("User %s was not removed, but their crowns were", user.Username),
			})
		} else if !cr && ur {
			s.CreateMessage(UInt64ToSnowflake(guild.GuildID), &disgord.CreateMessageParams{
				Content: fmt.Sprintf("User %s was removed, but their crowns were not", user.Username),
			})
		} else {
			s.CreateMessage(UInt64ToSnowflake(guild.GuildID), &disgord.CreateMessageParams{
				Content: fmt.Sprintf("User %s and crowns were not removed", user.Username),
			})
		}
	})

	client.On(disgord.EvtReady, func(s disgord.Session, evt *disgord.Ready) {
		for _, guild := range evt.Guilds {
			dbg := GetServer(guild.ID)

			if dbg.GuildID != SnowflakeToUInt64(guild.ID) {
				db, _ := OpenDB()
				now := time.Now()
				var s = Servers{
					GuildID:      SnowflakeToUInt64(guild.ID),
					LogChannelID: 0,
					ElevatedRole: 0,
					Time: Time{
						CreatedAt: &now,
						UpdatedAt: &now,
					},
				}
				_, err := db.Insert(&s)
				Check(err)
			}
		}

		for _, i := range []string{
			"Bot Online",
			fmt.Sprintf("Name: %s", evt.User.Username),
			fmt.Sprintf("ID: %s", evt.User.ID.String()),
		} {
			client.Logger.Info(i)
		}
		updatePresence()
	})

	// Handles the starboard
	// client.On(disgord.EvtMessageReactionAdd, func(s disgord.Session, evt *disgord.MessageReactionAdd) {
	// 	if evt.PartialEmoji.Name == "⭐" {
	// 		currentChannel, _ := s.GetChannel(evt.ChannelID)
	// 		message, _ := s.GetMessage(currentChannel.ID, evt.MessageID)
	// 		guild, _ := s.GetGuild(currentChannel.GuildID)

	// 		// only work if we've reached the activation count trigger limit
	// 		if len(message.Reactions) >= config.Starboard.ActivationCount {
	// 			for _, channel := range guild.Channels {
	// 				if channel.Name == config.Starboard.Channel {
	// 					_, t := AddEmbedFooter(message)
	// 					client.CreateMessage(channel.ID, &disgord.CreateMessageParams{
	// 						Embed: &disgord.Embed{
	// 							Title:       "Content",
	// 							URL:         GenerateMessageURL(guild.ID, message),
	// 							Description: message.Content,
	// 							Fields: []*disgord.EmbedField{
	// 								{
	// 									Name:   "Author",
	// 									Value:  message.Author.Mention(),
	// 									Inline: true,
	// 								},
	// 								{
	// 									Name:   "Channel",
	// 									Value:  currentChannel.Name,
	// 									Inline: true,
	// 								},
	// 							}, Timestamp: t,
	// 						},
	// 					})
	// 				}
	// 			}
	// 		}
	// 	}
	// })
}

func updatePresence() {
	for true {
		status := &disgord.UpdateStatusPayload{
			Game: &disgord.Activity{
				Name: "Metal",
				Type: 2,
			},
			Status: disgord.StatusOnline,
		}

		_client.UpdateStatus(status)
		time.Sleep(25 * time.Second) // run every 25 seconds
	}
}
