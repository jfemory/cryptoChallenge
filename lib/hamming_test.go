package lib

import (
	"fmt"
	"testing"
)

type hammingVector struct {
	input1 []byte
	input2 []byte
	output int
}

var EvalHamming = []hammingVector{
	{[]byte("this is a test"), []byte("wokka wokka!!!"), 37},
}

func TestHamming(t *testing.T) {
	for _, test := range EvalHamming {
		result := Hamming(test.input1, test.input2)
		if result != test.output {
			//number from 1 inclusive.
			fmt.Printf("Expected %d got %d", test.output, result)
			t.Errorf("Error in hamming.go ")
		}

	}
}
