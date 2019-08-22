package commands

import (
	"fmt"
	"persephone/utils"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/aurora"
)

// Help command.
type Help struct{ Command }

// InitHelp initializes the help command.
func InitHelp() Help {
	return Help{InitCmd(&CommandItem2{
		Name:        "help",
		Description: "Shows help message",
		Aliases:     []string{"h", "hh"},
		Usage:       "help whoknows",
		Parameters: []Parameter{
			{
				Name:        "command",
				Description: "Gets help on a specific command",
				Required:    false,
			},
		},
	})}
}

// Register registers and runs the help command.
func (c Help) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		if len(ctx.Args) > 0 {
			argcmd := ctx.Args[0]

			for _, command := range commands2 {
				// if argcmd == command.Name then
				// run help, otherwise, likely an
				// alias was used instead
				// which should also work
				if argcmd == command.Name {
					// c.processHelp(ctx, command)
					c.processHelp(ctx, command)
				} else {
					// check if argument was an alias
					for _, alias := range command.Aliases {
						if argcmd == alias {
							// c.processHelp(ctx, command)
							c.processHelp(ctx, command)
						}
					}
				}
			}
		} else {
			var cmdstrslc []string

			for _, command := range commands2 {
				cmdstrslc = append(cmdstrslc, fmt.Sprintf("`%s%s` - %s", config.Prefix, command.Name, command.Description))
			}

			ctx.Aurora.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Fields: []*disgord.EmbedField{
						{
							Name:  "Help",
							Value: "Listing all top-level commands. Specify a command to see more information.",
						},
						{
							Name:  "Commands",
							Value: utils.JoinString(cmdstrslc, "\n"),
						},
					},
					Color: 0x007FFF,
				},
			})
		}
	}

	return c.Command.CommandInterface
}

// processHelp processes help info defined in each command for command specific help pages
func (c Help) processHelp(ctx aurora.Context, command CommandItem2) {
	embedFields := []*disgord.EmbedField{
		{
			Name:  fmt.Sprintf("%s Help", utils.Ucwords(command.Name)),
			Value: fmt.Sprintf("`%s%s`: %s", config.Prefix, command.Name, command.Description),
		},
	}

	// Usage
	if command.Usage != "" {
		embedFields = append(embedFields, &disgord.EmbedField{
			Name:  "Example Usage",
			Value: fmt.Sprintf("`%s%s`", config.Prefix, command.Usage),
		})
	}

	// Parameters
	if len(command.Parameters) > 0 {
		var params []string

		for _, param := range command.Parameters {
			var (
				paramStr  string
				paramName string
			)

			if param.Value != "" {
				paramName = fmt.Sprintf("%s:%s", param.Name, param.Value)
			} else {
				paramName = param.Name
			}

			if param.Required {
				paramStr = fmt.Sprintf("<%s>", paramName)
			} else {
				paramStr = fmt.Sprintf("[%s]", paramName)
			}

			params = append(params, fmt.Sprintf("`%s` - %s",
				paramStr,
				param.Description,
			))
		}

		embedFields = append(embedFields, &disgord.EmbedField{
			Name:  "Parameters",
			Value: utils.JoinString(params, "\n"),
		})
	}

	// Aliases
	if len(command.Aliases) > 0 {
		var aliases []string

		for _, alias := range command.Aliases {
			aliases = append(aliases, fmt.Sprintf("`%s%s`", config.Prefix, alias))
		}

		embedFields = append(embedFields, &disgord.EmbedField{
			Name:  "Aliases",
			Value: utils.JoinString(aliases, ", "),
		})
	}

	ctx.Aurora.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
		Embed: &disgord.Embed{
			Fields: embedFields,
			Color:  0x007FFF,
		},
	})
}
