package lib

// Check is a super simple error handler.
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
