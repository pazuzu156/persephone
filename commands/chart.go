package commands

import (
	"fmt"
	"strings"

	"github.com/pazuzu156/atlas"
	"github.com/pazuzu156/lastfm-go"
)

// Chart command.
type Chart struct{ Command }

// InitChart initializes the chart command.
func InitChart() Chart {
	return Chart{Init(&CommandItem{
		Name:        "chart",
		Description: "Generates a chart for listens",
		Aliases:     []string{},
		Usage:       "chart --period weekly",
		Parameters: []Parameter{
			{
				Name:        "period",
				Value:       "overall|yearly|monthly|weekly",
				Description: "Gets your chart based on a given period of time",
				Required:    false,
			},
			{
				Name:        "type",
				Value:       "artist|album",
				Description: "Gets a specific type of chart",
				Required:    false,
			},
		},
	})}
}

// Register registers and runs the chart command.
func (c Chart) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		var (
			period = "overall"
			kind   = "artist"
		)

		if len(ctx.Args) > 0 {
			for i, arg := range ctx.Args {
				if strings.HasPrefix(arg, "--") {
					arg = strings.TrimLeft(arg, "--")
					argv, isset := ctx.Args[i+1]

					if !isset {
						ctx.Message.Reply(ctx.Atlas, "A value is required for that argument")

						return
					}

					switch strings.ToLower(arg) {
					case "type":
						switch strings.ToLower(argv) {
						case "album":
							kind = "album"
							break
						case "track":
							kind = "track"
							break
						}

						break
					case "period":
						switch strings.ToLower(argv) {
						case "yearly":
							period = "12month"
							break
						case "monthly":
							period = "1month"
							break
						case "weekly":
							period = "7day"
							break
						}
						break
					}
				}
			}
		}

		if kind == "artist" {
			chart, _ := c.Lastfm.User.GetTopArtists(lastfm.P{"user": c.getLastfmUser(ctx.Message.Author), "period": period})

			for n, ch := range chart.Artists {
				fmt.Printf("%d: %s\n", n+1, ch.Name)
			}
		} else if kind == "album" {
			// TODO: Album chart
		} else {
			// TODO: Track chart
		}
	}

	return c.CommandInterface
}
