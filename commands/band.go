package commands

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// 	"persephone/database"
// 	"persephone/fm"
// 	"persephone/lib"
// 	"strconv"
// 	"strings"

// 	"github.com/andersfylling/disgord"
// 	"github.com/cavaliercoder/grab"
// 	"github.com/fogleman/gg"
// 	"github.com/nfnt/resize"
// 	"github.com/pazuzu156/atlas"
// 	"github.com/pazuzu156/lastfm-go"
// 	"golang.org/x/text/language"
// 	"golang.org/x/text/message"
// )

// // album and track positions for grids
// var (
// 	albumPositions = []fm.AlbumPosition{
// 		{
// 			X: 355,
// 			Y: 170,
// 			Shadow: fm.Shadow{
// 				X: 350,
// 				Y: 165,
// 				R: 10,
// 			},
// 			Info: fm.InfoText{
// 				X: 350,
// 				Y: 340,
// 				Plays: fm.PlaysText{
// 					X: 350,
// 					Y: 360,
// 				},
// 			},
// 		},
// 		{
// 			X: 555,
// 			Y: 170,
// 			Shadow: fm.Shadow{
// 				X: 550,
// 				Y: 165,
// 				R: 10,
// 			},
// 			Info: fm.InfoText{
// 				X: 550,
// 				Y: 340,
// 				Plays: fm.PlaysText{
// 					X: 550,
// 					Y: 360,
// 				},
// 			},
// 		},
// 		{
// 			X: 355,
// 			Y: 390,
// 			Shadow: fm.Shadow{
// 				X: 350,
// 				Y: 385,
// 				R: 10,
// 			},
// 			Info: fm.InfoText{
// 				X: 350,
// 				Y: 560,
// 				Plays: fm.PlaysText{
// 					X: 350,
// 					Y: 580,
// 				},
// 			},
// 		},
// 		{
// 			X: 555,
// 			Y: 390,
// 			Shadow: fm.Shadow{
// 				X: 550,
// 				Y: 385,
// 				R: 10,
// 			},
// 			Info: fm.InfoText{
// 				X: 550,
// 				Y: 560,
// 				Plays: fm.PlaysText{
// 					X: 550,
// 					Y: 580,
// 				},
// 			},
// 		},
// 	}
// 	trackPositions = []fm.TrackPosition{
// 		{
// 			X: 720,
// 			Y: 180,
// 			Plays: fm.PlaysText{
// 				X: 870,
// 				Y: 180,
// 			},
// 		},
// 		{
// 			X: 720,
// 			Y: 210,
// 			Plays: fm.PlaysText{
// 				X: 870,
// 				Y: 210,
// 			},
// 		},
// 		{
// 			X: 720,
// 			Y: 240,
// 			Plays: fm.PlaysText{
// 				X: 870,
// 				Y: 240,
// 			},
// 		},
// 		{
// 			X: 720,
// 			Y: 270,
// 			Plays: fm.PlaysText{
// 				X: 870,
// 				Y: 270,
// 			},
// 		},
// 	}
// )

// // Band command.
// type Band struct{ Command }

// // InitBand initializes the band command.
// func InitBand() Band {
// 	return Band{Init(&CommandItem{
// 		Name:        "band",
// 		Description: "Gets information on the artist you're currently listening to",
// 		Aliases:     []string{"b"},
// 		Usage:       "band Gorguts",
// 		Parameters: []Parameter{
// 			{
// 				Name:        "artist",
// 				Description: "Gets information on a requested artist",
// 				Required:    false,
// 			},
// 		},
// 	})}
// }

// // Register registers and runs the help command.
// func (c Band) Register() *atlas.Command {
// 	c.CommandInterface.Run = func(ctx atlas.Context) {
// 		// this command takes a really long time to complete
// 		// this message lets the user know that the bot is working
// 		tempmsg, _ := ctx.Message.Reply(ctx.Atlas, "Please wait while the artist image is generated...")
// 		defer ctx.Atlas.DeleteMessage(tempmsg.ChannelID, tempmsg.ID) // delete message when command completes

// 		// Want to check if an artist is supplied or not
// 		if len(ctx.Args) > 0 {
// 			artistName := strings.Trim(strings.Join(ctx.Args, " "), " ")
// 			artist, err := c.getArtistInfo(artistName, ctx.Message.Author)

// 			if err != nil {
// 				ctx.Message.Reply(ctx.Atlas, err.Error())

// 				return
// 			}

// 			c.displayArtistInfo(ctx, artist) // display info with requested artist
// 		} else {
// 			// current track should have the artist info we need to do a new artist query
// 			track, err := fm.GetNowPlayingTrack(ctx.Message.Author, c.Lastfm)

// 			if err != nil {
// 				ctx.Message.Reply(ctx.Atlas, err.Error())

// 				return
// 			}

// 			artist, err := c.getArtistInfo(track.Artist.Name, ctx.Message.Author) // get full artist info

// 			if err != nil {
// 				ctx.Message.Reply(ctx.Atlas, "Couldn't find that artist")
// 			}

// 			c.displayArtistInfo(ctx, artist) // display info with current artist
// 		}
// 	}

// 	return c.CommandInterface
// }

// func (c Band) displayArtistInfo(ctx atlas.Context, artist lastfm.ArtistGetInfo) {
// 	albums := c.getAlbumsList(ctx, artist)      // gets users albums from artist
// 	tracks, err := c.getTracksList(ctx, artist) // gets users tracks from artist

// 	if err != nil {
// 		ctx.Message.Reply(ctx.Atlas, err.Error())

// 		return
// 	}

// 	lfmuser, _ := database.GetLastfmUserInfo(ctx.Message.Author, c.Lastfm)

// 	aimg := lib.GetArtistImage(artist) // artist image is scraped from metal-archives
// 	avres, _ := grab.Get(lib.LocGet("temp/"), lib.GenAvatarURL(ctx.Message.Author))
// 	bg := lib.OpenImage(lib.LocGet("static/images/background.png"))
// 	av := lib.OpenImage(avres.Filename)
// 	os.Remove(avres.Filename)

// 	air := resize.Resize(230, 230, aimg, resize.Bicubic)
// 	avr := resize.Resize(72, 72, av, resize.Bicubic)

// 	dc := gg.NewContext(1000, 600)
// 	dc.DrawImage(bg, 0, 0)

// 	// artist image shadow
// 	dc.SetRGBA(1, 1, 1, 0.2)
// 	dc.DrawRectangle(0, 50, 1000, 72)
// 	dc.Fill()

// 	// artist image
// 	dc.SetRGBA(0, 0, 0, 0.3)
// 	dc.DrawRoundedRectangle(50, 50, 240, 240, 10)
// 	dc.Fill()
// 	dc.DrawImage(air, 55, 55)

// 	// artist name and play count
// 	dc.SetRGB(0.9, 0.9, 0.9)
// 	dc.LoadFontFace(FontBold, 20)
// 	dc.DrawStringWrapped(artist.Name, 50, 310, 0, 0, 230, 1.5, gg.AlignCenter)
// 	dc.LoadFontFace(FontRegular, 20)
// 	dc.DrawStringWrapped(fmt.Sprintf("%s plays", artist.Stats.UserPlays), 50, 345, 0, 0, 235, 1.5, gg.AlignCenter)

// 	// separator between artist name and tags
// 	dc.DrawLine(50, 370, 285, 370)
// 	dc.SetLineWidth(0.5)
// 	dc.Stroke()

// 	// Get the artist tags, and stringify them
// 	var tags []string

// 	for _, tag := range artist.Tags {
// 		tags = append(tags, tag.Name)
// 	}

// 	// tags
// 	dc.DrawStringWrapped(lib.JoinString(tags, ", "), 50, 380, 0, 0, 235, 1.5, gg.AlignCenter)

// 	// user avatar/info
// 	dc.DrawImage(avr, 315, 50)
// 	dc.LoadFontFace(FontBold, 26)
// 	dc.SetRGB(0.9, 0.9, 0.9)
// 	dc.DrawString(ctx.Message.Author.Username+" ("+lfmuser.Name+")", 400, 80)
// 	// scrobble count
// 	dc.LoadFontFace(FontRegular, 20)
// 	dc.SetRGB(0.9, 0.9, 0.9)
// 	printer := message.NewPrinter(language.English)
// 	pc, _ := strconv.Atoi(lfmuser.PlayCount)
// 	dc.DrawString(fmt.Sprintf("%s scrobbles", printer.Sprintf("%d", pc)), 400, 110)

// 	dc.DrawString("Albums", 490, 150)

// 	// takes all albums and aranges them in a 2x2 grid
// 	for i, album := range albums {
// 		if i < len(albums) && i < 4 {
// 			ares, _ := grab.Get(lib.LocGet("temp/"), album.Images[3].URL)
// 			ai := lib.OpenImage(ares.Filename)
// 			os.Remove(ares.Filename)
// 			ar := resize.Resize(145, 145, ai, resize.Bicubic)
// 			pos := albumPositions[i]

// 			// shadow
// 			dc.SetRGBA(0, 0, 0, 0.3)
// 			dc.DrawRoundedRectangle(pos.Shadow.X, pos.Shadow.Y, 155, 155, pos.Shadow.R)
// 			dc.Fill()

// 			// album image
// 			dc.DrawImage(ar, pos.X, pos.Y)

// 			// album name/play count
// 			dc.SetRGBA(1, 1, 1, 0.9)
// 			dc.LoadFontFace(FontRegular, 20)
// 			dc.DrawString(lib.ShortStr(album.Name, 15), pos.Info.X, pos.Info.Y)
// 			dc.LoadFontFace(FontRegular, 16)
// 			dc.DrawString(fmt.Sprintf("%s plays", album.PlayCount), pos.Info.Plays.X, pos.Info.Plays.Y)
// 		}
// 	}

// 	dc.LoadFontFace(FontRegular, 20)
// 	dc.DrawString("Tracks", 790, 150)

// 	// takes top tracks and lists the top 4
// 	for i, track := range tracks {
// 		if i < len(tracks) && i < 4 {
// 			pos := trackPositions[i]
// 			dc.SetRGB(0.9, 0.9, 0.9)
// 			dc.LoadFontFace(FontRegular, 16)
// 			dc.DrawString(lib.ShortStr(track.Name, 15), pos.X, pos.Y)
// 			dc.LoadFontFace(FontBold, 16)
// 			dc.DrawString(fmt.Sprintf("%s plays", track.PlayCount), pos.Plays.X, pos.Plays.Y)
// 		}
// 	}

// 	lib.BrandImage(dc) // brand image

// 	dc.SavePNG(lib.LocGet("temp/" + ctx.Message.Author.ID.String() + "_band.png"))
// 	r, _ := os.Open(lib.LocGet("temp/" + ctx.Message.Author.ID.String() + "_band.png"))

// 	ctx.Atlas.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
// 		Files: []disgord.CreateMessageFileParams{
// 			{
// 				FileName: r.Name(),
// 				Reader:   r,
// 			},
// 		},
// 	})

// 	r.Close()
// 	os.Remove(lib.LocGet("temp/" + ctx.Message.Author.ID.String() + "_band.png"))
// }

// // getArtistInfo retrieves artist info for a given user.
// func (c Band) getArtistInfo(artist string, user *disgord.User) (lastfm.ArtistGetInfo, error) {
// 	dbu := database.GetUser(user)
// 	return c.Lastfm.Artist.GetInfo(lastfm.P{"artist": artist, "username": dbu.Lastfm})
// }

// // getAlbumsList gets albums for a user for a given artist.
// func (c Band) getAlbumsList(ctx atlas.Context, artist lastfm.ArtistGetInfo) []fm.TopAlbum {
// 	user := database.GetUser(ctx.Message.Author)
// 	// kinda gotta get as many albums as possible. we'll use totalPages to make more queries if need be
// 	alist, _ := c.Lastfm.User.GetTopAlbums(lastfm.P{"user": user.Lastfm, "limit": "1000"}) // limit max = 1000
// 	var albums = []fm.TopAlbum{}

// 	// add first batch of albums for artist into slice
// 	for _, album := range alist.Albums {
// 		if album.Artist.Name == artist.Name {
// 			albums = append(albums, album)
// 		}
// 	}

// 	// if more pages than 1, run another sweep
// 	if alist.TotalPages > 1 {
// 		for i := 1; i <= alist.TotalPages; i++ {
// 			al, _ := c.Lastfm.User.GetTopAlbums(lastfm.P{"user": user.Lastfm, "limit": "1000", "page": strconv.Itoa(i)})

// 			// add more albums for artist into slice
// 			if al.Albums[i].Artist.Name == artist.Name {
// 				albums = append(albums, al.Albums[i])
// 			}
// 		}
// 	}

// 	return albums
// }

// // getTracksList gets the users top tracks for a given artist.
// func (c Band) getTracksList(ctx atlas.Context, artist lastfm.ArtistGetInfo) ([]fm.TopTrack, error) {
// 	// this method works like getAlbumsList, won't comment
// 	user := database.GetUser(ctx.Message.Author)
// 	tlist, _ := c.Lastfm.User.GetTopTracks(lastfm.P{"user": user.Lastfm, "limit": "1000"})
// 	var tracks = []fm.TopTrack{}

// 	for _, track := range tlist.Tracks {
// 		if track.Artist.Name == artist.Name {
// 			tracks = append(tracks, track)
// 		}
// 	}

// 	if tlist.TotalPages > 1 {
// 		for i := 1; i <= tlist.TotalPages; i++ {
// 			tl, err := c.Lastfm.User.GetTopTracks(lastfm.P{"user": user.Lastfm, "limit": "1000", "page": strconv.Itoa(i)})

// 			if err != nil {
// 				return nil, errors.New("An error occurred while trying to retireve this info. Please try again later")
// 			}

// 			if tl.Tracks[i].Artist.Name == artist.Name {
// 				tracks = append(tracks, tl.Tracks[i])
// 			}
// 		}
// 	}

// 	return tracks, nil
// }
