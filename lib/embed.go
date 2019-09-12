package lib

import (
	"fmt"
	"time"

	"github.com/andersfylling/disgord"
)

// AddEmbedFooter returns a footer and timestamp for disgord embeds
func AddEmbedFooter(msg *disgord.Message) (f *disgord.EmbedFooter, t disgord.Time) {
	f = &disgord.EmbedFooter{
		IconURL: GenAvatarURL(msg.Author),
		Text:    fmt.Sprintf("Command invoked by: %s#%s", msg.Author.Username, msg.Author.Discriminator),
	}

	t = disgord.Time{
		Time: time.Now(),
	}

	return
}
