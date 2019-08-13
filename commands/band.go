package commands

import (
	"fmt"
	"image"
	"net/url"
	"os"
	"persephone/database"
	"persephone/lib"
	"persephone/utils"
	"strconv"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/cavaliercoder/grab"
	"github.com/fogleman/gg"
	"github.com/gocolly/colly"
	"github.com/nfnt/resize"
	"github.com/pazuzu156/aurora"
	"github.com/pazuzu156/lastfm-go"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
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
		tempmsg, _ := ctx.Message.Reply(ctx.Aurora, "Please wait while the artist image is generated...")
		defer ctx.Aurora.DeleteMessage(tempmsg.ChannelID, tempmsg.ID)
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

func (c Band) displayArtistInfo(ctx aurora.Context, artist lastfm.ArtistGetInfo) {
	user := database.GetUser(ctx.Message.Author)
	alist, _ := c.Command.Lastfm.User.GetTopAlbums(lastfm.P{"user": user.Lastfm, "limit": "1000"}) // limit max = 1000
	lfmuser, _ := database.GetLastfmUserInfo(ctx.Message.Author, c.Command.Lastfm)
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

	aimg := c.getArtistImage(artist)
	avres, _ := grab.Get("temp/", "https://cdn.discordapp.com/avatars/"+ctx.Message.Author.ID.String()+"/"+*ctx.Message.Author.Avatar+".png")
	bg := lib.OpenImage("static/images/background.png")
	av := lib.OpenImage(avres.Filename)
	os.Remove(avres.Filename)

	air := resize.Resize(230, 230, aimg, resize.Bicubic)
	avr := resize.Resize(72, 72, av, resize.Bicubic)

	dc := gg.NewContext(1000, 600)
	dc.DrawImage(bg, 0, 0)

	dc.SetRGBA(1, 1, 1, 0.2)
	dc.DrawRectangle(0, 100, 1000, 72)
	dc.Fill()

	dc.SetRGBA(0, 0, 0, 0.3)
	dc.DrawRoundedRectangle(50, 100, 240, 240, 10)
	dc.Fill()
	dc.DrawImage(air, 55, 105)

	dc.SetRGB(0.9, 0.9, 0.9)
	dc.LoadFontFace(FontBold, 20)
	dc.DrawStringWrapped(artist.Name, 50, 360, 0, 0, 230, 1.5, gg.AlignCenter)

	dc.LoadFontFace(FontRegular, 20)
	dc.DrawStringWrapped(fmt.Sprintf("%s plays", artist.Stats.UserPlays), 50, 375, 0, 0, 235, 1.5, gg.AlignCenter)

	dc.DrawImage(avr, 315, 100)
	dc.LoadFontFace(FontBold, 26)
	dc.SetRGB(0.9, 0.9, 0.9)
	dc.DrawString(ctx.Message.Author.Username+" ("+user.Lastfm+")", 390, 130)
	// scrobble count
	dc.LoadFontFace(FontRegular, 20)
	dc.SetRGB(0.9, 0.9, 0.9)
	printer := message.NewPrinter(language.English)
	pc, _ := strconv.Atoi(lfmuser.PlayCount)
	dc.DrawString(fmt.Sprintf("%s scrobbles", printer.Sprintf("%d", pc)), 390, 160)

	dc.SavePNG("temp/" + ctx.Message.Author.ID.String() + "_band.png")
	r, _ := os.Open("temp/" + ctx.Message.Author.ID.String() + "_band.png")

	ctx.Aurora.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
		Files: []disgord.CreateMessageFileParams{
			{
				FileName: r.Name(),
				Reader:   r,
			},
		},
	})

	r.Close()
	os.Remove("temp/" + ctx.Message.Author.ID.String() + "_band.png")
}

func (c Band) getArtistInfo(artist string) (lastfm.ArtistGetInfo, error) {
	return c.Command.Lastfm.Artist.GetInfo(lastfm.P{"artist": artist})
}

func (c Band) getArtistImage(artist lastfm.ArtistGetInfo) image.Image {
	col := colly.NewCollector()
	var imgsrc string
	col.OnHTML(".band_img a#photo", func(e *colly.HTMLElement) {
		imgsrc = e.ChildAttr("img", "src")
	})
	col.Visit(fmt.Sprintf("https://metal-archives.com/bands/%s", url.QueryEscape(artist.Name)))

	if imgsrc != "" {
		res, _ := grab.Get("temp/", imgsrc)
		img := lib.OpenImage(res.Filename)

		os.Remove(res.Filename)

		return img
	}

	ares, _ := grab.Get("temp/", artist.Images[3].Url)
	aimg := lib.OpenImage(ares.Filename)
	os.Remove(ares.Filename)

	return aimg
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
