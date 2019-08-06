package commands

import (
	"persephone/utils"

	"github.com/polaron/aurora"
	"github.com/shkh/lastfm-go/lastfm"
)

// Command is the base command object for all commands.
type Command struct {
	CommandInterface *aurora.Command
	Lastfm           *lastfm.Api
}

// Init initializes aurora commands.
func Init(name string, description string, aliases ...string) Command {
	cmd := aurora.NewCommand(name).SetDescription(description)

	if aliases != nil {
		cmd.SetAliases(aliases...)
	}

	config := utils.Config()

	lfm := lastfm.New(config.Lastfm.APIKey, config.Lastfm.Secret)

	return Command{cmd, lfm}
}
