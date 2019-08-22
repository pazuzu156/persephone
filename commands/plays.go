package commands

import (
	"fmt"
	"persephone/utils"

	"github.com/pazuzu156/aurora"
)

// Plays command.
type Plays struct{ Command }

// InitPlays initializes the plays command.
func InitPlays() Plays {
	return Plays{Init(&CommandItem{
		Name:        "plays",
		Description: "Gets the number of plays for the current playing track",
		Aliases:     []string{"p", "recent"},
		Usage:       "plays ...",
		Parameters:  []Parameter{},
	})}
}

// Register registers and runs the plays command.
func (c Plays) Register() *aurora.Command {
	c.CommandInterface.Run = func(ctx aurora.Context) {
		np, _ := utils.GetNowPlayingTrack(ctx.Message.Author, c.Lastfm)
		if len(ctx.Args) > 0 {
			fmt.Println(np)
		}
	}

	return c.CommandInterface
}
