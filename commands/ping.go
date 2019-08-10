package commands

import "github.com/pazuzu156/aurora"

// Ping is a simple testing command.
type Ping struct {
	Command Command
}

// InitPing initializes the ping command.
func InitPing(aliases ...string) Ping {
	return Ping{Init("ping", "Ping/Pong")}
}

// Register registers and runs the ping command.
func (c Ping) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		ctx.Message.RespondString(ctx.Aurora, "Pong")
	}

	return c.Command.CommandInterface
}
