package commands

import "github.com/pazuzu156/aurora"

// Plays command.
type Plays struct{ Command }

// InitPlays initializes the plays command.
func InitPlays() Plays {
    return Plays{Init(&CommandItem{
        Name: "plays",
        Description: "Gets the number of plays for the current playing track",
        Aliases: []string{},
        Usage: "plays ...",
        Parameters: []Parameter{},
    })}
}

// Register registers and runs the plays command.
func (c Plays) Register() *aurora.Command {
    c.Command.CommandInterface.Run = func(ctx aurora.Context) {
        ctx.Message.Reply(ctx.Aurora, "Hello, Plays!")
    }

    return c.Command.CommandInterface
}
