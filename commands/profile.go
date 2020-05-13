package commands

import (
	"fmt"
	"image"
	"os"
	"strconv"

	"github.com/andersfylling/disgord"
	"github.com/cavaliercoder/grab"
	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"github.com/pazuzu156/atlas"
	"github.com/pazuzu156/lastfm-go"
	"github.com/pazuzu156/persephone/lib"
)

// Profile command.
type Profile struct{ Command }

// InitProfile initializes the profile command.
func InitProfile() Profile {
	return Profile{Init(&CommandItem{
		Name:        "profile",
		Description: "Shows your top everything",
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
		if user := lib.GetUser(ctx.Message.Author); user.Username != "" {
			lfmuser, _ := lib.GetLastfmUserInfo(ctx.Message.Author, c.Lastfm)

			avURL, _ := ctx.Message.Author.AvatarURL(256, false)
			res, _ := lib.Grab(avURL)
			av, avf := lib.OpenImage(res.Filename)

			os.Remove(avf.Name())

			avr := resize.Resize(72, 72, av, resize.Bicubic)

			artists, _ := c.Lastfm.User.GetTopArtists(lastfm.P{"user": lfmuser.Name, "period": "overall", "limit": "5"})
			albums, _ := c.Lastfm.User.GetTopAlbums(lastfm.P{"user": lfmuser.Name, "limit": "4", "period": "overall"})
			tracks, _ := c.Lastfm.User.GetTopTracks(lastfm.P{"user": c.getLastfmUserFromCtx(ctx), "limit": "8"})

			bg, _ := lib.OpenImage(lib.LocGet("static/images/background.png"))

			dc := gg.NewContext(1000, 600)
			dc.DrawImage(bg, 0, 0)

			dc.SetRGBA(1, 1, 1, 0.2)
			dc.DrawRectangle(0, 50, 1000, 72)
			dc.Fill()

			// Draw avatar and add username + scrobble count
			dc.DrawImage(avr, 315, 50)
			dc.LoadFontFace(FontBold, 26)
			lib.DrawStringWithShadow(ctx.Message.Author.Username+" ("+lfmuser.Name+")", 390, 80, dc)
			// scrobble count
			dc.LoadFontFace(FontRegular, 20)
			lib.DrawStringWithShadow(fmt.Sprintf("%s scrobbles", lib.HumanNumber(lfmuser.PlayCount)), 390, 110, dc)

			// Draw white box that goes behind album art + draw album art
			dc.SetRGBA(1, 1, 1, 0.2)
			dc.DrawRectangle(50, 0, 250, 600)
			dc.Fill()

			dc.LoadFontFace(FontBold, 40)
			lib.DrawStringWithShadow("Top Artists", 65, 90, dc)

			var aa image.Image = nil

			if len(artists.Artists) > 0 {
				aa = lib.GetArtistImage(artists.Artists[0])
			}

			if aa == nil {
				aa, _ = lib.OpenImage(lib.LocGet("static/images/bm.png"))
			}

			aar := resize.Resize(240, 240, aa, resize.Bicubic)
			dc.DrawImage(aar, 55, 105)

			dc.LoadFontFace(FontRegular, 25)

			if len(artists.Artists) > 0 {
				lib.DrawStringWithShadow(lib.ShortStr(artists.Artists[0].Name, 33), 60, 370, dc)
				lib.DrawStringWithShadow(fmt.Sprintf("%s plays", artists.Artists[0].PlayCount), 60, 400, dc)
			}

			x := float64(61)
			y := float64(445)

			// loop and display artists
			for n, artist := range artists.Artists {
				if n > 0 {
					dc.LoadFontFace(FontRegular, 20)

					lib.DrawStringWithShadow(lib.ShortStr(artist.Name, 18), x, y, dc)
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
			for i, albumItem := range albums.Albums {
				album, _ := c.Lastfm.Album.GetInfo(lastfm.P{"artist": albumItem.Artist.Name,
					"album": albumItem.Name, "username": c.getLastfmUserFromCtx(ctx)})
				ares, _ := grab.Get(lib.LocGet("temp/"), album.Images[3].URL)
				ai, _ := lib.OpenImage(ares.Filename)
				os.Remove(ares.Filename) // possible suspect to deletion of bm.png. keep a watch on this
				ar := resize.Resize(145, 145, ai, resize.Bicubic)
				pos := AlbumPositions[i]

				dc.SetRGBA(0, 0, 0, 0.3)
				dc.DrawRoundedRectangle(pos.Shadow.X, pos.Shadow.Y, 155, 155, pos.Shadow.R)
				dc.Fill()

				dc.DrawImage(ar, pos.X, pos.Y)

				dc.LoadFontFace(FontRegular, 20)
				lib.DrawStringWithShadow(lib.ShortStr(album.Name, 14), pos.Info.X, pos.Info.Y, dc)

				dc.LoadFontFace(FontRegular, 16)
				plays, _ := strconv.Atoi(album.UserPlayCount)
				plays = plays / len(album.Tracks)
				lib.DrawStringWithShadow(fmt.Sprintf("%d album plays", plays), pos.Info.Plays.X, pos.Info.Plays.Y, dc)
			}

			dc.LoadFontFace(FontRegular, 20)
			lib.DrawStringWithShadow("Top Tracks", 800, 150, dc)

			for i, track := range tracks.Tracks {
				pos := TrackPositions[i]

				dc.LoadFontFace(FontRegular, 20)
				lib.DrawStringWithShadow(lib.ShortStr(track.Name, 20), pos.X, pos.Y, dc)

				dc.LoadFontFace(FontRegular, 16)
				lib.DrawStringWithShadow(fmt.Sprintf("%s | %s plays", lib.ShortStr(track.Artist.Name, 10), track.PlayCount), pos.Plays.X, pos.Plays.Y, dc)
			}

			lib.BrandImage(dc)

			r, _ := lib.SaveImage(dc, ctx, "profile")

			ctx.Atlas.CreateMessage(ctx.Context, ctx.Message.ChannelID, &disgord.CreateMessageParams{
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

			r.Close()
			os.Remove(r.Name())
		} else {
			ctx.Message.Reply(ctx.Context, ctx.Atlas, "You're not currently logged in with Last.fm")
		}
	}

	return c.CommandInterface
}
