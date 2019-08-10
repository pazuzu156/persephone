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

var commands = map[string]*aurora.Command{}

// Init initializes aurora commands.
func Init(name string, description string, aliases ...string) Command {
	cmd := aurora.NewCommand(name).SetDescription(description)
	commands[cmd.Name] = cmd // used for the help command

	// register aliases
	if aliases != nil {
		cmd.SetAliases(aliases...)
	}

	config := utils.Config()

	lfm := lastfm.New(config.Lastfm.APIKey, config.Lastfm.Secret)

	return Command{cmd, lfm}
}
