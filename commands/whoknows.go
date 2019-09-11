package commands

import (
	"fmt"
	"persephone/database"
	"persephone/lib"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/aurora"
	"github.com/pazuzu156/lastfm-go"
)

// Whoknows command.
type Whoknows struct{ Command }

// InitWhoknows initialized the whoknows command.
func InitWhoknows() Whoknows {
	return Whoknows{Init(&CommandItem{
		Name:        "whoknows",
		Description: "Shows who knows a specific artist",
		Aliases:     []string{"wk"},
		Usage:       "whoknows Judas Iscariot",
		Parameters: []Parameter{
			{
				Name:        "artist",
				Description: "shows a list of users who know the requested artist",
				Required:    false,
			},
		},
	})}
}

// Register registers and runs the whoknows command.
func (c Whoknows) Register() *aurora.Command {
	c.CommandInterface.Run = func(ctx aurora.Context) {
		// if args > 0, an artist is likely provided
		// so .wk <artist> runs the command on a requested artist
		if len(ctx.Args) > 0 {
			artist := strings.TrimRight(strings.Join(ctx.Args, " "), " ")
			a, err := c.Lastfm.Artist.GetInfo(lastfm.P{"artist": artist})

			if err != nil {
				ctx.Message.Reply(ctx.Aurora, "Artist could not be found on Last.fm")

				return
			}

			go c.displayWhoKnows(ctx, a) // runs the whoknows logic and displays the embed
		} else {
			// get the user from the sender
			if user := database.GetUser(ctx.Message.Author); user.Username != "" {
				np, err := c.Lastfm.User.GetRecentTracks(lastfm.P{"user": user.Lastfm, "limit": "2"})

				if err != nil {
					ctx.Message.Reply(ctx.Aurora, "Artist could not be found on Last.fm")

					return
				}

				// Loop through their recent tracks
				for _, track := range np.Tracks {
					if track.NowPlaying == "true" {
						npa, _ := c.Lastfm.Artist.GetInfo(lastfm.P{"artist": track.Artist.Name})
						go c.displayWhoKnows(ctx, npa) // runs the whoknows logic and displays the embed
						break
					}
				}
			} else {
				ctx.Message.Reply(ctx.Aurora, "You're not currently logged in with Last.fm")
			}
		}
	}

	return c.CommandInterface
}

// displayWhoKnows displays an embed with a list of top users who have scrobbled a given artist.
func (c Whoknows) displayWhoKnows(ctx aurora.Context, artist lastfm.ArtistGetInfo) {
	users := database.GetUsers()

	// user representation for the wk slice
	type U struct {
		DiscordID uint64
		Name      string
		Plays     int
	}

	// Gets all logged in users
	var wk = []U{}
	for _, user := range users {
		a, _ := c.Lastfm.Artist.GetInfo(lastfm.P{"username": user.Lastfm, "artist": artist.Name})
		plays, _ := strconv.Atoi(a.Stats.UserPlays)

		// add all users who have scrobbled the artist to the slice
		if plays > 0 {
			wk = append(wk, U{DiscordID: user.DiscordID, Name: user.Username, Plays: plays})
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
					db, _ := database.OpenDB()
					defer db.Close()
					crowns := database.GetCrownsList()
					now := time.Now()

					updated := false
					for _, crown := range crowns {
						if artist.Name == crown.Artist {
							updated = true
							crown.DiscordID = user.DiscordID
							crown.PlayCount = user.Plays
							crown.Time.UpdatedAt = &now
							db.Update(crown)
						}
					}

					if !updated {
						crown := []database.Crowns{
							{
								DiscordID: database.GetUInt64ID(ctx.Message.Author),
								Artist:    artist.Name,
								PlayCount: user.Plays,
								Time: database.Time{
									CreatedAt: &now,
									UpdatedAt: &now,
								},
							},
						}
						db.Insert(crown)
					}

					desc += fmt.Sprintf("👑 **%s** with **%d** plays\n", user.Name, user.Plays) // has the most plays
				} else {
					desc += fmt.Sprintf("🎶 **%s** with **%d** plays\n", user.Name, user.Plays) // all other scrobblers
				}
			}
		}

		f, t := c.embedFooter(ctx)
		ctx.Aurora.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       fmt.Sprintf("Who knows %s?", artist.Name),
				URL:         fmt.Sprintf("https://last.fm/music/%s", strings.Replace(artist.Name, " ", "+", len(artist.Name))),
				Description: desc,
				Color:       lib.RandomColor(),
				Footer:      f, Timestamp: t,
			},
		})
	} else {
		ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("No one has scrobbled %s yet", artist.Name))
	}
}
