package commands

import (
	"persephone/lib"

	"github.com/pazuzu156/aurora"
	"github.com/pazuzu156/lastfm-go"
)

// Command is the base command object for all commands.
type Command struct {
	CommandInterface *aurora.Command
	Lastfm           *lastfm.Api
}

// CommandItem is the base command item object for the help command.
type CommandItem struct {
	Name        string
	Description string
	Aliases     []string
	Usage       []UsageItem
}

// UsageItem is the base usage object for the help command.
type UsageItem struct {
	Command     string
	Description string
}

// var commands = map[string]*aurora.Command{}
var commands = []CommandItem{}

// Init initializes aurora commands.
func Init(name string, description string, usage []UsageItem, aliases ...string) Command {
	cmd := aurora.NewCommand(name).SetDescription(description)

	// register aliases
	if aliases != nil {
		cmd.SetAliases(aliases...)
	}

	// Add all command info to slice for help command
	commands = append(commands, CommandItem{
		Name:        cmd.Name,
		Description: cmd.Description,
		Aliases:     cmd.Aliases,
		Usage:       usage,
	})

	config := lib.Config()
	lfm := lastfm.New(config.Lastfm.APIKey, config.Lastfm.Secret)

	return Command{cmd, lfm}
}
