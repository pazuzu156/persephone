package commands

import (
	"fmt"
	"persephone/database"
	"persephone/fm"
	"persephone/lib"
	"strconv"
	"strings"

	"github.com/pazuzu156/atlas"
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
		Usage:   "plays --artist Grabak",
		Parameters: []Parameter{
			{
				Name:        "artist",
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
func (c Plays) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		np, _ := fm.GetNowPlayingTrack(ctx.Message.Author, c.Lastfm)

		if len(ctx.Args) > 0 {
			for n, arg := range ctx.Args {
				if strings.HasPrefix(arg, "--") {
					arg = strings.TrimLeft(arg, "--")
					argv, isset := ctx.Args[n+1]

					switch strings.ToLower(arg) {
					case "album": // album
						var album lastfm.AlbumGetInfo

						if isset {
							delete(ctx.Args, 0)
							argvs := strings.Split(lib.JoinStringMap(ctx.Args, " "), ":")

							if len(argvs) > 1 {
								al := argvs[0]
								ar := argvs[1]
								album, _ = c.Lastfm.Album.GetInfo(lastfm.P{"artist": ar, "album": al, "username": c.getLastfmUser(ctx.Message.Author)})
							} else {
								ctx.Message.Reply(ctx.Atlas, "Invalid argument syntax. The argument value should look like: album:artist")

								break
							}
						} else {
							album, _ = c.Lastfm.Album.GetInfo(lastfm.P{"artist": np.Artist.Name, "album": np.Album.Name, "username": c.getLastfmUser(ctx.Message.Author)})
						}

						plays, _ := strconv.Atoi(album.UserPlayCount)

						if plays > 0 {
							ctx.Message.Reply(ctx.Atlas, fmt.Sprintf("**%s** has scrobbled %s by %s **%d** times", ctx.Message.Author.Username, album.Name, album.Artist, plays))
						} else {
							ctx.Message.Reply(ctx.Atlas, fmt.Sprintf("**%s** has not scrobbled this album yet", ctx.Message.Author.Username))
						}
						break
					case "artist": // artist
						var artist lastfm.ArtistGetInfo

						if isset {
							artist, _ = c.Lastfm.Artist.GetInfo(lastfm.P{"artist": argv, "username": c.getLastfmUser(ctx.Message.Author)})
						} else {
							artist, _ = c.Lastfm.Artist.GetInfo(lastfm.P{"artist": np.Artist.Name, "username": c.getLastfmUser(ctx.Message.Author)})
						}

						plays, _ := strconv.Atoi(artist.Stats.UserPlays)

						if plays > 0 {
							ctx.Message.Reply(ctx.Atlas, fmt.Sprintf("**%s** has scrobbled %s **%d** times", ctx.Message.Author.Username, artist.Name, plays))
						} else {
							ctx.Message.Reply(ctx.Atlas, fmt.Sprintf("**%s** has not scrobbled this artist yet", ctx.Message.Author.Username))
						}
						break
					}
				}
			}
		} else {
			// show play count for current playing track
			track, _ := c.Lastfm.Track.GetInfo(lastfm.P{"track": np.Name, "artist": np.Artist.Name, "username": database.GetUser(ctx.Message.Author).Lastfm})
			plays, _ := strconv.Atoi(track.UserPlayCount)

			if plays > 0 {
				ctx.Message.Reply(ctx.Atlas, fmt.Sprintf("**%s** has scrobbled %s by %s **%d** times", ctx.Message.Author.Username, track.Name, track.Artist.Name, plays))
			} else {
				ctx.Message.Reply(ctx.Atlas, fmt.Sprintf("**%s** has not scrobbled this track yet", ctx.Message.Author.Username))
			}
		}
	}

	return c.CommandInterface
}
