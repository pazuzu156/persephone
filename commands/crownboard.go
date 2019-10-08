package commands

import (
	"fmt"
	"persephone/database"
	"persephone/lib"
	"sort"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
)

// CrownBoard command.
type CrownBoard struct{ Command }

// InitCrownBoard initializes the crownboard command.
func InitCrownBoard() CrownBoard {
	return CrownBoard{Init(&CommandItem{
		Name:        "crownboard",
		Description: "Crowns leaderboard",
		Aliases:     []string{"cb", "leaders"},
	})}
}

// Register registers and runs the crownboard command.
func (c CrownBoard) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		users := database.GetUsers()
		crownsCount := len(database.GetCrownsList())

		sort.SliceStable(users, func(i, j int) bool {
			return len(users[i].Crowns()) > len(users[j].Crowns())
		})

		var (
			descar []string
			limit  = 10
		)

		for n, user := range users {
			if n < limit {
				descar = append(descar, fmt.Sprintf("%d. 👑 %s with %d crowns", n+1, user.Username, len(user.Crowns())))
			}

			n++
		}

		f, t := c.embedFooter(ctx)

		ctx.Atlas.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				// Title:       "Crown Leaderboards",
				Title:       fmt.Sprintf("Crowns Leaderboards • %s total crowns", lib.HumanNumber(crownsCount)),
				Description: lib.JoinString(descar, "\n"),
				Color:       lib.RandomColor(),
				Footer:      f, Timestamp: t,
			},
		})
	}

	return c.CommandInterface
}
