package lib

import "github.com/pazuzu156/atlas"

// Check is a super simple error handler.
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// CanRun checks if a user can run a given elevated command.
func CanRun(ctx atlas.Context) bool {
	member, _ := ctx.Atlas.GetMember(ctx.Message.GuildID, ctx.Message.Author.ID)

	for _, r := range member.Roles {
		groles, _ := ctx.Atlas.GetGuildRoles(ctx.Message.GuildID)

		for _, gr := range groles {
			if gr.ID == r {
				if r.String() == config.ElevatedRole {
					return true
				}
			}
		}
	}

	return false
}
