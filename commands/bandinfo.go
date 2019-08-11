package commands

import (
	"strings"

	"github.com/pazuzu156/aurora"
	"github.com/pazuzu156/lastfm-go"
)

type Bandinfo struct {
	Command Command
}

func InitBandinfo(aliases ...string) Bandinfo {
	return Bandinfo{Init(
		"bandinfo",
		"Gets information on a band",
		[]UsageItem{
			{
				Command:     "bandinfo",
				Description: "Gets information on the artist you're currently listening to",
			},
			{
				Command:     "bandinfo [artist]",
				Description: "Gets information on a requested artist",
			},
		},
		aliases...,
	)}
}

func (c Bandinfo) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		if len(ctx.Args) > 0 {
			artist := strings.Trim(strings.Join(ctx.Args, ", "), ", ")
			a, err := c.Command.Lastfm.Artist.GetInfo(lastfm.P{"artist": artist})

			if err != nil {
				ctx.Message.RespondString(ctx.Aurora, "Artist could not be found on Last.fm")

				return
			}

			c.displayArtistInfo(ctx, a)
		}
	}

	return c.Command.CommandInterface
}

func (c Bandinfo) displayArtistInfo(context aurora.Context, artist lastfm.ArtistGetInfo) {

}
