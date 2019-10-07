package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"persephone/fm"
	"persephone/lib"

	"github.com/pazuzu156/atlas"
)

// Youtube command.
type Youtube struct {
	Command
	APIKey  string
	RootURL string
}

// InitYoutube initializes the youtube command.
func InitYoutube() Youtube {
	config := lib.Config()

	return Youtube{
		Command: Init(&CommandItem{
			Name:        "youtube",
			Description: "Gets a youtube link to your current playing track",
			Aliases:     []string{"yt"},
			Parameters: []Parameter{
				{
					Name:        "query",
					Description: "Gets a youtube link from the given search query",
					Required:    false,
				},
			},
		}),
		APIKey:  config.YouTube.APIKey,
		RootURL: "https://www.googleapis.com/youtube/v3/search",
	}
}

// Register registers and runs the youtube command.
func (c Youtube) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		ss := map[string]string{"key": c.APIKey, "part": "snippet", "type": "video"}

		if len(ctx.Args) > 0 {
			ss["q"] = lib.JoinStringMap(ctx.Args, " ")
		} else {
			track, err := fm.GetNowPlayingTrack(ctx.Message.Author, c.Lastfm)

			if err != nil {
				ctx.Message.Reply(ctx.Atlas, err.Error())

				return
			}

			ss["q"] = fmt.Sprintf("%s %s", track.Artist.Name, track.Name)
		}

		c.displayResults(ctx, ss)
	}

	return c.CommandInterface
}

// displayResults displays the results of a youtube search.
func (c Youtube) displayResults(ctx atlas.Context, ss map[string]string) {
	qstring := c.stringify(ss)

	resp, err := http.Get(c.RootURL + qstring)

	if err != nil {
		ctx.Message.Reply(ctx.Atlas, err.Error())

		return
	}

	defer resp.Body.Close()

	buf, _ := ioutil.ReadAll(resp.Body)
	var body fm.YouTubeResponse
	json.Unmarshal(buf, &body)

	if len(body.Items) > 0 {
		ctx.Message.Reply(ctx.Atlas, fmt.Sprintf("Result for **%s**: https://youtu.be/%s",
			ss["q"], body.Items[0].ID.VideoID))
	} else {
		ctx.Message.Reply(ctx.Atlas, "No results could be found for what you're listening to")
	}
}

// strigify generates a URL GET query string.
func (c Youtube) stringify(ss map[string]string) (enc string) {
	var urlss []string

	for k, v := range ss {
		if k != "apikey" {
			urlss = append(urlss, fmt.Sprintf("%s=%s", k, url.QueryEscape(v)))
		} else {
			urlss = append(urlss, fmt.Sprintf("%s=%s", k, v))
		}
	}

	enc = fmt.Sprintf("?%s", lib.JoinString(urlss, "&"))

	return
}
