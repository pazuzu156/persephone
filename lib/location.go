package lib

import "runtime"

func Storage() string {
	return runtime.GOOS
}
