package commands

import (
	"fmt"
	"persephone/database"

	"github.com/pazuzu156/aurora"
	"github.com/shkh/lastfm-go/lastfm"
)

// Login command.
type Login struct {
	Command Command
}

// InitLogin initializes the login command.
func InitLogin(aliases ...string) Login {
	return Login{Init("login", "Login to the bot with your Lastfm Username")}
}

// Register registers and runs the login command.
func (c Login) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		db, _ := database.OpenDB()
		defer db.Close()

		if len(ctx.Args) > 0 {
			dbu := database.GetUser(ctx.Message.Author)

			if len(dbu) == 0 {
				lfmun := ctx.Args[0]
				lfmuser, err := c.Command.Lastfm.User.GetInfo(lastfm.P{"user": lfmun})

				if err != nil {
					ctx.Message.RespondString(ctx.Aurora, "A user with that username could not be found")

					return
				}

				newuser := []database.User{
					{
						Username:  ctx.Message.Author.Username,
						DiscordID: database.GetUInt64ID(ctx.Message.Author),
						Lastfm:    lfmuser.Name,
					},
				}

				n, _ := db.Insert(newuser)

				if n > 0 {
					ctx.Message.RespondString(ctx.Aurora, fmt.Sprintf("%s You have logged in with Last.fm username: `%s`", ctx.Message.Author.Mention(), lfmuser.Name))
				} else {
					ctx.Message.RespondString(ctx.Aurora, "There was a problem saving your information. Please try again later")
				}
			} else {
				ctx.Message.RespondString(ctx.Aurora, fmt.Sprintf("%s You're already logged in to Last.fm", ctx.Message.Author.Mention()))
			}
		} else {
			ctx.Message.RespondString(ctx.Aurora, "You need to provide your Last.fm username to log in")
		}
	}

	return c.Command.CommandInterface
}
