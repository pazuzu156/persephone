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

var albumPositions = []lib.AlbumPosition{
	{
		X: 355,
		Y: 170,
		Shadow: lib.Shadow{
			X: 350,
			Y: 165,
			R: 10,
		},
		Info: lib.InfoText{
			X: 350,
			Y: 340,
			Plays: lib.PlaysText{
				X: 350,
				Y: 360,
			},
		},
	},
	{
		X: 555,
		Y: 170,
		Shadow: lib.Shadow{
			X: 550,
			Y: 165,
			R: 10,
		},
		Info: lib.InfoText{
			X: 550,
			Y: 340,
			Plays: lib.PlaysText{
				X: 550,
				Y: 360,
			},
		},
	},
	{
		X: 355,
		Y: 390,
		Shadow: lib.Shadow{
			X: 350,
			Y: 385,
			R: 10,
		},
		Info: lib.InfoText{
			X: 350,
			Y: 560,
			Plays: lib.PlaysText{
				X: 350,
				Y: 580,
			},
		},
	},
	{
		X: 555,
		Y: 390,
		Shadow: lib.Shadow{
			X: 550,
			Y: 385,
			R: 10,
		},
		Info: lib.InfoText{
			X: 550,
			Y: 560,
			Plays: lib.PlaysText{
				X: 550,
				Y: 580,
			},
		},
	},
}

var trackPositions = []lib.TrackPosition{
	{
		X: 720,
		Y: 180,
		Plays: lib.PlaysText{
			X: 870,
			Y: 180,
		},
	},
	{
		X: 720,
		Y: 210,
		Plays: lib.PlaysText{
			X: 870,
			Y: 210,
		},
	},
	{
		X: 720,
		Y: 240,
		Plays: lib.PlaysText{
			X: 870,
			Y: 240,
		},
	},
	{
		X: 720,
		Y: 270,
		Plays: lib.PlaysText{
			X: 870,
			Y: 270,
		},
	},
}

// Band command.
type Band struct {
	Command Command
}

// InitBand initializes the band command.
func InitBand() Band {
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
		"b",
	)}
}

// Register registers and runs the help command.
func (c Band) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		tempmsg, _ := ctx.Message.Reply(ctx.Aurora, "Please wait while the artist image is generated...")
		defer ctx.Aurora.DeleteMessage(tempmsg.ChannelID, tempmsg.ID)
		if len(ctx.Args) > 0 {
			artistName := strings.Trim(strings.Join(ctx.Args, " "), " ")
			artist, err := c.getArtistInfo(artistName, ctx.Message.Author)

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

			artist, err := c.getArtistInfo(track.Artist.Name, ctx.Message.Author)

			if err != nil {
				ctx.Message.Reply(ctx.Aurora, "Couldn't find that artist")
			}

			c.displayArtistInfo(ctx, artist)
		}
	}

	return c.Command.CommandInterface
}

func (c Band) displayArtistInfo(ctx aurora.Context, artist lastfm.ArtistGetInfo) {
	albums := c.getAlbumsList(ctx, artist)
	tracks := c.getTracksList(ctx, artist)
	lfmuser, _ := database.GetLastfmUserInfo(ctx.Message.Author, c.Command.Lastfm)

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
	dc.DrawRectangle(0, 50, 1000, 72)
	dc.Fill()

	dc.SetRGBA(0, 0, 0, 0.3)
	dc.DrawRoundedRectangle(50, 50, 240, 240, 10)
	dc.Fill()
	dc.DrawImage(air, 55, 55)

	dc.SetRGB(0.9, 0.9, 0.9)
	dc.LoadFontFace(FontBold, 20)
	dc.DrawStringWrapped(artist.Name, 50, 310, 0, 0, 230, 1.5, gg.AlignCenter)

	dc.LoadFontFace(FontRegular, 20)
	dc.DrawStringWrapped(fmt.Sprintf("%s plays", artist.Stats.UserPlays), 50, 345, 0, 0, 235, 1.5, gg.AlignCenter)

	dc.DrawLine(50, 370, 285, 370)
	dc.SetLineWidth(0.5)
	dc.Stroke()

	var tags string

	for _, tag := range artist.Tags {
		tags += tag.Name + ", "
	}

	dc.DrawStringWrapped(strings.TrimRight(tags, ", "), 50, 380, 0, 0, 235, 1.5, gg.AlignCenter)

	dc.DrawImage(avr, 315, 50)
	dc.LoadFontFace(FontBold, 26)
	dc.SetRGB(0.9, 0.9, 0.9)
	dc.DrawString(ctx.Message.Author.Username+" ("+lfmuser.Name+")", 400, 80)
	// scrobble count
	dc.LoadFontFace(FontRegular, 20)
	dc.SetRGB(0.9, 0.9, 0.9)
	printer := message.NewPrinter(language.English)
	pc, _ := strconv.Atoi(lfmuser.PlayCount)
	dc.DrawString(fmt.Sprintf("%s scrobbles", printer.Sprintf("%d", pc)), 400, 110)

	dc.DrawString("Albums", 490, 150)

	for i, album := range albums {
		if i < len(albums) && i < 4 {
			ares, _ := grab.Get("temp/", album.Images[3].Url)
			ai := lib.OpenImage(ares.Filename)
			os.Remove(ares.Filename)
			ar := resize.Resize(145, 145, ai, resize.Bicubic)
			pos := albumPositions[i]

			dc.SetRGBA(0, 0, 0, 0.3)
			dc.DrawRoundedRectangle(pos.Shadow.X, pos.Shadow.Y, 155, 155, pos.Shadow.R)
			dc.Fill()

			dc.DrawImage(ar, pos.X, pos.Y)

			dc.SetRGBA(1, 1, 1, 0.9)
			dc.LoadFontFace(FontRegular, 20)
			dc.DrawString(utils.ShortStr(album.Name, 15), pos.Info.X, pos.Info.Y)
			dc.LoadFontFace(FontRegular, 16)
			dc.DrawString(fmt.Sprintf("%s plays", album.PlayCount), pos.Info.Plays.X, pos.Info.Plays.Y)
		}
	}

	dc.LoadFontFace(FontRegular, 20)
	dc.DrawString("Tracks", 790, 150)

	for i, track := range tracks {
		if i < len(tracks) && i < 4 {
			pos := trackPositions[i]
			dc.SetRGB(0.9, 0.9, 0.9)
			dc.LoadFontFace(FontRegular, 16)
			dc.DrawString(utils.ShortStr(track.Name, 15), pos.X, pos.Y)
			dc.LoadFontFace(FontBold, 16)
			dc.DrawString(fmt.Sprintf("%s plays", track.PlayCount), pos.Plays.X, pos.Plays.Y)

		}
	}

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

func (c Band) getArtistInfo(artist string, user *disgord.User) (lastfm.ArtistGetInfo, error) {
	dbu := database.GetUser(user)
	return c.Command.Lastfm.Artist.GetInfo(lastfm.P{"artist": artist, "username": dbu.Lastfm})
}

func (c Band) getArtistImage(artist lastfm.ArtistGetInfo) image.Image {
	col := colly.NewCollector()
	var imgsrc string
	col.OnHTML(".band_img a#photo", func(e *colly.HTMLElement) {
		imgsrc = e.ChildAttr("img", "src")
	})

	maartist := lib.GetMaArtist(artist.Name)

	if maartist.ID != 0 {
		col.Visit(fmt.Sprintf("https://metal-archives.com/bands/%s/%d", url.QueryEscape(maartist.Name), maartist.ID))
	} else {
		col.Visit(fmt.Sprintf("https://metal-archives.com/bands/%s", url.QueryEscape(artist.Name)))
	}

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

func (c Band) getAlbumsList(ctx aurora.Context, artist lastfm.ArtistGetInfo) []lib.TopAlbum {
	user := database.GetUser(ctx.Message.Author)
	alist, _ := c.Command.Lastfm.User.GetTopAlbums(lastfm.P{"user": user.Lastfm, "limit": "1000"}) // limit max = 1000
	var albums = []lib.TopAlbum{}

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

	return albums
}

func (c Band) getTracksList(ctx aurora.Context, artist lastfm.ArtistGetInfo) []lib.TopTrack {
	user := database.GetUser(ctx.Message.Author)
	tlist, _ := c.Command.Lastfm.User.GetTopTracks(lastfm.P{"user": user.Lastfm, "limit": "1000"})
	var tracks = []lib.TopTrack{}

	for _, track := range tlist.Tracks {
		if track.Artist.Name == artist.Name {
			tracks = append(tracks, track)
		}
	}

	if tlist.TotalPages > 1 {
		for i := 1; i <= tlist.TotalPages; i++ {
			tl, _ := c.Command.Lastfm.User.GetTopTracks(lastfm.P{"user": user.Lastfm, "limit": "1000", "page": strconv.Itoa(i)})

			if tl.Tracks[i].Artist.Name == artist.Name {
				tracks = append(tracks, tl.Tracks[i])
			}
		}
	}

	return tracks
}
