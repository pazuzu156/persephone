package utils

import (
	"persephone/lib"
	"strconv"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/aurora"
)

// GetBot returns the bot object.
func GetBot(ctx aurora.Context) *disgord.Member {
	config := lib.Config()
	id, _ := strconv.Atoi(config.BotID)
	bot, _ := ctx.Aurora.GetMember(ctx.Message.GuildID, disgord.NewSnowflake(uint64(id)))

	return bot
}

// GetBotUser returns the bot User object.
func GetBotUser(ctx aurora.Context) *disgord.User {
	return GetBot(ctx).User
}
