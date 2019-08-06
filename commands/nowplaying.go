package commands

import (
	"fmt"

	"github.com/polaron/aurora"
	"github.com/shkh/lastfm-go/lastfm"
)

// Nowplaying command.
type Nowplaying struct {
	Command Command
}

// InitNowPlaying initializes the nowplaying command.
func InitNowPlaying(aliases ...string) Nowplaying {
	return Nowplaying{Init(
		"nowplaying",
		"Shows what you're currently listening to",
		aliases...,
	)}
}

// Register registers and runs the nowplaying command.
func (c Nowplaying) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		res, _ := c.Command.Lastfm.User.GetRecentTracks(lastfm.P{
			"user":  "Pazuzu156",
			"limit": "1",
		})

		track := res.Tracks[0]

		if track.NowPlaying == "true" {
			ctx.Message.RespondString(ctx.Aurora, fmt.Sprintf("Now Playing: %s - %s", track.Artist.Name, track.Name))
		} else {
			ctx.Message.RespondString(ctx.Aurora, "You're currently not listening to anything")
		}
	}

	return c.Command.CommandInterface
}
