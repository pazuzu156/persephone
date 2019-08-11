package commands

import (
	"persephone/utils"

	"github.com/pazuzu156/aurora"
	"github.com/pazuzu156/lastfm-go"
)

// Command is the base command object for all commands.
type Command struct {
	CommandInterface *aurora.Command
	Lastfm           *lastfm.Api
}

// Usage is the base usage object for the help command.
type Usage struct {
	CommandName string
	Usage       []UsageItem
}

// UsageItem is the base usage object for the help command.
type UsageItem struct {
	Command     string
	Description string
}

var commands = map[string]*aurora.Command{}
var usageMap = []Usage{}

// Init initializes aurora commands.
func Init(name string, description string, usage []UsageItem, aliases ...string) Command {
	cmd := aurora.NewCommand(name).SetDescription(description)
	commands[cmd.Name] = cmd // used for the help command

	usageMap = append(usageMap, Usage{
		CommandName: cmd.Name,
		Usage:       usage,
	})

	// register aliases
	if aliases != nil {
		cmd.SetAliases(aliases...)
	}

	config := utils.Config()

	lfm := lastfm.New(config.Lastfm.APIKey, config.Lastfm.Secret)

	return Command{cmd, lfm}
}
