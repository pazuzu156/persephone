package utils

// ShortStr truncates a string by n length
func ShortStr(str string, n int) string {
	runes := []rune(str)

	if len(runes) > n {
		return string(runes[:n]) + "..."
	}

	return str
}
