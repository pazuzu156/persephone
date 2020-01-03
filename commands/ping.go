package commands

import (
	"fmt"

	"github.com/pazuzu156/atlas"
)

// Ping is a simple testing command.
type Ping struct{ Command }

// InitPing initializes the ping command.
func InitPing() Ping {
	return Ping{Init(&CommandItem{
		Name:        "ping",
		Description: "Ping/Pong",
		Usage:       "ping Pong!",
		Parameters: []Parameter{
			{
				Name:        "string",
				Description: "A string to send back to yourself",
				Required:    false,
			},
		},
	})}
}

// Register registers and runs the ping command.
func (c Ping) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		if len(ctx.Args) > 0 {
			fmt.Println(ctx.Message.Content)
			ctx.Message.Reply(ctx.Context, ctx.Atlas, ctx.Message.Content)
		} else {
			ctx.Message.Reply(ctx.Context, ctx.Atlas, "Pong")
		}
	}

	return c.CommandInterface
}
