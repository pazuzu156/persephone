package commands

import (
	"persephone/database"
	"sort"

	"github.com/pazuzu156/atlas"
)

// TopCrowns command.
type TopCrowns struct{ Command }

// InitTopCrowns initializes the topcrowns command.
func InitTopCrowns() TopCrowns {
	return TopCrowns{Init(&CommandItem{
		Name:        "topcrowns",
		Description: "simple crowns leaderboard",
	})}
}

// Register registers and runs the topcrowns command.
func (c TopCrowns) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		crowns := database.GetCrownsList()
		sort.SliceStable(crowns, func(i, j int) bool {
			return crowns[i].PlayCount > crowns[j].PlayCount
		})

	}

	return c.CommandInterface
}
