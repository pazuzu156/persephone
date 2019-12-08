package commands

import (
	"fmt"
	"persephone/lib"

	"github.com/andersfylling/disgord"
	"github.com/fogleman/gg"
	"github.com/pazuzu156/atlas"
)

// Profile command.
type Profile struct{ Command }

// InitProfile initializes the profile command.
func InitProfile() Profile {
	return Profile{Init(&CommandItem{
		Name:        "profile",
		Description: "Shows your top everything",
		Aliases:     []string{"p"},
		Usage:       "profile [member]",
		Parameters: []Parameter{
			{
				Name:        "member",
				Description: "The user you want to see a profile of",
			},
		},
	})}
}

// Register registers and runs the profile command.
func (c Profile) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		lfmuser, _ := lib.GetLastfmUserInfo(ctx.Message.Author, c.Lastfm)

		bg, _ := lib.OpenImage(lib.LocGet("static/images/background.png"))

		dc := gg.NewContext(1000, 600)
		dc.DrawImage(bg, 0, 0)

		lib.BrandImage(dc)

		r, _ := lib.SaveImage(dc, ctx, "profile")

		ctx.Atlas.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
			Files: []disgord.CreateMessageFileParams{
				{
					FileName: lib.TagImageName(ctx, "profile") + ".png",
					Reader:   r,
				},
			},
			Embed: &disgord.Embed{
				Title: fmt.Sprintf("View %s's Profile on Last.fm", ctx.Message.Author.Username),
				URL:   fmt.Sprintf("https://last.fm/user/%s", lfmuser.Name),
				Color: lib.RandomColor(),
			},
		})
	}

	return c.CommandInterface
}
