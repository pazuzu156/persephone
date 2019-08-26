package commands

import (
	"fmt"
	"persephone/fm"
	"persephone/lib"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/aurora"
	"github.com/pazuzu156/lastfm-go"
)

// Bandinfo command.
type Bandinfo struct{ Command }

// InitBandinfo initializes the bandinfo command.
func InitBandinfo() Bandinfo {
	return Bandinfo{Init(&CommandItem{
		Name:        "bandinfo",
		Description: "Gets information on the current playing band",
		Aliases:     []string{"bi"},
		Usage:       "bandinfo Darkthrone",
		Parameters:  []Parameter{},
	})}
}

// Register registers and runs the bandinfo command.
func (c Bandinfo) Register() *aurora.Command {
	c.CommandInterface.Run = func(ctx aurora.Context) {
		// ctx.Message.Reply(ctx.Aurora, "Hello, Bandinfo!")
		if len(ctx.Args) > 0 {
			artistName := lib.JoinString(ctx.Args, " ")
			artist, err := c.Lastfm.Artist.GetInfo(lastfm.P{"artist": artistName})

			if err != nil {
				ctx.Message.Reply(ctx.Aurora, "That artist could not be found")

				return
			}

			c.displayBandInfo(ctx, artist)
		} else {
			np, err := fm.GetNowPlayingTrack(ctx.Message.Author, c.Lastfm)

			if err != nil {
				ctx.Message.Reply(ctx.Aurora, "You're not currently listening to anything")

				return
			}

			artist, _ := c.Lastfm.Artist.GetInfo(lastfm.P{"artist": np.Artist.Name})
			c.displayBandInfo(ctx, artist)
		}
	}

	return c.CommandInterface
}

func (c Bandinfo) displayBandInfo(ctx aurora.Context, artist lastfm.ArtistGetInfo) {
	fmt.Println(artist.Bio.Links[0].URL)
	ctx.Aurora.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
		Embed: &disgord.Embed{
			Title: artist.Name,
			URL:   artist.URL,
			Color: lib.RandomColor(),
			Thumbnail: &disgord.EmbedThumbnail{
				URL: lib.GetArtistImageURL(artist),
			},
			Fields: []*disgord.EmbedField{
				{
					Name:  "Summary",
					Value: lib.ShortStr(artist.Bio.Content, 850, fmt.Sprintf(" [Read More...](%s)", artist.Bio.Links[0].URL)),
				},
				{
					Name:  "Tags",
					Value: lib.JoinString(c.tags(artist), ", "),
				},
				{
					Name:  "Similar Artists",
					Value: lib.JoinString(c.similar(artist), ", "),
				},
				{
					Name:   "Total Listeners",
					Value:  lib.HumanNumber(artist.Stats.Listeners),
					Inline: true,
				},
				{
					Name:   "Total Play Count",
					Value:  lib.HumanNumber(artist.Stats.Plays),
					Inline: true,
				},
			},
		},
	})
}

func (c Bandinfo) tags(artist lastfm.ArtistGetInfo) (tags []string) {
	for _, tag := range artist.Tags {
		tags = append(tags, fmt.Sprintf("[%s](%s)", tag.Name, tag.URL))
	}

	return
}

func (c Bandinfo) similar(artist lastfm.ArtistGetInfo) (similar []string) {
	for _, sim := range artist.Similars {
		similar = append(similar, fmt.Sprintf("[%s](%s)", sim.Name, sim.URL))
	}

	return
}
