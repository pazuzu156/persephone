package commands

import (
	"fmt"
	"persephone/utils"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/aurora"
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
func (c Recent) Register() *aurora.Command {
	c.CommandInterface.Run = func(ctx aurora.Context) {
		recent, _ := utils.GetRecentTracks(ctx.Message.Author, c.Lastfm, "4")
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
				Value: utils.JoinString(tracksslc, "\n"),
			})

			ctx.Aurora.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title: "Recent Tracks",
					Thumbnail: &disgord.EmbedThumbnail{
						URL: utils.GenAvatarURL(ctx.Message.Author),
					},
					Fields: tracks,
					Color:  utils.RandomColor(),
				},
			})
		}
	}

	return c.CommandInterface
}
