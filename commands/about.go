package commands

import (
	"fmt"
	"persephone/utils"
	"strconv"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/snowflake/v3"
	"github.com/pazuzu156/aurora"
)

// About command.
type About struct {
	Command Command
	Version string
}

// Version holds the bot's version number
const Version = "0.0.1"

// InitAbout initialized the about command.
func InitAbout() About {
	return About{Init("about", "Gets information about the bot", []UsageItem{}, []Parameter{}), Version}
}

// Register registers and runs the about command.
func (c About) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		id, _ := strconv.Atoi(config.BotID)
		bot, _ := ctx.Aurora.GetMember(ctx.Message.GuildID, snowflake.NewSnowflake(uint64(id)))

		var roles []string
		for _, r := range bot.Roles {
			groles, _ := ctx.Aurora.GetGuildRoles(ctx.Message.GuildID)

			for _, gr := range groles {
				if gr.ID == r {
					roles = append(roles, gr.Name)
				}
			}
		}

		ctx.Aurora.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       "About Persephone",
				Description: fmt.Sprintf("Persephone is a bot written in Go. Version v%s", c.Version),
				Color:       0x7FFF00,
				Fields: []*disgord.EmbedField{
					{
						Name:  "Name",
						Value: fmt.Sprintf("%s#%s", bot.User.Username, bot.User.Discriminator),
					},
					{
						Name:  "ID",
						Value: bot.User.ID.String(),
					},
					{
						Name:  "Roles",
						Value: utils.JoinString(roles, ", "),
					},
				},
			},
		})
	}

	return c.Command.CommandInterface
}
