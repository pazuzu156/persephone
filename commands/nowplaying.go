package commands

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"persephone/models"
	"strconv"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/cavaliercoder/grab"
	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"github.com/polaron/aurora"
	"github.com/shkh/lastfm-go/lastfm"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Nowplaying command.
type Nowplaying struct {
	Command Command
}

// InitNowPlaying initializes the nowplaying command.
func InitNowPlaying(aliases ...string) Nowplaying {
	return Nowplaying{Init(
		"nowplaying",
		"Shows what you're currently listening to",
		aliases...,
	)}
}

// Register registers and runs the nowplaying command.
func (c Nowplaying) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		models.GetUser(ctx.Message.Author)

		np, _ := c.Command.Lastfm.User.GetRecentTracks(lastfm.P{
			"user":  "Pazuzu156",
			"limit": "3",
		})

		track := np.Tracks[0]                                                   // want first track
		user, _ := c.Command.Lastfm.User.GetInfo(lastfm.P{"user": "Pazuzu156"}) // gets the user

		if track.NowPlaying == "true" {
			res, _ := grab.Get("temp/", track.Images[3].Url)
			avres, _ := grab.Get("temp/", "https://cdn.discordapp.com/avatars/"+ctx.Message.Author.ID.String()+"/"+*ctx.Message.Author.Avatar+".png")

			// Open base images
			bg := openImage("static/images/background.png")
			aa := openImage(res.Filename)
			av := openImage(avres.Filename)

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
			dc.LoadFontFace("static/fonts/NotoSans-Bold.ttf", 26)
			dc.SetRGB(0.9, 0.9, 0.9)
			dc.DrawString(ctx.Message.Author.Username+" ("+user.Name+")", 390, 130)
			// scrobble count
			dc.LoadFontFace("static/fonts/NotoSans-Regular.ttf", 20)
			dc.SetRGB(0.9, 0.9, 0.9)
			printer := message.NewPrinter(language.English)
			pc, _ := strconv.Atoi(user.PlayCount)
			dc.DrawString(fmt.Sprintf("%s scrobbles", printer.Sprintf("%d", pc)), 390, 155)

			// Draw white box that goes behind album art + draw album art
			dc.SetRGBA(1, 1, 1, 0.2)
			dc.DrawRectangle(50, 0, 250, 600)
			dc.Fill()
			dc.DrawImage(aar, 55, 105)

			// Draw artist name
			dc.LoadFontFace("static/fonts/NotoSans-Bold.ttf", 20)
			dc.SetRGB(0.9, 0.9, 0.9)
			dc.DrawStringWrapped(track.Artist.Name, 70, 370, 0, 0, 200, 1.5, gg.AlignLeft)

			// Draw album + track name
			dc.LoadFontFace("static/fonts/NotoSans-Regular.ttf", 20)
			dc.SetRGB(0.9, 0.9, 0.9)
			dc.DrawStringWrapped(track.Album.Name+" - "+track.Name, 70, 420, 0, 0, 200, 1.5, gg.AlignLeft)

			// This gets the last 3 listened tracks and draws
			// images + text for each
			// tracks are layered in reverse order 3 -> 2 -> 1 displayed
			// in ascending order
			if len(np.Tracks) > 3 {
				t1 := np.Tracks[1] // most recent track
				t2 := np.Tracks[2] // second most recent track
				t3 := np.Tracks[3] // third most recent track

				// This needs to be rendered first, as it'll be in the back
				// behind the other 2 recent tracks
				if img := t3.Images[3].Url; img != "" {
					dc.SetRGBA(0, 0, 0, 0.3)
					dc.DrawRoundedRectangle(336, 396, 168, 168, 85)
					dc.Fill()

					dc.DrawRoundedRectangle(340, 400, 160, 160, 80)
					dc.Clip()
					i, _ := grab.Get("temp/", img)
					ii := openImage(i.Filename)
					iir := resize.Resize(160, 160, ii, resize.Bicubic)
					dc.DrawImage(iir, 340, 400)
					dc.ResetClip()
					os.Remove(i.Filename)

					dc.SetRGB(0.9, 0.9, 0.9)
					dc.LoadFontFace("static/fonts/NotoSans-Bold.ttf", 25)
					dc.DrawString(t3.Artist.Name, 510, 480)

					dc.LoadFontFace("static/fonts/NotoSans-Regular.ttf", 25)
					dc.DrawString(t3.Name, 510, 520)
				}

				// Track 2
				if img := t2.Images[3].Url; img != "" {
					dc.SetRGBA(0, 0, 0, 0.3)
					dc.DrawRoundedRectangle(336, 306, 168, 168, 85)
					dc.Fill()

					dc.DrawRoundedRectangle(340, 310, 160, 160, 80)
					dc.Clip()
					i, _ := grab.Get("temp/", img)
					ii := openImage(i.Filename)
					iir := resize.Resize(160, 160, ii, resize.Bicubic)
					dc.DrawImage(iir, 340, 310)
					dc.ResetClip()
					os.Remove(i.Filename)

					dc.SetRGB(0.9, 0.9, 0.9)
					dc.LoadFontFace("static/fonts/NotoSans-Bold.ttf", 25)
					dc.DrawString(t2.Artist.Name, 510, 380)

					dc.LoadFontFace("static/fonts/NotoSans-Regular.ttf", 25)
					dc.DrawString(t2.Name, 510, 420)
				}

				// First most recent track
				if img := t1.Images[3].Url; img != "" {
					dc.SetRGBA(0, 0, 0, 0.3)
					dc.DrawRoundedRectangle(336, 216, 168, 168, 85)
					dc.Fill()

					dc.DrawRoundedRectangle(340, 220, 160, 160, 80)
					dc.Clip()
					i, _ := grab.Get("temp/", img)
					ii := openImage(i.Filename)
					iir := resize.Resize(160, 160, ii, resize.Bicubic)
					dc.DrawImage(iir, 340, 220)
					dc.ResetClip()
					os.Remove(i.Filename)

					dc.SetRGB(0.9, 0.9, 0.9)
					dc.LoadFontFace("static/fonts/NotoSans-Bold.ttf", 25)
					dc.DrawString(t1.Artist.Name, 510, 280)

					dc.LoadFontFace("static/fonts/NotoSans-Regular.ttf", 25)
					dc.DrawString(t1.Name, 510, 320)
				}
			}

			dc.SavePNG("temp/" + ctx.Message.Author.ID.String() + ".png")      // save generated image
			r, _ := os.Open("temp/" + ctx.Message.Author.ID.String() + ".png") // open generated image into memory

			// create new message with the image + embed with a link to the user's Last.fm profile page
			ctx.Aurora.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
				Files: []disgord.CreateMessageFileParams{
					{FileName: ctx.Message.Author.ID.String() + ".png", Reader: r},
				},
				Embed: &disgord.Embed{
					Title: fmt.Sprintf("View %s's Profile on Last.fm", ctx.Message.Author.Username),
					URL:   fmt.Sprintf("https://last.fm/user/%s", user.Name),
					Color: 0x4d00ff,
				},
			})

			// Close and delete the image
			r.Close()
			os.Remove("temp/" + ctx.Message.Author.ID.String() + ".png")
		} else {
			ctx.Message.RespondString(ctx.Aurora, "You're currently not listening to anything") // Not currently playing anything
		}
	}

	return c.Command.CommandInterface
}

// getExt returns the extension of a given file name
func getExt(filename string) string {
	s := strings.Split(filename, ".")

	return s[len(s)-1]
}

// openImage returns an image.Image instance of a given file
func openImage(filename string) image.Image {
	in, _ := os.Open(filename)
	defer in.Close()
	var img image.Image

	switch getExt(filename) {
	case "png":
		img, _ = png.Decode(in)
		break
	case "jpeg":
		img, _ = jpeg.Decode(in)
		break
	case "jpg":
		img, _ = jpeg.Decode(in)
		break
	}

	return img
}
