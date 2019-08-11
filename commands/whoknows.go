package commands

import (
	"fmt"
	"persephone/database"
	"sort"
	"strconv"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/aurora"
	"github.com/pazuzu156/lastfm-go"
)

// Whoknows command.
type Whoknows struct {
	Command Command
}

// InitWhoknows initialized the whoknows command.
func InitWhoknows(aliases ...string) Whoknows {
	return Whoknows{Init("whoknows",
		"Shows who knows a specific artist",
		[]string{
			"whoknows",
			"whoknows [artist]",
		},
		aliases...,
	)}
}

// Register registers and runs the whoknows command.
func (c Whoknows) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		// if args > 0, an artist is likely provided
		// so .wk <artist> runs the command on a requested artist
		if len(ctx.Args) > 0 {
			var artist = ""

			// merge all args into one string
			if len(ctx.Args) > 1 {
				artist = strings.Join(ctx.Args, " ")
			} else {
				artist = ctx.Args[0]
			}

			a, err := c.Command.Lastfm.Artist.GetInfo(lastfm.P{"artist": artist})

			if err != nil {
				ctx.Message.RespondString(ctx.Aurora, "Artist could not be found on Last.fm")

				return
			}

			c.displayWhoKnows(ctx, a) // runs the whoknows logic and displays the embed
		} else {
			// get the user from the sender
			if user := database.GetUser(ctx.Message.Author); user.Username != "" {
				np, err := c.Command.Lastfm.User.GetRecentTracks(lastfm.P{"user": user.Lastfm, "limit": "2"})

				if err != nil {
					ctx.Message.RespondString(ctx.Aurora, "Artist could not be found on Last.fm")

					return
				}

				// Loop through their recent tracks
				for _, track := range np.Tracks {
					if track.NowPlaying == "true" {
						npa, _ := c.Command.Lastfm.Artist.GetInfo(lastfm.P{"artist": track.Artist.Name})
						c.displayWhoKnows(ctx, npa) // runs the whoknows logic and displays the embed
						break
					}
				}
			} else {
				ctx.Message.RespondString(ctx.Aurora, "You're not currently logged in with Last.fm")
			}
		}
	}

	return c.Command.CommandInterface
}

func (c Whoknows) displayWhoKnows(ctx aurora.Context, artist lastfm.ArtistGetInfo) {
	users := database.GetUsers()

	// user representation for the wk slice
	type U struct {
		Name  string
		Plays int
	}

	// Gets all logged in users
	var wk = []U{}
	for _, user := range users {
		a, _ := c.Command.Lastfm.Artist.GetInfo(lastfm.P{"username": user.Lastfm, "artist": artist.Name})
		plays, _ := strconv.Atoi(a.Stats.UserPlays)

		// add all users who have scrobbled the artist to the slice
		// otherwise, break from the loop
		// Also, don't loop through all the logged in users if the last
		// seen user has no scrobbles
		if plays > 0 {
			wk = append(wk, U{Name: user.Username, Plays: plays})
		} else if plays == 0 {
			break
		}
	}

	// Did we actually get any results?
	// If so, display them
	if len(wk) > 0 {
		// Sorts the slice in descending order by number of plays
		sort.SliceStable(wk, func(i, j int) bool {
			return wk[i].Plays > wk[j].Plays
		})

		var max = 10 // display a max of 10 users
		var desc = fmt.Sprintf("%d users have scrobbled %s\n", len(wk), artist.Name)
		for i := 0; i < len(wk); i++ {
			if i < max {
				user := wk[i]
				if i == 0 {
					desc += fmt.Sprintf("👑 **%s** with **%d** plays\n", user.Name, user.Plays) // has the most plays
				} else {
					desc += fmt.Sprintf("🎶 **%s** with **%d** plays\n", user.Name, user.Plays) // all other scrobblers
				}
			}
		}

		ctx.Aurora.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       fmt.Sprintf("Who knows %s?", artist.Name),
				URL:         fmt.Sprintf("https://last.fm/music/%s", artist.Name),
				Description: desc,
				Color:       0xb3ff00,
			},
		})
	} else {
		ctx.Message.RespondString(ctx.Aurora, fmt.Sprintf("No one has scrobbled %s yet", artist.Name))
	}
}
