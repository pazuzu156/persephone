package commands

import (
	"fmt"
	"persephone/database"
	"persephone/utils"
	"strconv"
	"strings"

	"github.com/pazuzu156/aurora"
	"github.com/pazuzu156/lastfm-go"
)

type Band struct {
	Command Command
}

func InitBand(aliases ...string) Band {
	return Band{Init(
		"bandinfo",
		"Gets information on a band",
		[]UsageItem{
			{
				Command:     "band",
				Description: "Gets information on the artist you're currently listening to",
			},
			{
				Command:     "band [artist]",
				Description: "Gets information on a requested artist",
			},
		},
		aliases...,
	)}
}

func (c Band) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		if len(ctx.Args) > 0 {
			artistName := strings.Trim(strings.Join(ctx.Args, " "), " ")
			artist, err := c.getArtistInfo(artistName)

			if err != nil {
				ctx.Message.Reply(ctx.Aurora, err.Error())

				return
			}

			c.displayArtistInfo(ctx, artist)
		} else {
			track, err := utils.GetNowPlayingTrack(ctx.Message.Author, c.Command.Lastfm)

			if err != nil {
				ctx.Message.Reply(ctx.Aurora, err.Error())

				return
			}

			artist, err := c.getArtistInfo(track.Artist.Name)

			if err != nil {
				ctx.Message.Reply(ctx.Aurora, "Couldn't find that artist")
			}

			c.displayArtistInfo(ctx, artist)
		}
	}

	return c.Command.CommandInterface
}

func (c Band) displayArtistInfo(context aurora.Context, artist lastfm.ArtistGetInfo) {
	user := database.GetUser(context.Message.Author)
	alist, _ := c.Command.Lastfm.User.GetTopAlbums(lastfm.P{"user": user.Lastfm, "limit": "1000"}) // limit max = 1000
	var albums = []Album{}

	for _, album := range alist.Albums {
		if album.Artist.Name == artist.Name {
			albums = append(albums, album)
		}
	}

	if alist.TotalPages > 1 {
		for i := 1; i <= alist.TotalPages; i++ {
			al, _ := c.Command.Lastfm.User.GetTopAlbums(lastfm.P{"user": user.Lastfm, "limit": "1000", "page": strconv.Itoa(i)})

			if al.Albums[i].Artist.Name == artist.Name {
				albums = append(albums, al.Albums[i])
			}
		}
	}

	fmt.Println(albums) // These are displayed as played albums
}

func (c Band) getArtistInfo(artist string) (lastfm.ArtistGetInfo, error) {
	return c.Command.Lastfm.Artist.GetInfo(lastfm.P{"artist": artist})
}

type Album struct {
	Rank      string `xml:"rank,attr"`
	Name      string `xml:"name"`
	PlayCount string `xml:"playcount"`
	Mbid      string `xml:"mbid"`
	Url       string `xml:"url"`
	Artist    struct {
		Name string `xml:"name"`
		Mbid string `xml:"mbid"`
		Url  string `xml:"url"`
	} `xml:"artist"`
	Images []struct {
		Size string `xml:"size,attr"`
		Url  string `xml:",chardata"`
	} `xml:"image"`
}
