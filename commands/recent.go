package commands

import (
	"fmt"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
	"github.com/pazuzu156/persephone/fm"
	"github.com/pazuzu156/persephone/lib"
)

// Recent command.
type Recent struct{ Command }

// InitRecent initializes the recent command.
func InitRecent() Recent {
	return Recent{Init(&CommandItem{
		Name:        "recent",
		Description: "Shows a list of recent tracks",
	})}
}

// Register registers and runs the recent command.
func (c Recent) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		// restrict users who haven't updated lastfm login
		if c.restrict(ctx) {
			ctx.Message.Reply(ctx.Context, ctx.Atlas, "You must re-login to Persephone before gaining access to this command. Please see #announcements for more info")

			return
		}

		recent, _ := fm.GetRecentTracks(ctx.Message.Author, c.Lastfm, "4")
		var tracks []*disgord.EmbedField

		if len(recent) > 0 {
			var tracksslc []string
			for _, track := range recent {
				if track.NowPlaying == "true" {
					tracks = append(tracks, &disgord.EmbedField{
						Name:  "Currently Playing",
						Value: fmt.Sprintf("%s - %s", track.Artist.Name, track.Name),
					})
				} else {
					tracksslc = append(tracksslc, fmt.Sprintf("%s - %s", track.Artist.Name, track.Name))
				}
			}

			tracks = append(tracks, &disgord.EmbedField{
				Name:  "Previous Tracks",
				Value: lib.JoinString(tracksslc, "\n"),
			})

			footer, time := c.embedFooter(ctx)
			ctx.Atlas.CreateMessage(ctx.Context, ctx.Message.ChannelID, &disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title: "Recent Tracks",
					Thumbnail: &disgord.EmbedThumbnail{
						URL: lib.GenAvatarURL(ctx.Message.Author),
					},
					Fields: tracks,
					Color:  lib.RandomColor(),
					Footer: footer, Timestamp: time,
				},
			})
		}
	}

	return c.CommandInterface
}
