package utils

import (
	"persephone/lib"
	"strconv"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/snowflake/v3"
	"github.com/pazuzu156/aurora"
)

func GetBot(ctx aurora.Context) *disgord.Member {
	config := lib.Config()
	id, _ := strconv.Atoi(config.BotID)
	bot, _ := ctx.Aurora.GetMember(ctx.Message.GuildID, snowflake.NewSnowflake(uint64(id)))

	return bot
}

func GetBotUser(ctx aurora.Context) *disgord.User {
	return GetBot(ctx).User
}
