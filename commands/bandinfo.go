package commands

import (
	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/aurora"
)

// Bandinfo command.
type Bandinfo struct{ Command }

// InitBandinfo initializes the bandinfo command.
func InitBandinfo() Bandinfo {
	return Bandinfo{Init(&CommandItem{
		Name:        "bandinfo",
		Description: "Gets information on the current playing band",
		Aliases:     []string{"bi"},
		Usage:       "bandinfo Darkthrone",
		Parameters:  []Parameter{},
	})}
}

// Register registers and runs the bandinfo command.
func (c Bandinfo) Register() *aurora.Command {
	c.CommandInterface.Run = func(ctx aurora.Context) {
		// ctx.Message.Reply(ctx.Aurora, "Hello, Bandinfo!")
		ctx.Aurora.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{})
	}

	return c.CommandInterface
}
