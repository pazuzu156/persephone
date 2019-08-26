package commands

import (
	"persephone/database"

	"github.com/pazuzu156/aurora"
)

// Login command.
type Login struct{ Command }

// InitLogin initializes the login command.
func InitLogin() Login {
	return Login{Init(&CommandItem{
		Name:        "login",
		Description: "Log into the bot with your Last.fm username",
		Aliases:     []string{"li"},
	})}
}

// Register registers and runs the login command.
func (c Login) Register() *aurora.Command {
	c.CommandInterface.Run = func(ctx aurora.Context) {
		db, _ := database.OpenDB()
		defer db.Close()

		// if len(ctx.Args) > 0 {
		// 	if user := database.GetUser(ctx.Message.Author); user.Username == "" {
		// 		lfmun := ctx.Args[0]
		// 		lfmuser, err := c.Lastfm.User.GetInfo(lastfm.P{"user": lfmun})

		// 		if err != nil {
		// 			ctx.Message.Reply(ctx.Aurora, "A user with that username could not be found")

		// 			return
		// 		}

		// 		newuser := []database.User{
		// 			{
		// 				Username:  ctx.Message.Author.Username,
		// 				DiscordID: database.GetUInt64ID(ctx.Message.Author),
		// 				Lastfm:    lfmuser.Name,
		// 			},
		// 		}

		// 		n, _ := db.Insert(newuser)

		// 		if n > 0 {
		// 			ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("%s You have logged in with your Last.fm username: `%s`", ctx.Message.Author.Mention(), lfmuser.Name))
		// 		} else {
		// 			ctx.Message.Reply(ctx.Aurora, "There was a problem saving your information. Please try again later")
		// 		}
		// 	} else {
		// 		ctx.Message.Reply(ctx.Aurora, "You are already logged in with Last.fm")
		// 	}
		// } else {
		// 	ctx.Message.Reply(ctx.Aurora, "You need to provide your Last.fm username to log in")
		// }
	}

	return c.CommandInterface
}
