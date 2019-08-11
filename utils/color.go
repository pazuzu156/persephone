package utils

import (
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// RandomColor returns a random color.
func RandomColor() int {
	chars := []string{
		"a", "b", "c", "d", "e", "f", // letters
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", // digits
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))

	var max = 6
	var color = "0x"

	for i, j := range r.Perm(len(chars)) {
		if i < max {
			color += chars[j]
		}
	}

	h, _ := hexutil.DecodeUint64(color)

	return int(h)
}
