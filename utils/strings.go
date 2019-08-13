package utils

import "strings"

// ShortStr truncates a string by n length
func ShortStr(str string, n int) string {
	runes := []rune(str)

	if len(runes) > n {
		return string(runes[:n]) + "..."
	}

	return str
}

// JoinString joins a string slice with a char, and removes the end char
func JoinString(strs []string, char string) string {
	return strings.TrimRight(strings.Join(strs, char), char)
}
