package commands

import (
	"fmt"
	"math"
	"persephone/database"
	"persephone/utils"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/aurora"
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
func (c Crowns) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		// check for command arguments
		if len(ctx.Args) > 0 {
			var (
				user *disgord.User
				page = 1
				err  error
			)

			// loop through arguments
			// they don't have a particular order
			for _, arg := range ctx.Args {
				// Check if a user is supplied
				if strings.Contains(arg, "<@") {
					did, _ := strconv.Atoi(strings.TrimLeft(strings.TrimRight(arg, ">"), "<@"))
					discordid := disgord.NewSnowflake(uint64(did))
					user, err = ctx.Aurora.GetUser(discordid)

					if err != nil {
						ctx.Message.Reply(ctx.Aurora, "That user could not be found in the server")
					}
				}

				// check if a page number is requested
				if strings.Contains(arg, "page:") {
					a := strings.Split(arg, ":")

					if a[1] != "" {
						page, err = strconv.Atoi(a[1])

						if err != nil {
							ctx.Message.Reply(ctx.Aurora, "Invalid parameter passed to `page`")

							return
						}
					}

					fmt.Println(page)
				}

				if user == nil {
					user = ctx.Message.Author
				}

				c.displayCrowns(ctx, user, page) // display crowns

			}
		} else {
			c.displayCrowns(ctx, ctx.Message.Author, 1) // display crowns
		}

	}

	return c.Command.CommandInterface
}

// displayCrowns displays all crowns for users logged in with lastfm.
func (c Crowns) displayCrowns(ctx aurora.Context, user *disgord.User, page int) {
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

			fmt.Println(offset)

			// query database with limit/offset for each page
			db, _ := database.OpenDB()
			db.Select(&crowns, db.Where("discord_id", "=", database.GetUInt64ID(user)), db.From(database.Crown{}), db.Limit(maxPerPage), db.Offset(offset))

			// Sorts the slice in descending order by number of plays
			sort.SliceStable(crowns, func(i, j int) bool {
				return crowns[i].PlayCount > crowns[j].PlayCount
			})

			var descar []string

			// add each crown to string slice for embed
			for n, crown := range crowns {
				descar = append(descar, fmt.Sprintf("%d. 👑 %s with %d plays", n+1, crown.Artist, crown.PlayCount))
			}

			ctx.Aurora.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:       fmt.Sprintf("%d crowns for %s", count, user.Username),
					Description: utils.JoinString(descar, "\n"),
					Color:       utils.RandomColor(),
					Footer: &disgord.EmbedFooter{
						IconURL: utils.GenAvatarURL(utils.GetBotUser(ctx)),
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

		ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("%s Invalid page count", ctx.Message.Author.Mention()))
	}
}
