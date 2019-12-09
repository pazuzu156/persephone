package commands

import (
	"fmt"
	"os"
	"persephone/lib"

	"github.com/andersfylling/disgord"
	"github.com/cavaliercoder/grab"
	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"github.com/pazuzu156/atlas"
	"github.com/pazuzu156/lastfm-go"
)

// Profile command.
type Profile struct{ Command }

// InitProfile initializes the profile command.
func InitProfile() Profile {
	return Profile{Init(&CommandItem{
		Name:        "profile",
		Description: "Shows your top everything",
		Aliases:     []string{"p"},
		Usage:       "profile [member]",
		Parameters: []Parameter{
			{
				Name:        "member",
				Description: "The user you want to see a profile of",
			},
		},
	})}
}

// Register registers and runs the profile command.
func (c Profile) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		lfmuser, _ := lib.GetLastfmUserInfo(ctx.Message.Author, c.Lastfm)

		avURL, _ := ctx.Message.Author.AvatarURL(256, false)
		res, _ := lib.Grab(avURL)
		av, avf := lib.OpenImage(res.Filename)

		os.Remove(avf.Name())

		avr := resize.Resize(72, 72, av, resize.Bicubic)

		// TODO: Get Top Artists
		artists, _ := c.Lastfm.User.GetTopArtists(lastfm.P{"user": lfmuser.Name, "period": "overall", "limit": "5"})

		// TODO: Get Top Albums
		albums, _ := c.Lastfm.User.GetTopAlbums(lastfm.P{"user": lfmuser.Name, "limit": "4", "period": "overall"})

		// TODO: Get Top Tags

		bg, _ := lib.OpenImage(lib.LocGet("static/images/background.png"))

		dc := gg.NewContext(1000, 600)
		dc.DrawImage(bg, 0, 0)

		dc.SetRGBA(1, 1, 1, 0.2)
		dc.DrawRectangle(0, 50, 1000, 72)
		dc.Fill()

		// Draw avatar and add username + scrobble count
		dc.DrawImage(avr, 315, 50)
		dc.LoadFontFace(FontBold, 26)
		dc.SetRGB(0.2, 0.2, 0.2)
		dc.DrawString(ctx.Message.Author.Username+" ("+lfmuser.Name+")", 391, 80)
		dc.SetRGB(0.9, 0.9, 0.9)
		dc.DrawString(ctx.Message.Author.Username+" ("+lfmuser.Name+")", 390, 79)
		// scrobble count
		dc.LoadFontFace(FontRegular, 20)
		dc.SetRGB(0.2, 0.2, 0.2)
		dc.DrawString(fmt.Sprintf("%s scrobbles", lib.HumanNumber(lfmuser.PlayCount)), 391, 111)
		dc.SetRGB(0.9, 0.9, 0.9)
		dc.DrawString(fmt.Sprintf("%s scrobbles", lib.HumanNumber(lfmuser.PlayCount)), 390, 110)

		// Draw white box that goes behind album art + draw album art
		dc.SetRGBA(1, 1, 1, 0.2)
		dc.DrawRectangle(50, 0, 250, 600)
		dc.Fill()

		dc.SetRGB(0.2, 0.2, 0.2)
		dc.LoadFontFace(FontBold, 40)
		dc.DrawString("Top Artists", 66, 91)
		dc.SetRGB(0.8, 0.8, 0.8)
		dc.DrawString("Top Artists", 65, 90)

		// display top artist and the other top 4
		aa := lib.GetArtistImage(artists.Artists[0])

		if aa == nil {
			aa, _ = lib.OpenImage(lib.LocGet("static/images/bm.png"))
		}

		aar := resize.Resize(240, 240, aa, resize.Bicubic)
		dc.DrawImage(aar, 55, 105)

		dc.LoadFontFace(FontRegular, 25)
		lib.DrawStringWithShadow(lib.ShortStr(artists.Artists[0].Name, 33), 60, 370, dc)
		lib.DrawStringWithShadow(fmt.Sprintf("%s plays", artists.Artists[0].PlayCount), 60, 400, dc)

		x := float64(61)
		y := float64(445)

		// loop and display artists
		for n, artist := range artists.Artists {
			if n > 0 {
				dc.LoadFontFace(FontRegular, 20)

				lib.DrawStringWithShadow(lib.ShortStr(artist.Name, 25), x, y, dc)
				lib.DrawWrappedStringWithShadow(artist.PlayCount, 50, y, 0, 1, 240, 0, gg.AlignRight, dc)

				dc.SetRGB(0.2, 0.2, 0.2)
				dc.SetLineWidth(0.5)
				dc.DrawLine(x, y+5, x+230, y+5)
				dc.Stroke()
				dc.SetRGB(0.9, 0.9, 0.9)
				dc.DrawLine(x-1, y+4, x+229, y+4)
				dc.Stroke()

				y = y + 40
			}
		}

		dc.LoadFontFace(FontRegular, 20)
		lib.DrawStringWithShadow("Top Albums", 475, 150, dc)

		// display album grid.
		for i, album := range albums.Albums {
			ares, _ := grab.Get(lib.LocGet("temp/"), album.Images[3].URL)
			ai, _ := lib.OpenImage(ares.Filename)
			os.Remove(ares.Filename)
			ar := resize.Resize(145, 145, ai, resize.Bicubic)
			pos := AlbumPositions[i]

			dc.SetRGBA(0, 0, 0, 0.3)
			dc.DrawRoundedRectangle(pos.Shadow.X, pos.Shadow.Y, 155, 155, pos.Shadow.R)
			dc.Fill()

			dc.DrawImage(ar, pos.X, pos.Y)

			dc.LoadFontFace(FontRegular, 20)
			lib.DrawStringWithShadow(lib.ShortStr(album.Name, 15), pos.Info.X, pos.Info.Y, dc)

			dc.LoadFontFace(FontRegular, 16)
			lib.DrawStringWithShadow(fmt.Sprintf("%s plays", album.PlayCount), pos.Info.Plays.X, pos.Info.Plays.Y, dc)
		}

		lib.BrandImage(dc)

		r, _ := lib.SaveImage(dc, ctx, "profile")

		ctx.Atlas.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
			Files: []disgord.CreateMessageFileParams{
				{
					FileName: lib.TagImageName(ctx, "profile") + ".png",
					Reader:   r,
				},
			},
			Embed: &disgord.Embed{
				Title: fmt.Sprintf("View %s's Profile on Last.fm", ctx.Message.Author.Username),
				URL:   fmt.Sprintf("https://last.fm/user/%s", lfmuser.Name),
				Color: lib.RandomColor(),
			},
		})
	}

	return c.CommandInterface
}
