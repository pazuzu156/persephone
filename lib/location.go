package lib

import (
	"fmt"
	"os"
	"runtime"
)

// Storage returns storage locations for configuration and static folders
func Storage() string {
	if runtime.GOOS == "linux" {
		return fmt.Sprintf("%s/.config/persephone", os.Getenv("HOME"))
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
