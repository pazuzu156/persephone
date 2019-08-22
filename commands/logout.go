package commands

import (
	"fmt"
	"persephone/database"

	"github.com/pazuzu156/aurora"
)

// Logout command.
type Logout struct{ Command }

// InitLogout initializes the logout command.
func InitLogout() Logout {
	return Logout{Init(&CommandItem{
		Name:        "logout",
		Description: "Logs the user out of the Last.fm integration",
		Aliases:     []string{"lo"},
	})}
}

// Register registers and runs the logout command.
func (c Logout) Register() *aurora.Command {
	c.CommandInterface.Run = func(ctx aurora.Context) {
		db, _ := database.OpenDB()
		defer db.Close()

		if user := database.GetUser(ctx.Message.Author); user.Username != "" {
			n, _ := db.Delete(&user)

			if n > 0 {
				ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("%s You have logged out successfully", ctx.Message.Author.Mention()))
			} else {
				ctx.Message.Reply(ctx.Aurora, "There was an issue logging you out. Please try again later")
			}
		} else {
			ctx.Message.Reply(ctx.Aurora, "You are not logged in with Last.fm")
		}
	}

	return c.CommandInterface
}
