package commands

import (
	"fmt"
	"sort"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
	"github.com/pazuzu156/lastfm-go"
	"github.com/pazuzu156/persephone/fm"
	"github.com/pazuzu156/persephone/lib"
)

// Taste command.
type Taste struct{ Command }

// InitTaste initializes the taste command.
func InitTaste() Taste {
	return Taste{Init(&CommandItem{
		Name:        "taste",
		Description: "Stack your musical tastes up with others",
		Aliases:     []string{},
		Usage:       "taste ...",
		Parameters:  []Parameter{},
	})}
}

// MatchData holds data of matched artists for caller and recipient.
type MatchData struct {
	UserArtistData   fm.Artist
	AuthorArtistData fm.Artist
}

// Register registers and runs the taste command.
func (c Taste) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		if len(ctx.Args) > 0 {
			var (
				user       *disgord.User
				callLimit  = "700"
				matches    []MatchData
				matchLimit = 10
				descar     []*disgord.EmbedField
			)

			if strings.HasPrefix(ctx.Args[0], "<@") {
				mention, _ := lib.GetDiscordIDFromMention(ctx.Args[0])
				user, _ = ctx.Atlas.GetUser(ctx.Context, mention)
			}

			if user == nil {
				// TODO: bug with string usernames....
				// dbu := lib.GetUserFromString(ctx.Args[0])

				// if dbu.Username != "" {
				// 	user, _ = ctx.Atlas.GetUser(dbu.GetDiscordID())
				// }

				ctx.Message.Reply(ctx.Context, ctx.Atlas, "You need to supply a user to taste!")

				return
			}

			if user.ID == ctx.Message.Author.ID {
				ctx.Message.Reply(ctx.Context, ctx.Atlas, "You cannot taste yourself")

				return
			}

			dba := lib.GetUser(ctx.Message.Author)
			dbu := lib.GetUser(user)
			authorData, _ := c.Lastfm.User.GetTopArtists(lastfm.P{"user": dba.Lastfm, "period": "overall", "limit": callLimit})
			userData, _ := c.Lastfm.User.GetTopArtists(lastfm.P{"user": dbu.Lastfm, "period": "overall", "limit": callLimit})

			for _, x := range userData.Artists {
				b, a := c.contains(authorData.Artists, x.Name)

				if b {
					matches = append(matches, MatchData{
						UserArtistData:   x,
						AuthorArtistData: a,
					})
				}
			}

			if len(matches) == 0 {
				ctx.Message.Reply(ctx.Context, ctx.Atlas, fmt.Sprintf("You and %s share no common artists", dbu.Username))

				return
			}

			sort.SliceStable(matches, func(i, j int) bool {
				return matches[i].UserArtistData.PlayCount < matches[j].UserArtistData.PlayCount
			})

			for n, match := range matches {
				if n < matchLimit {
					descar = append(descar, &disgord.EmbedField{
						Name: match.UserArtistData.Name,
						Value: fmt.Sprintf("%s plays - %s plays",
							match.AuthorArtistData.PlayCount,
							match.UserArtistData.PlayCount,
						),
						Inline: true,
					})
				}

				n++
			}

			f, t := lib.AddEmbedFooter(ctx.Message)

			ctx.Atlas.CreateMessage(ctx.Context, ctx.Message.ChannelID, &disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title: fmt.Sprintf("%s and %s taste comparison",
						ctx.Message.Author.Username,
						user.Username,
					),
					Thumbnail: &disgord.EmbedThumbnail{
						URL: lib.GenAvatarURL(ctx.Message.Author),
					},
					Fields: descar,
					Color:  lib.RandomColor(),
					Footer: f, Timestamp: t,
				},
			})
		} else {
			ctx.Message.Reply(ctx.Context, ctx.Atlas, "You need to supply a user to taste!")
		}
	}

	return c.CommandInterface
}

func (c Taste) contains(a fm.Artists, x string) (bool, fm.Artist) {
	for _, n := range a {
		if x == n.Name {
			return true, n
		}
	}

	return false, fm.Artist{}
}
