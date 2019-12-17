package commands

import (
	"fmt"
	"persephone/fm"
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
		user := lib.GetUser(ctx.Message.Author)

		if len(ctx.Args) > 0 {
			args := ctx.Args
			artistName := lib.JoinStringMap(args, " ")

			fmt.Println(artistName)

			artist, err := c.Lastfm.Artist.GetInfo(lastfm.P{"artist": artistName, "username": user.Lastfm})

			if err != nil {
				ctx.Message.Reply(ctx.Atlas, "That artist could not be found")

				return
			}

			c.displayBandInfo(ctx, artist)
		} else {
			np, err := fm.GetNowPlayingTrack(ctx.Message.Author, c.Lastfm)

			if err != nil {
				ctx.Message.Reply(ctx.Atlas, "You're not currently listening to anything")

				return
			}

			artist, err := c.Lastfm.Artist.GetInfo(lastfm.P{"artist": np.Artist.Name, "username": user.Lastfm})

			if err != nil {
				ctx.Message.Reply(ctx.Atlas, "That artist could not be found")

				return
			}

			c.displayBandInfo(ctx, artist)
		}
	}

	return c.CommandInterface
}

func (c Bandinfo) appendToFields(fields []*disgord.EmbedField, name string, item string, inline bool) []*disgord.EmbedField {
	if "" != item {
		field := &disgord.EmbedField{
			Name:   name,
			Value:  item,
			Inline: inline,
		}

		fields = append(fields, field)
	}

	return fields
}

func (c Bandinfo) displayBandInfo(ctx atlas.Context, artist lastfm.ArtistGetInfo) {
	f, t := c.embedFooter(ctx)
	bio := lib.ShortStr(lib.HTMLParse(artist.Bio.Content), 850, fmt.Sprintf(" [Read More...](%s)", artist.Bio.Links[0].URL))

	if "" == bio {
		bio = "No bio found"
	}

	fields := []*disgord.EmbedField{
		{
			Name:  "Summary",
			Value: bio,
		},
	}

	fields = c.appendToFields(fields, "Tags", lib.JoinString(c.tags(artist), ", "), false)
	fields = c.appendToFields(fields, "Similar Artists", lib.JoinString(c.similar(artist), ", "), false)
	fields = c.appendToFields(fields, "Total Listeners", lib.HumanNumber(artist.Stats.Listeners), true)
	fields = c.appendToFields(fields, "Total Play Count", lib.HumanNumber(artist.Stats.Plays), true)
	fields = c.appendToFields(fields, "Your Play Count", lib.HumanNumber(artist.Stats.UserPlays), true)

	ctx.Atlas.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
		Embed: &disgord.Embed{
			Title: artist.Name,
			URL:   artist.URL,
			Color: lib.RandomColor(),
			Thumbnail: &disgord.EmbedThumbnail{
				URL: lib.GetArtistImageURLFromFmArtist(artist),
			},
			Fields: fields,
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
