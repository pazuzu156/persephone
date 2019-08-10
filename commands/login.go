package commands

import (
	"fmt"
	"persephone/models"
	"persephone/utils"
	"strconv"

	"github.com/polaron/aurora"
	"github.com/shkh/lastfm-go/lastfm"
)

type Login struct {
	Command Command
}

func InitLogin(aliases ...string) Login {
	return Login{Init("login", "Login to the bot with your Lastfm Username")}
}

func (c Login) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		db, _ := utils.OpenDB()
		defer db.Close()

		if len(ctx.Args) > 0 {
			did, _ := strconv.Atoi(ctx.Message.Author.ID.String())
			var res []models.User
			db.Select(&res, db.Where("discord_id", "=", did))

			if len(res) == 0 {
				lfmun := ctx.Args[0]
				lfmuser, err := c.Command.Lastfm.User.GetInfo(lastfm.P{"user": lfmun})
				fmt.Println(lfmuser)

				if err != nil {
					ctx.Message.RespondString(ctx.Aurora, "A user with that username could not be found")

					return
				}

				user := []models.User{
					{
						Username:  lfmuser.Name,
						DiscordID: uint64(did),
						Lastfm:    lfmuser.Name,
					},
				}
				n, _ := db.Insert(user)

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
