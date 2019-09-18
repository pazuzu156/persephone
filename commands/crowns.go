package commands

import (
	"fmt"
	"math"
	"persephone/database"
	"persephone/lib"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/naoina/genmai"
	"github.com/pazuzu156/atlas"
)

// Crowns command.
type Crowns struct{ Command }

// InitCrowns initializes the crowns command.
func InitCrowns() Crowns {
	return Crowns{Init(&CommandItem{
		Name:        "crowns",
		Description: "Shows a list of all your crowns (limit 10 per page)",
		Usage:       "crowns @Apollyon#6666",
		Parameters: []Parameter{
			{
				Name:        "member",
				Description: "Shows a list of crowns for the requested user",
				Required:    false,
			},
			{
				Name:        "page",
				Value:       "#",
				Description: "Shows the requested page of results",
				Required:    false,
			},
		},
	})}
}

// Register registers and runs the crowns command.
func (c Crowns) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		// check for command arguments
		if len(ctx.Args) > 0 {
			var (
				user *disgord.User
				page = 1
				err  error
			)

			for _, arg := range ctx.Args {
				if strings.HasPrefix(arg, "<@") {
					mention, _ := lib.GetDiscordIDFromMention(arg)
					user, _ = ctx.Atlas.GetUser(mention)
				}

				if strings.HasPrefix(arg, "page:") {
					a := strings.Split(arg, ":")

					if a[1] != "" {
						page, err = strconv.Atoi(a[1])

						if err != nil {
							ctx.Message.Reply(ctx.Atlas, "Invalid parameter passed to `page`")

							return
						}
					}
				}

				if user == nil {
					// TODO: bug with string usernames....
					// dbu := database.GetUserFromString(ctx.Args[0])

					// if dbu.Username != "" {
					// 	user, _ = ctx.Atlas.GetUser(dbu.GetDiscordID())
					// }

					user = ctx.Message.Author
				}

				// if user == nil {
				// 	dbu := database.GetUserFromString(arg)

				// 	if dbu.Username != "" {
				// 		user, _ = ctx.Atlas.GetUser(dbu.GetDiscordID())
				// 	} else {
				// 		user = ctx.Message.Author
				// 	}
				// }
				c.displayCrowns(ctx, user, page)
			}
		} else {
			c.displayCrowns(ctx, ctx.Message.Author, 1)
		}

	}

	return c.CommandInterface
}

// displayCrowns displays all crowns for users logged in with lastfm.
func (c Crowns) displayCrowns(ctx atlas.Context, user *disgord.User, page int) {
	if user == nil {
		ctx.Message.Reply(ctx.Atlas, "That username couldn't be found")

		return
	}

	crowns := database.GetUser(user).Crowns() // get crowns

	if len(crowns) > 0 {
		var (
			count      = len(crowns)
			maxPerPage = 10
			pages      = int(math.Ceil(float64(count) / float64(maxPerPage))) // gets total number of pages
			offset     = 0
		)

		// page sanity check
		if page <= pages {
			if page > 1 {
				offset = (page-1)*maxPerPage + 1 // fucking pagination
			}

			// query database with limit/offset for each page
			db, _ := database.OpenDB()
			db.Select(&crowns, db.Where("discord_id", "=", database.GetUInt64ID(user)).OrderBy("play_count", genmai.DESC), db.From(database.Crowns{}), db.Limit(maxPerPage).Offset(offset))

			// Sorts the slice in descending order by number of plays
			sort.SliceStable(crowns, func(i, j int) bool {
				return crowns[i].PlayCount > crowns[j].PlayCount
			})

			var descar []string

			// add each crown to string slice for embed
			for n, crown := range crowns {
				descar = append(descar, fmt.Sprintf("%d. 👑 %s with %d plays", n+1, crown.Artist, crown.PlayCount))
			}

			ctx.Atlas.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:       fmt.Sprintf("%d crowns for %s", count, user.Username),
					Description: lib.JoinString(descar, "\n"),
					Color:       lib.RandomColor(),
					Footer: &disgord.EmbedFooter{
						IconURL: lib.GenAvatarURL(c.getBotUser(ctx)),
						Text: fmt.Sprintf("Command invoked by: %s#%s | Page %d/%d",
							ctx.Message.Author.Username,
							ctx.Message.Author.Discriminator,
							page, pages,
						),
					},
					Timestamp: disgord.Time{
						Time: time.Now(),
					},
				},
			})

			return
		}

		ctx.Message.Reply(ctx.Atlas, fmt.Sprintf("%s Invalid page count", ctx.Message.Author.Mention()))

		return
	}

	ctx.Message.Reply(ctx.Atlas, "That user hasn't logged in to the bot yet.")
}
