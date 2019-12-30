package lib

import (
	"fmt"
	"os"
	"runtime"
)

// os.Getenv seems to continue giving me goroutine issues
// hard setting HOME path. Please set your path accordingly
var homepath string = "/home/kaleb"

// Storage returns storage locations for configuration and static folders
func Storage() string {
	if runtime.GOOS == "linux" {
		// return fmt.Sprintf("%s/.config/persephone", os.Getenv("HOME"))

		return fmt.Sprintf("%s/.config/persephone", homepath)
	}

	return ""
}

// LocGet returns file location within bot's storage directory.
func LocGet(file string) string {
	storage := Storage()

	if storage != "" {
		if fileExists(file) {
			return file
		}

		return fmt.Sprintf("%s/%s", storage, file)
	}

	return file
}

func fileExists(file string) bool {
	info, err := os.Stat(file)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
