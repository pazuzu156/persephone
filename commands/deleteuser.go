package commands

import (
	"fmt"
	"strconv"

	"github.com/pazuzu156/atlas"
	"github.com/pazuzu156/persephone/lib"
)

// DeleteUser command.
type DeleteUser struct{ Command }

// InitDeleteUser initializes the deleteuser command.
func InitDeleteUser() DeleteUser {
	return DeleteUser{Init(&CommandItem{
		Name:        "deleteuser",
		Description: "Deletes a user and all related data",
		Aliases:     []string{"du", "removeuser", "ru"},
		Usage:       "deleteuser <discord_id>",
		Parameters: []Parameter{
			{
				Name:        "discord_id",
				Description: "The Discord ID of the user to remove",
				Required:    true,
			},
		},
		Admin: true,
	})}
}

// Register registers and runs the deleteuser command.
func (c DeleteUser) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		if lib.CanRun(ctx) {
			u64, err := strconv.ParseUint(lib.JoinStringMap(ctx.Args, " "), 10, 64)

			if err != nil {
				ctx.Message.Reply(ctx.Context, ctx.Atlas, "You didn't pass in a valid Discord User ID")

				return
			}

			db, _ := lib.OpenDB()
			var users []lib.Users
			db.Select(&users, db.Where("discord_id", "=", u64), db.From(lib.Users{}))

			if len(users) > 0 {
				user := users[0]
				ur, cr := user.Delete()

				if cr && ur {
					ctx.Message.Reply(ctx.Context, ctx.Atlas, fmt.Sprintf("User with ID %d and their crowns have been removed", u64))
				} else if cr && !ur {
					ctx.Message.Reply(ctx.Context, ctx.Atlas, fmt.Sprintf("User with ID %d was not removed, but their crowns were. You might need to run this again", u64))
				} else if !cr && ur {
					ctx.Message.Reply(ctx.Context, ctx.Atlas, fmt.Sprintf("User with ID %d was removed, but their crowns were not You might need to run this again", u64))
				} else {
					ctx.Message.Reply(ctx.Context, ctx.Atlas, fmt.Sprintf("User with ID %d was not removed, and neither were their crowns. Please try again later", u64))
				}
			}
		}
	}

	return c.CommandInterface
}
