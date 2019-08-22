package commands

import (
	"fmt"

	"github.com/pazuzu156/aurora"
)

// Ping is a simple testing command.
type Ping struct{ Command }

// InitPing initializes the ping command.
func InitPing() Ping {
	return Ping{InitCmd(&CommandItem2{
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
func (c Ping) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		if len(ctx.Args) > 0 {
			fmt.Println(ctx.Message.Content)
			ctx.Message.Reply(ctx.Aurora, ctx.Message.Content)
		} else {
			ctx.Message.Reply(ctx.Aurora, "Pong")
		}
	}

	return c.Command.CommandInterface
}
