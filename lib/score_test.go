package lib

import (
	"math"
	"testing"
)

type scoreBytesVector struct {
	input  []byte
	output float64
}

var EvalScoreBytes = []scoreBytesVector{
	{[]byte{31, 65, 66}, 0.6666},
	{[]byte{127, 126, 122, 111, 64, 37, 32, 12, 11, 10}, 0.3},
}

func TestScoreBytes(t *testing.T) {
	for i, test := range EvalScoreBytes {
		result := scoreBytes(test.input)
		if math.Abs(result-test.output) > 0.001 {
			//number from 1 inclusive.
			t.Errorf("Error in lib/score/scoreBytes([]bytes) float64. Check number %d iteration. ", i+1)
		}
	}
}

type scoreHexVector struct {
	input  Hexstr
	output float64
}

var EvalHexScore = []scoreHexVector{
	{Hexstr("6166564f4d410c0a0a0a"), 0.6},
}

func TestScoreStringHex(t *testing.T) {
	for i, test := range EvalHexScore {
		result := ScoreString(test.input)
		if math.Abs(result-test.output) > 0.001 {
			//number from 1 inclusive.
			t.Errorf("Error in lib/score/ScoreString(Hexstr) float64. Check number %d iteration. ", i+1)
		}
	}
}

type scoreBase64Vector struct {
	input  Base64str
	output float64
}

var EvalBase64Score = []scoreBase64Vector{
	{Base64str("aGVsbG8zNTI4MQ=="), 0.5},
}

func TestScoreStringBase64(t *testing.T) {
	for i, test := range EvalBase64Score {
		result := ScoreString(test.input)
		if math.Abs(result-test.output) > 0.001 {
			//number from 1 inclusive.
			t.Errorf("Error in lib/score/ScoreString(Base64str) float64. Check number %d iteration. ", i+1)
		}
	}
}
