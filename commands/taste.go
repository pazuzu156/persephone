package commands

import (
	"github.com/pazuzu156/aurora"
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

// Register registers and runs the taste command.
func (c Taste) Register() *aurora.Command {
	c.CommandInterface.Run = func(ctx aurora.Context) {
		if len(ctx.Args) > 0 {
			// user, _ := lib.GetDiscordIDFromMention(ctx.Args[0])
		}

		ctx.Message.Reply(ctx.Aurora, "You need to supply a user to taste!")
	}

	return c.CommandInterface
}
