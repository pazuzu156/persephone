package commands

import (
	"fmt"
	"strings"

	"github.com/pazuzu156/atlas"
)

// Chart command.
type Chart struct{ Command }

// InitChart initializes the chart command.
func InitChart() Chart {
	return Chart{Init(&CommandItem{
		Name:        "chart",
		Description: "Generates a chart for listens",
		Aliases:     []string{},
		Usage:       "chart ...",
		Parameters:  []Parameter{},
	})}
}

// Register registers and runs the chart command.
func (c Chart) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		if len(ctx.Args) > 0 {
			for i, arg := range ctx.Args {
				if strings.HasPrefix(arg, "--") {
					arg = strings.TrimLeft(arg, "--")
					argv := ctx.Args[i+1]
					fmt.Printf("%s: %s\n", arg, argv)
				}
			}
		}
	}

	return c.CommandInterface
}
