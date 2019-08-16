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
	"github.com/andersfylling/snowflake/v3"
	"github.com/pazuzu156/aurora"
)

// Crowns command.
type Crowns struct {
	Command Command
}

// InitCrowns initializes the crowns command.
func InitCrowns() Crowns {
	return Crowns{Init(
		"crowns",
		"List your crowns",
		[]UsageItem{
			{
				Command: "crowns",
			},
		},
	)}
}

// Register registers and runs the crowns command.
func (c Crowns) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		if len(ctx.Args) > 0 {
			var (
				user *disgord.User
				page = 1
				err  error
			)

			for _, arg := range ctx.Args {
				if strings.Contains(arg, "<@") {
					did, _ := strconv.Atoi(strings.TrimLeft(strings.TrimRight(arg, ">"), "<@"))
					discordid := snowflake.NewSnowflake(uint64(did))
					user, err = ctx.Aurora.GetUser(discordid)

					if err != nil {
						ctx.Message.Reply(ctx.Aurora, "That user could not be found in the server")
					}
				}

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

				c.displayCrowns(ctx, user, page)

			}
		} else {
			c.displayCrowns(ctx, ctx.Message.Author, 1)
		}

	}

	return c.Command.CommandInterface
}

func (c Crowns) displayCrowns(ctx aurora.Context, user *disgord.User, page int) {
	crowns := database.GetUser(user).Crowns()

	if len(crowns) > 0 {
		var (
			count      = len(crowns)
			maxPerPage = 10
			pages      = int(math.Ceil(float64(count) / float64(maxPerPage)))
			offset     = 0
		)

		if page <= pages {
			if page > 1 {
				offset = (page-1)*maxPerPage + 1
			}

			fmt.Println(offset)

			db, _ := database.OpenDB()
			db.Select(&crowns, db.Where("discord_id", "=", database.GetUInt64ID(user)), db.From(database.Crown{}), db.Limit(maxPerPage), db.Offset(offset))

			// Sorts the slice in descending order by number of plays
			sort.SliceStable(crowns, func(i, j int) bool {
				return crowns[i].PlayCount > crowns[j].PlayCount
			})

			var descar []string

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
