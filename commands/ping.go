package commands

import "github.com/polaron/aurora"

// Ping is a simple testing command.
type Ping struct {
	Command Command
}

// InitPing initializes the ping command.
func InitPing(aliases ...string) Ping {
	return Ping{Init("ping", "Ping/Pong", aliases...)}
}

// Register registers and runs the ping command.
func (c Ping) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		ctx.Message.RespondString(ctx.Aurora, "Pong")
	}

	return c.Command.CommandInterface
}
