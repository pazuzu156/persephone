package commands

import (
	"fmt"
	"persephone/lib"
	"strconv"
	"time"

	"github.com/andersfylling/disgord"
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
	Usage       string
	Parameters  []Parameter
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

// Init initializes aurora commands
func Init(t *CommandItem) Command {
	cmd := aurora.NewCommand(t.Name).SetDescription(t.Description)

	if t.Aliases != nil {
		cmd.SetAliases(t.Aliases...)
	}

	commands = append(commands, *t)
	config := lib.Config()
	lfm := lastfm.New(config.Lastfm.APIKey, config.Lastfm.Secret)

	return Command{cmd, lfm}
}

// embedFooter returns a footer and timestamp for disgord embeds
func (c Command) embedFooter(ctx aurora.Context) (f *disgord.EmbedFooter, t disgord.Time) {
	f = &disgord.EmbedFooter{
		IconURL: lib.GenAvatarURL(ctx.Message.Author),
		Text:    fmt.Sprintf("Command invoked by: %s#%s", ctx.Message.Author.Username, ctx.Message.Author.Discriminator),
	}

	t = disgord.Time{
		Time: time.Now(),
	}

	return
}

// getBot returns the bot object.
func (c Command) getBot(ctx aurora.Context) *disgord.Member {
	config := lib.Config()
	id, _ := strconv.Atoi(config.BotID)
	bot, _ := ctx.Aurora.GetMember(ctx.Message.GuildID, disgord.NewSnowflake(uint64(id)))

	return bot
}

// getBotUser returns the bot User object.
func (c Command) getBotUser(ctx aurora.Context) *disgord.User {
	return c.getBot(ctx).User
}
