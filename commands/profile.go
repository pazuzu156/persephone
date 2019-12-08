package commands

import "github.com/pazuzu156/atlas"

// Profile command.
type Profile struct{ Command }

// InitProfile initializes the profile command.
func InitProfile() Profile {
    return Profile{Init(&CommandItem{
        Name:        "profile",
        Description: "shows your top everything",
        Aliases:     []string{},
        Usage:       "profile ...",
        Parameters:  []Parameter{},
    })}
}

// Register registers and runs the profile command.
func (c Profile) Register() *atlas.Command {
    c.CommandInterface.Run = func(ctx atlas.Context) {
        ctx.Message.Reply(ctx.Atlas, "Hello, Profile!")
    }

    return c.CommandInterface
}
