package utils

import (
	"fmt"
	"strings"

	"github.com/andersfylling/disgord"
)

// ShortStr truncates a string by n length.
func ShortStr(str string, n int) string {
	runes := []rune(str)

	if len(runes) > n {
		return string(runes[:n]) + "..."
	}

	return str
}

// JoinString joins a string slice with a char, and removes the end char.
func JoinString(strs []string, char string) string {
	return strings.TrimRight(strings.Join(strs, char), char)
}

// GenAvatarURL generates a URL used to get a user avatar.
func GenAvatarURL(user *disgord.User) string {
	return fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", user.ID.String(), *user.Avatar)
}
