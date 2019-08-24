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

// Plays command.
type Plays struct{ Command }

// InitPlays initializes the plays command.
func InitPlays() Plays {
	return Plays{Init(&CommandItem{
		Name: "plays",
		Description: `Gets the number of plays for a given artist/album.
Giving no parameters will get the play count for the current playing track.
Passing no value to a parameter will get the plays for said parameter using the current playing track`,
		Aliases: []string{"p"},
		Usage:   "plays band:Grabak",
		Parameters: []Parameter{
			{
				Name:        "band",
				Value:       "name",
				Description: "Gets play count for a given artist",
				Required:    false,
			},
			{
				Name:        "album",
				Value:       "name:artist",
				Description: "Gets play count for a given album",
				Required:    false,
			},
		},
	})}
}

// Register registers and runs the plays command.
func (c Plays) Register() *aurora.Command {
	c.CommandInterface.Run = func(ctx aurora.Context) {
		np, _ := utils.GetNowPlayingTrack(ctx.Message.Author, c.Lastfm)
		leave := false

		if len(ctx.Args) > 0 {
			for _, arg := range ctx.Args {
				if strings.Contains(arg, ":") {
					arg = utils.JoinString(ctx.Args, " ")
					as := strings.Split(arg, ":")

					switch strings.ToLower(as[0]) {
					case "band": // show play count for requested artist
						a := as[1]
						artist, _ := c.Lastfm.Artist.GetInfo(lastfm.P{"artist": a, "username": database.GetUser(ctx.Message.Author).Lastfm})
						plays, _ := strconv.Atoi(artist.Stats.UserPlays)

						if plays > 0 {
							ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("**%s** has scrobbled %s **%d** times", ctx.Message.Author.Username, artist.Name, plays))
						} else {
							ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("**%s** has not scrobbled this artist yet", ctx.Message.Author.Username))
						}
						leave = true
						break
					case "album": // show play count for requested album
						if len(as) != 3 {
							ctx.Message.Reply(ctx.Aurora, "You need to provide all values to the album parameter!")
							leave = true
						} else {
							al := as[1]
							ar := as[2]
							album, _ := c.Lastfm.Album.GetInfo(lastfm.P{"artist": ar, "album": al, "username": database.GetUser(ctx.Message.Author).Lastfm})
							plays, _ := strconv.Atoi(album.UserPlayCount)

							if plays > 0 {
								ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("**%s** has scrobbled %s by %s **%d** times", ctx.Message.Author.Username, album.Name, album.Artist, plays))
							} else {
								ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("**%s** has not scrobbled this album yet", ctx.Message.Author.Username))
							}
						}
						leave = true
						break
					}

					// break from loop if we don't need to be here anymore
					if leave {
						break
					}
				} else {
					if strings.ToLower(ctx.Args[0]) == "band" { // show play count for current band
						artist, _ := c.Lastfm.Artist.GetInfo(lastfm.P{"artist": np.Artist.Name, "username": database.GetUser(ctx.Message.Author).Lastfm})
						plays, _ := strconv.Atoi(artist.Stats.UserPlays)

						if plays > 0 {
							ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("**%s** has scrobbled %s **%d** times", ctx.Message.Author.Username, artist.Name, plays))
						} else {
							ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("**%s** has not scrobbled this artist yet", ctx.Message.Author.Username))
						}
					} else if strings.ToLower(ctx.Args[0]) == "album" { // show play count for current album
						album, _ := c.Lastfm.Album.GetInfo(lastfm.P{"artist": np.Artist.Name, "album": np.Album.Name, "username": database.GetUser(ctx.Message.Author).Lastfm})
						plays, _ := strconv.Atoi(album.UserPlayCount)

						if plays > 0 {
							ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("**%s** has scrobbled %s by %s **%d** times", ctx.Message.Author.Username, album.Name, album.Artist, plays))
						} else {
							ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("**%s** has not scrobbled this album yet", ctx.Message.Author.Username))
						}
					}
				}
			}
		} else {
			// show play count for current playing track
			track, _ := c.Lastfm.Track.GetInfo(lastfm.P{"track": np.Name, "artist": np.Artist.Name, "username": database.GetUser(ctx.Message.Author).Lastfm})
			plays, _ := strconv.Atoi(track.UserPlayCount)

			if plays > 0 {
				ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("**%s** has scrobbled %s by %s **%d** times", ctx.Message.Author.Username, track.Name, track.Artist.Name, plays))
			} else {
				ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("**%s** has not scrobbled this track yet", ctx.Message.Author.Username))
			}
		}
	}

	return c.CommandInterface
}
