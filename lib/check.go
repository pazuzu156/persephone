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
	member, _ := ctx.Atlas.GetMember(ctx.Context, ctx.Message.GuildID, ctx.Message.Author.ID)

	for _, r := range member.Roles {
		groles, _ := ctx.Atlas.GetGuildRoles(ctx.Context, ctx.Message.GuildID)
		guild := GetServer(ctx.Message.GuildID)

		for _, gr := range groles {
			if gr.ID == r {
				if SnowflakeToUInt64(r) == guild.ElevatedRole {
					return true
				}
			}
		}
	}

	return false
}
