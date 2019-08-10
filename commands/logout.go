package commands

import (
	"fmt"
	"persephone/database"

	"github.com/pazuzu156/aurora"
)

// Logout command.
type Logout struct {
	Command Command
}

// InitLogout initializes the logout command.
func InitLogout(aliases ...string) Logout {
	return Logout{Init("logout", "Logs the user out of the Last.fm integration")}
}

// Register registers and runs the logout command.
func (c Logout) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		dbu := database.GetUser(ctx.Message.Author)

		if len(dbu) > 0 {
			user := dbu[0]

			db, _ := database.OpenDB()
			n, err := db.Delete(&user)

			if err != nil {
				fmt.Println(err)
			}

			if n > 0 {
				ctx.Message.RespondString(ctx.Aurora, fmt.Sprintf("%s You have logged out!", ctx.Message.Author.Mention()))
			} else {
				ctx.Message.RespondString(ctx.Aurora, "There was an issue logging you out. Please try again")
			}
		} else {
			ctx.Message.RespondString(ctx.Aurora, "You are not logged in with Last.fm")
		}
	}

	return c.Command.CommandInterface
}
