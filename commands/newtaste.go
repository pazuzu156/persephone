package commands

import (
	"os"
	"persephone/lib"

	"github.com/andersfylling/disgord"
	"github.com/fogleman/gg"
	"github.com/pazuzu156/atlas"
)

// Newtaste command.
type Newtaste struct{ Command }

// InitNewtaste initializes the newtaste command.
func InitNewtaste() Newtaste {
	return Newtaste{Init(&CommandItem{
		Name:        "newtaste",
		Description: "new taste",
		Aliases:     []string{"nt"},
		Usage:       "newtaste ...",
		Parameters:  []Parameter{},
	})}
}

// Register registers and runs the newtaste command.
func (c Newtaste) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		dc := gg.NewContext(1000, 600)
		dc.SetRGB(0.2, 0.2, 0.2)
		dc.Clear()

		r, _ := lib.SaveImage(dc, ctx, "taste")

		ctx.Atlas.CreateMessage(ctx.Context, ctx.Message.ChannelID, &disgord.CreateMessageParams{
			Files: []disgord.CreateMessageFileParams{
				{
					FileName: lib.TagImageName(ctx, "taste") + ".png",
					Reader:   r,
				},
			},
		})

		r.Close()
		os.Remove(r.Name())
	}

	return c.CommandInterface
}
