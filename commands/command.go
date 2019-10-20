package commands

import (
	"persephone/lib"
	"strconv"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
	"github.com/pazuzu156/lastfm-go"
)

// Command is the base command object for all commands.
type Command struct {
	CommandInterface *atlas.Command
	Lastfm           *lastfm.API
}

// CommandItem is the base command item object for the help command.
type CommandItem struct {
	Name        string
	Description string
	Aliases     []string
	Usage       string
	Parameters  []Parameter
	Admin       bool
}

// Parameter is the base parameter object for the help command.
type Parameter struct {
	Name        string // parameter name
	Value       string // value representation
	Description string // parameter description
	Required    bool   // is parameter required?
}

var (
	commands = []CommandItem{}
	config   = lib.Config()

	// FontRegular is the name for the regular typed font.
	FontRegular = lib.LocGet("static/fonts/NotoSans-Regular.ttf")

	// FontBold is the name for the bold typed font.
	FontBold = lib.LocGet("static/fonts/NotoSans-Bold.ttf")
)

// Init initializes atlas commands
func Init(t *CommandItem) Command {
	cmd := atlas.NewCommand(t.Name).SetDescription(t.Description)

	if t.Aliases != nil {
		cmd.SetAliases(t.Aliases...)
	}

	commands = append(commands, *t)
	config := lib.Config()
	lfm := lastfm.New(config.Lastfm.APIKey, config.Lastfm.Secret)

	return Command{cmd, lfm}
}

// embedFooter returns a footer and timestamp for disgord embeds
func (c Command) embedFooter(ctx atlas.Context) (f *disgord.EmbedFooter, t disgord.Time) {
	f, t = lib.AddEmbedFooter(ctx.Message)

	return
}

// getBot returns the bot object.
func (c Command) getBot(ctx atlas.Context) *disgord.Member {
	config := lib.Config()
	id, _ := strconv.Atoi(config.BotID)
	bot, _ := ctx.Atlas.GetMember(ctx.Message.GuildID, disgord.NewSnowflake(uint64(id)))

	return bot
}

// getBotUser returns the bot User object.
func (c Command) getBotUser(ctx atlas.Context) *disgord.User {
	return c.getBot(ctx).User
}

// getLastfmUser returns a Last.FM username from the database from a given discord user.
func (c Command) getLastfmUser(user *disgord.User) string {
	return lib.GetUser(user).Lastfm
}
