package commands

import (
	"fmt"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/aurora"
)

// Help command.
type Help struct {
	Command Command
}

// InitHelp initializes the help command.
func InitHelp(aliases ...string) Help {
	return Help{Init(
		"help",
		"Displays help information for commands",
		[]UsageItem{
			{
				Command:     "help",
				Description: "Shows the master help list",
			},
			{
				Command:     "help [command]",
				Description: "Gets help on a specific command",
			},
		},
		aliases...,
	)}
}

// Register registers and runs the help command.
func (c Help) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		if len(ctx.Args) > 0 {
			if cmd, ok := commands[ctx.Args[0]]; ok {
				embedFields := []*disgord.EmbedField{
					{
						Name:  "Help",
						Value: fmt.Sprintf("`%s`: %s", cmd.Name, cmd.Description),
					},
				}

				// Usage
				if len(usageMap) > 0 {
					var usage []string

					for _, i := range usageMap {
						fmt.Println(c.Command.CommandInterface.Name)
						if cmd.Name == i.CommandName {
							for _, j := range i.Usage {
								usage = append(usage, fmt.Sprintf("`%s` - %s", j.Command, j.Description))
							}
						}
					}

					embedFields = append(embedFields, &disgord.EmbedField{
						Name:  "Usage",
						Value: strings.Join(usage, "\n"),
					})
				}

				// Aliases
				if len(cmd.Aliases) > 0 {
					var aliases []string

					for _, alias := range cmd.Aliases {
						aliases = append(aliases, fmt.Sprintf("`%s`", alias))
					}

					embedFields = append(embedFields, &disgord.EmbedField{
						Name:  "Aliases",
						Value: strings.TrimRight(strings.Join(aliases, ", "), ", "),
					})
				}

				ctx.Aurora.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
					Embed: &disgord.Embed{
						Fields: embedFields,
						Color:  0x007FFF,
					},
				})
			}
		} else {
			var cmdstrslc []string

			for name := range commands {
				cmdstrslc = append(cmdstrslc, fmt.Sprintf("`%s`", name))
			}

			cmdstr := strings.Join(cmdstrslc, ", ")

			ctx.Aurora.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Fields: []*disgord.EmbedField{
						{
							Name:  "Help",
							Value: "Listing all top-level commands. Specify a command to see more information.",
						},
						{
							Name:  "Commands",
							Value: strings.TrimRight(cmdstr, ", "),
						},
					},
					Color: 0x007FFF,
				},
			})
		}
	}

	return c.Command.CommandInterface
}
