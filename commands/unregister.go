package commands

import (
	"fmt"
	"persephone/database"

	"github.com/pazuzu156/aurora"
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
func (c Unregister) Register() *aurora.Command {
	c.CommandInterface.Run = func(ctx aurora.Context) {
		db, _ := database.OpenDB()
		defer db.Close()

		if user := database.GetUser(ctx.Message.Author); user.Username != "" {
			crowns := user.Crowns()
			n, _ := db.Delete(&user)

			if n > 0 {
				for _, crown := range crowns {
					db.Delete(&crown)
				}

				ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("%s You have unregistered successfully", ctx.Message.Author.Mention()))
			} else {
				ctx.Message.Reply(ctx.Aurora, "There was an issue unregistering you out. Please try again later")
			}
		} else {
			ctx.Message.Reply(ctx.Aurora, "You are not registered with Last.fm")
		}
	}

	return c.CommandInterface
}
