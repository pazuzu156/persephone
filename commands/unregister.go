package commands

import (
	"fmt"
	"persephone/lib"

	"github.com/pazuzu156/atlas"
)

// Unregister command.
type Unregister struct{ Command }

// InitUnregister initializes the logout command.
func InitUnregister() Unregister {
	return Unregister{Init(&CommandItem{
		Name:        "unregister",
		Description: "Logs the user out of the Last.fm integration",
	})}
}

// Register registers and runs the logout command.
func (c Unregister) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		db, _ := lib.OpenDB()
		defer db.Close()

		if user := lib.GetUser(ctx.Message.Author); user.Username != "" {
			crowns := user.Crowns()
			n, _ := db.Delete(&user)

			if n > 0 {
				for _, crown := range crowns {
					db.Delete(&crown)
				}

				ctx.Message.Reply(ctx.Context, ctx.Atlas, fmt.Sprintf("%s You have unregistered successfully", ctx.Message.Author.Mention()))
			} else {
				ctx.Message.Reply(ctx.Context, ctx.Atlas, "There was an issue unregistering you out. Please try again later")
			}
		} else {
			ctx.Message.Reply(ctx.Context, ctx.Atlas, "You are not registered with Last.fm")
		}
	}

	return c.CommandInterface
}
