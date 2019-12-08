package commands

import (
	"fmt"
	"persephone/lib"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
	"github.com/pazuzu156/lastfm-go"
)

// Bandinfo command.
type Bandinfo struct{ Command }

// InitBandinfo initializes the bandinfo command.
func InitBandinfo() Bandinfo {
	return Bandinfo{Init(&CommandItem{
		Name:        "bandinfo",
		Description: "Gets information on the current playing band",
		Aliases:     []string{"bi", "artistinfo", "ai"},
		Usage:       "bandinfo Darkthrone",
		Parameters:  []Parameter{},
	})}
}

// Register registers and runs the bandinfo command.
func (c Bandinfo) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		// ctx.Message.Reply(ctx.Atlas, "Hello, Bandinfo!")
		if len(ctx.Args) > 0 {
			args := ctx.Args
			artistName := lib.JoinStringMap(args, " ")

			fmt.Println(artistName)

			artist, err := c.Lastfm.Artist.GetInfo(lastfm.P{"artist": artistName})

			if err != nil {
				ctx.Message.Reply(ctx.Atlas, "That artist could not be found")

				return
			}

			c.displayBandInfo(ctx, artist)
		} else {
			np, err := lib.GetNowPlayingTrack(ctx.Message.Author, c.Lastfm)

			if err != nil {
				ctx.Message.Reply(ctx.Atlas, "You're not currently listening to anything")

				return
			}

			artist, _ := c.Lastfm.Artist.GetInfo(lastfm.P{"artist": np.Artist.Name})
			c.displayBandInfo(ctx, artist)
		}
	}

	return c.CommandInterface
}

func (c Bandinfo) displayBandInfo(ctx atlas.Context, artist lastfm.ArtistGetInfo) {
	f, t := c.embedFooter(ctx)
	ctx.Atlas.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
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
			Footer: f, Timestamp: t,
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
