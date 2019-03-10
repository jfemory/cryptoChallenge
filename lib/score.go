package lib

import (
	"encoding/base64"
	"encoding/hex"
)

type score interface {
	score() float64
}

func ScoreString(s score) float64 {
	out := s.score()
	return out
}

func (str Hexstr) score() float64 {
	decoded, _ := hex.DecodeString(string(str))
	out := scoreBytes(decoded)
	return out
}

func (str Base64str) score() float64 {
	decoded, _ := base64.StdEncoding.DecodeString(string(str))
	out := scoreBytes(decoded)
	return out
}

func scoreBytes(b []byte) float64 {
	score := 0
	for _, v := range b {
		if (v > 64 && v < 91) || (v > 96 && v < 127) || (v == 32) {
			score++
		}
	}
	return (float64(score) / float64(len(b)))
	//TODO: Write more tests and make more finegrained
}
