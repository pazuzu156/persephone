package commands

import (
	"fmt"
	"math"
	"persephone/database"
	"persephone/utils"
	"sort"
	"strconv"
	"strings"

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
			// TODO: requested user
			var (
				user *disgord.User
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
					fmt.Println(user)
				}

				if strings.Contains(arg, "page:") {
					var (
						page = 1
						err  error
					)
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

			}
			// if strings.Contains(ctx.Args[0], "<@") {
			// 	did, _ := strconv.Atoi(strings.TrimLeft(strings.TrimRight(ctx.Args[0], ">"), "<@"))
			// 	discordid := snowflake.NewSnowflake(uint64(did))
			// 	user, err = ctx.Aurora.GetUser(discordid)

			// 	if err != nil {
			// 		ctx.Message.Reply(ctx.Aurora, "That user could not be found in the server")
			// 	}

			// 	c.displayCrowns(ctx, user)
			// } else {

			// }
		} else {
			c.displayCrowns(ctx, ctx.Message.Author, 1)
			// ctx.Message.Reply(ctx.Aurora, utils.JoinString(descar, "\n"))

		}

		// dbu := database.GetUser(ctx.Message.Author)

		// sql := db.DB()
		// stmt, _ := sql.Prepare("SELECT * FROM crown LIMIT 1 OFFSET 0")
		// res, _ := stmt.Exec()
		// var desc = ""
		// dbu.DB().Select(&out, dbu.DB().From(database.User{}))

		// ctx.Message.Reply(ctx.Aurora, dbu.Crowns())
		// fmt.Println(out)

	}

	return c.Command.CommandInterface
}

func (c Crowns) displayCrowns(ctx aurora.Context, user *disgord.User, page int) {
	crowns := database.GetUser(user).Crowns()

	if len(crowns) > 0 {
		count := len(crowns)
		maxPerPage := 5
		pages := math.Ceil(float64(count) / float64(maxPerPage))
		page := page
		ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("Page %d of %d", page, int(pages)))

		db, _ := database.OpenDB()
		db.Select(&crowns, db.From(database.Crown{}), db.Limit(maxPerPage), db.Offset(0))

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
			},
		})
	}
}
