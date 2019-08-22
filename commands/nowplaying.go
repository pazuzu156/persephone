package commands

import (
	"fmt"
	"os"
	"persephone/database"
	"persephone/lib"
	"persephone/utils"
	"strconv"

	"github.com/andersfylling/disgord"
	"github.com/cavaliercoder/grab"
	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"github.com/pazuzu156/aurora"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Nowplaying command.
type Nowplaying struct{ Command }

// InitNowPlaying initializes the nowplaying command.
func InitNowPlaying() Nowplaying {
	return Nowplaying{InitCmd(&CommandItem2{
		Name:        "nowplaying",
		Description: "Shows what you're currently listening to",
		Aliases:     []string{"np"},
	})}
}

// Register registers and runs the nowplaying command.
func (c Nowplaying) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		track, err := utils.GetNowPlayingTrack(ctx.Message.Author, c.Command.Lastfm)
		lfmuser, _ := database.GetLastfmUserInfo(ctx.Message.Author, c.Command.Lastfm)

		if err == nil {
			res, _ := grab.Get(lib.LocGet("temp/"), track.Images[3].URL)
			avres, _ := grab.Get(lib.LocGet("temp/"), "https://cdn.discordapp.com/avatars/"+ctx.Message.Author.ID.String()+"/"+*ctx.Message.Author.Avatar+".png")

			// Open base images
			bg := lib.OpenImage(lib.LocGet("static/images/background.png"))
			aa := lib.OpenImage(res.Filename)
			av := lib.OpenImage(avres.Filename)

			// delete downloaded images (they're already loaded into memory)
			os.Remove(res.Filename)
			os.Remove(avres.Filename)

			// Some resizing for avatar and album art
			aar := resize.Resize(240, 240, aa, resize.Bicubic)
			avr := resize.Resize(72, 72, av, resize.Bicubic)

			// New image context, and add background image
			dc := gg.NewContext(1000, 600)
			dc.DrawImage(bg, 0, 0)

			// Draw avatar (also add the white bar that goes behind the avatar image)
			dc.SetRGBA(1, 1, 1, 0.2)
			dc.DrawRectangle(0, 100, 1000, 72)
			dc.Fill()

			// Draw avatar and add username + scrobble count
			dc.DrawImage(avr, 315, 100)
			dc.LoadFontFace(FontBold, 26)
			dc.SetRGB(0.9, 0.9, 0.9)
			dc.DrawString(ctx.Message.Author.Username+" ("+lfmuser.Name+")", 390, 130)
			// scrobble count
			dc.LoadFontFace(FontRegular, 20)
			dc.SetRGB(0.9, 0.9, 0.9)
			printer := message.NewPrinter(language.English)
			pc, _ := strconv.Atoi(lfmuser.PlayCount)
			dc.DrawString(fmt.Sprintf("%s scrobbles", printer.Sprintf("%d", pc)), 390, 160)

			// Draw white box that goes behind album art + draw album art
			dc.SetRGBA(1, 1, 1, 0.2)
			dc.DrawRectangle(50, 0, 250, 600)
			dc.Fill()
			dc.DrawImage(aar, 55, 105)

			// Draw artist name
			dc.LoadFontFace(FontBold, 20)
			dc.SetRGB(0.9, 0.9, 0.9)
			dc.DrawStringWrapped(track.Artist.Name, 70, 370, 0, 0, 200, 1.5, gg.AlignLeft)

			// Draw album + track name
			dc.LoadFontFace(FontRegular, 20)
			dc.SetRGB(0.9, 0.9, 0.9)
			dc.DrawStringWrapped(track.Album.Name+" - "+track.Name, 70, 420, 0, 0, 200, 1.5, gg.AlignLeft)

			// This gets the last 3 listened tracks and draws
			// images + text for each
			// tracks are layered in reverse order 3 -> 2 -> 1 displayed
			// in ascending order
			tracks, _ := utils.GetRecentTracks(ctx.Message.Author, c.Command.Lastfm, "3")
			if len(tracks) > 3 {
				t1 := tracks[1] // most recent track
				t2 := tracks[2] // second most recent track
				t3 := tracks[3] // third most recent track

				// This needs to be rendered first, as it'll be in the back
				// behind the other 2 recent tracks
				if img := t3.Images[3].URL; img != "" {
					dc.SetRGBA(0, 0, 0, 0.3)
					dc.DrawRoundedRectangle(336, 396, 168, 168, 85)
					dc.Fill()

					dc.DrawRoundedRectangle(340, 400, 160, 160, 80)
					dc.Clip()
					i, _ := grab.Get(lib.LocGet("temp/"), img)
					ii := lib.OpenImage(i.Filename)
					iir := resize.Resize(160, 160, ii, resize.Bicubic)
					dc.DrawImage(iir, 340, 400)
					dc.ResetClip()
					os.Remove(i.Filename)

					dc.SetRGB(0.9, 0.9, 0.9)
					dc.LoadFontFace(FontBold, 25)
					dc.DrawString(t3.Artist.Name, 510, 480)

					dc.LoadFontFace(FontRegular, 25)
					dc.DrawString(t3.Name, 510, 520)
				}

				// Track 2
				if img := t2.Images[3].URL; img != "" {
					dc.SetRGBA(0, 0, 0, 0.3)
					dc.DrawRoundedRectangle(336, 306, 168, 168, 85)
					dc.Fill()

					dc.DrawRoundedRectangle(340, 310, 160, 160, 80)
					dc.Clip()
					i, _ := grab.Get(lib.LocGet("temp/"), img)
					ii := lib.OpenImage(i.Filename)
					iir := resize.Resize(160, 160, ii, resize.Bicubic)
					dc.DrawImage(iir, 340, 310)
					dc.ResetClip()
					os.Remove(i.Filename)

					dc.SetRGB(0.9, 0.9, 0.9)
					dc.LoadFontFace(FontBold, 25)
					dc.DrawString(t2.Artist.Name, 510, 380)

					dc.LoadFontFace(FontRegular, 25)
					dc.DrawString(t2.Name, 510, 420)
				}

				// First most recent track
				if img := t1.Images[3].URL; img != "" {
					dc.SetRGBA(0, 0, 0, 0.3)
					dc.DrawRoundedRectangle(336, 216, 168, 168, 85)
					dc.Fill()

					dc.DrawRoundedRectangle(340, 220, 160, 160, 80)
					dc.Clip()
					i, _ := grab.Get(lib.LocGet("temp/"), img)
					ii := lib.OpenImage(i.Filename)
					iir := resize.Resize(160, 160, ii, resize.Bicubic)
					dc.DrawImage(iir, 340, 220)
					dc.ResetClip()
					os.Remove(i.Filename)

					dc.SetRGB(0.9, 0.9, 0.9)
					dc.LoadFontFace(FontBold, 25)
					dc.DrawString(t1.Artist.Name, 510, 280)

					dc.LoadFontFace(FontRegular, 25)
					dc.DrawString(t1.Name, 510, 320)
				}
			}

			lib.BrandImage(dc) // brand image

			dc.SavePNG(lib.LocGet("temp/" + ctx.Message.Author.ID.String() + "_np.png"))      // save generated image
			r, _ := os.Open(lib.LocGet("temp/" + ctx.Message.Author.ID.String() + "_np.png")) // open generated image into memory

			// create new message with the image + embed with a link to the user's Last.fm profile page
			ctx.Aurora.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
				Files: []disgord.CreateMessageFileParams{
					{FileName: ctx.Message.Author.ID.String() + ".png", Reader: r},
				},
				Embed: &disgord.Embed{
					Title: fmt.Sprintf("View %s's Profile on Last.fm", ctx.Message.Author.Username),
					URL:   fmt.Sprintf("https://last.fm/user/%s", lfmuser.Name),
					Color: utils.RandomColor(),
				},
			})

			// Close and delete the image
			r.Close()
			os.Remove(lib.LocGet("temp/" + ctx.Message.Author.ID.String() + "_np.png"))
		} else {
			ctx.Message.Reply(ctx.Aurora, err.Error())
		}
	}

	return c.Command.CommandInterface
}
