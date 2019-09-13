package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pazuzu156/aurora"
)

// Register command.
type Register struct{ Command }

// InitRegister initializes the login command.
func InitRegister() Register {
	return Register{Init(&CommandItem{
		Name:        "register",
		Description: "Starts the bot's registration flow",
	})}
}

// RegisterResponse represents the login API response body.
type RegisterResponse struct {
	Token         string `json:"request_token"`
	Expires       int32  `json:"expires"`
	ExpiresString string `json:"expires_string"`
	Error         bool   `json:"error"`
	ErrorMessage  string `json:"message"`
}

// Register registers and runs the login command.
func (c Register) Register() *aurora.Command {
	c.CommandInterface.Run = func(ctx aurora.Context) {

		res, err := http.Get(fmt.Sprintf("%s/login/request_token/%s", config.Website.APIURL, ctx.Message.Author.ID.String()))

		if err != nil {
			ctx.Message.Reply(ctx.Aurora, "An error occurred when attempting to communitate with the authentication server. Please try again later")

			return
		}

		defer res.Body.Close()

		var lr RegisterResponse
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &lr)

		if lr.Error == true {
			ctx.Message.Reply(ctx.Aurora, lr.ErrorMessage)

			return
		}

		url := fmt.Sprintf("%s/auth/authenticate/%s/%s", config.Website.AppURL, ctx.Message.Author.ID.String(), lr.Token)

		ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("Your login request was received. Use this link to begin the login process: %s", url))
		ctx.Message.Reply(ctx.Aurora, fmt.Sprintf("This link %s", lr.ExpiresString))
	}

	return c.CommandInterface
}
