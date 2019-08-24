package commands

import (
	"fmt"
	"persephone/utils"
	"strings"

	"github.com/pazuzu156/aurora"
)

// Plays command.
type Plays struct{ Command }

// InitPlays initializes the plays command.
func InitPlays() Plays {
	return Plays{Init(&CommandItem{
		Name: "plays",
		Description: `Gets the number of plays for a given artist/album/track.
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
				Value:       "name",
				Description: "Gets play count for a given album",
				Required:    false,
			},
			{
				Name:        "track",
				Value:       "name",
				Description: "Gets play count for a given track (default uses current playing track)",
				Required:    false,
			},
		},
	})}
}

// Register registers and runs the plays command.
func (c Plays) Register() *aurora.Command {
	c.CommandInterface.Run = func(ctx aurora.Context) {
		// np, _ := utils.GetNowPlayingTrack(ctx.Message.Author, c.Lastfm)

		leave := false

		if len(ctx.Args) > 0 {
			for _, arg := range ctx.Args {
				if strings.Contains(arg, ":") {
					arg = utils.JoinString(ctx.Args, " ")
					as := strings.Split(arg, ":")

					switch strings.ToLower(as[0]) {
					case "band":
						artist := as[1]
						fmt.Println(artist)
						leave = true
					}

					// break from loop if we don't need to be here anymore
					if leave {
						break
					}

					// switch as[0] {
					// case "band":
					// 	break
					// }

					// if strings.ToLower(as[0]) == "band" {
					// 	// TODO: band
					// } else if strings.ToLower(as[0]) == "album" {
					// }
					// artist, err := c.Lastfm.Artist.GetInfo(lastfm.P{"artist": "artist", "username": database.GetUser(ctx.Message.Author).Lastfm})
				} else {
					// TODO: match argument
				}
			}
		}
	}

	return c.CommandInterface
}
