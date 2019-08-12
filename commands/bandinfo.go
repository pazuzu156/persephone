package commands

import (
	"persephone/utils"
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
		var artist lastfm.ArtistGetInfo
		var err error
		if len(ctx.Args) > 0 {
			artistName := strings.Trim(strings.Join(ctx.Args, ", "), ", ")
			artist, err = c.getArtistInfo(artistName)
		} else {
			var track utils.Track
			track, err = utils.GetNowPlayingTrack(ctx.Message.Author, c.Command.Lastfm)
			artist, err = c.getArtistInfo(track.Artist.Name)
		}

		if err != nil {
			ctx.Message.RespondString(ctx.Aurora, err.Error())

			return
		}

		c.displayArtistInfo(ctx, artist)
	}

	return c.Command.CommandInterface
}

func (c Bandinfo) displayArtistInfo(context aurora.Context, artist lastfm.ArtistGetInfo) {
	albums := c.Command.Lastfm.Artist.GetTopAlbums()
}

func (c Bandinfo) getArtistInfo(artist string) (lastfm.ArtistGetInfo, error) {
	return c.Command.Lastfm.Artist.GetInfo(lastfm.P{"artist": artist})
}
