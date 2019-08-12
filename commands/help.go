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
			argcmd := ctx.Args[0]

			for _, command := range commands {
				// if argcmd == command.Name then
				// run help, otherwise, likely an
				// alias was used instead
				// which should also work
				if argcmd == command.Name {
					// TODO: Command name given, process help for command
					c.processHelp(ctx, command)
				} else {
					// check if argument was an alias
					for _, alias := range command.Aliases {
						if argcmd == alias {
							// TODO: Alias given, process help for command
							c.processHelp(ctx, command)
						}
					}
				}
			}
		} else {
			var cmdstrslc []string

			for _, command := range commands {
				cmdstrslc = append(cmdstrslc, fmt.Sprintf("`%s`", command.Name))
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

func (c Help) processHelp(ctx aurora.Context, command CommandItem) {
	embedFields := []*disgord.EmbedField{
		{
			Name:  "Help",
			Value: fmt.Sprintf("`%s`: %s", command.Name, command.Description),
		},
	}

	// Usage
	if len(command.Usage) > 0 {
		var usage []string

		for _, i := range command.Usage {
			usage = append(usage, fmt.Sprintf("`%s` - %s", i.Command, i.Description))
		}

		embedFields = append(embedFields, &disgord.EmbedField{
			Name:  "Usage",
			Value: strings.Join(usage, "\n"),
		})
	}

	// Aliases
	if len(command.Aliases) > 0 {
		var aliases []string

		for _, alias := range command.Aliases {
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
