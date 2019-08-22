package commands

import (
	"persephone/lib"

	"github.com/pazuzu156/aurora"
	"github.com/pazuzu156/lastfm-go"
)

// Command is the base command object for all commands.
type Command struct {
	CommandInterface *aurora.Command
	Lastfm           *lastfm.API
}

// CommandItem is the base command item object for the help command.
type CommandItem struct {
	Name        string
	Description string
	Aliases     []string
	Usage       []UsageItem
	Parameters  []Parameter
}

// CommandItem2 is the base command item object for the help command.
type CommandItem2 struct {
	Name        string
	Description string
	Aliases     []string
	Usage       []UsageItem
	Parameters  []Parameter
}

// UsageItem is the base usage object for the help command.
type UsageItem struct {
	Command     string
	Description string
}

// UsageItem2 is the base usage object for the help command.
type UsageItem2 struct {
	Command     string
	Description string
}

// Parameter is the base parameter object for the help command.
type Parameter struct {
	Name        string // parameter name
	Value       string // value representation
	Description string // parameter description
	Required    bool   // is parameter required?
}

var (
	// var commands = map[string]*aurora.Command{}
	commands  = []CommandItem{} // Deprecated: reference commands2 instead
	commands2 = []CommandItem2{}
	config    = lib.Config()

	// FontRegular is the name for the regular typed font.
	FontRegular = lib.LocGet("static/fonts/NotoSans-Regular.ttf")

	// FontBold is the name for the bold typed font.
	FontBold = lib.LocGet("static/fonts/NotoSans-Bold.ttf")
)

// Init initializes aurora commands
// Deprecated: Use InitCmd instead.
func Init(name string, description string, usage []UsageItem, params []Parameter, aliases ...string) Command {
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
		Parameters:  params,
	})

	config := lib.Config()
	lfm := lastfm.New(config.Lastfm.APIKey, config.Lastfm.Secret)

	return Command{cmd, lfm}
}

func InitCmd(t *CommandItem2) Command {
	cmd := aurora.NewCommand(t.Name).SetDescription(t.Description)

	if t.Aliases != nil {
		cmd.SetAliases(t.Aliases...)
	}

	commands2 = append(commands2, *t)
	config := lib.Config()
	lfm := lastfm.New(config.Lastfm.APIKey, config.Lastfm.Secret)

	return Command{cmd, lfm}
}
