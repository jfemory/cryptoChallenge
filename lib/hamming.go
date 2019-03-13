package lib

import (
	"math/bits"
)

func Hamming(a, b []byte) int {
	xored, _ := xorBytes(a, b)
	var out int
	for _, v := range xored {
		out = out + bits.OnesCount(uint(v))
	}
	return out
}
