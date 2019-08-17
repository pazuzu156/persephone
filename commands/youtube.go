package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"persephone/lib"
	"persephone/utils"

	"github.com/pazuzu156/aurora"
)

// Youtube command.
type Youtube struct {
	Command Command
	APIKey  string
	RootURL string
}

// InitYoutube initializes the youtube command.
func InitYoutube() Youtube {
	config := lib.Config()

	return Youtube{
		Command: Init(
			"youtube",
			"Gets a youtube video from query or current playing tack",
			[]UsageItem{},
			[]Parameter{
				{
					Name:        "query",
					Description: "Gets a youtube video from the given search query",
					Required:    false,
				},
			},
			"yt",
		),
		APIKey:  config.YouTube.APIKey,
		RootURL: "https://www.googleapis.com/youtube/v3/search",
	}
}

// Register registers and runs the youtube command.
func (c Youtube) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		ss := map[string]string{"key": c.APIKey, "part": "snippet", "type": "video"}

		if len(ctx.Args) > 0 {
			ss["q"] = utils.JoinString(ctx.Args, " ")
		} else {
			track, err := utils.GetNowPlayingTrack(ctx.Message.Author, c.Command.Lastfm)

			if err != nil {
				ctx.Message.Reply(ctx.Aurora, err.Error())

				return
			}

			ss["q"] = fmt.Sprintf("%s %s", track.Artist.Name, track.Name)
		}

		c.displayResults(ctx, ss)
	}

	return c.Command.CommandInterface
}

// displayResults displays the results of a youtube search.
func (c Youtube) displayResults(ctx aurora.Context, ss map[string]string) {
	qstring := c.stringify(ss)

	resp, err := http.Get(c.RootURL + qstring)

	if err != nil {
		ctx.Message.Reply(ctx.Aurora, err.Error())

		return
	}

	defer resp.Body.Close()

	buf, _ := ioutil.ReadAll(resp.Body)
	var body lib.YouTubeResponse
	json.Unmarshal(buf, &body)

	if len(body.Items) > 0 {
		ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("Result for **%s**: https://youtu.be/%s",
			ss["q"], body.Items[0].ID.VideoID))
	} else {
		ctx.Message.Reply(ctx.Aurora, "No results could be found for what you're listening to")
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

	enc = fmt.Sprintf("?%s", utils.JoinString(urlss, "&"))

	return
}
