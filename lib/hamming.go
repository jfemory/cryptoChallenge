package lib

import (
	"errors"
	"math/bits"
)

type scoredKeySize struct {
	size      int
	normScore float64
}

//Hamming calculates the hamming distance between two byte slices.
func Hamming(a, b []byte) int {
	xored, _ := XorBytes(a, b)
	var out int
	for _, v := range xored {
		out = out + bits.OnesCount(uint(v))
	}
	return out
}

//FindKeySize guess the most likely keysize in a range for a presumed vigernere cipher bsed on the hamming distance of adjacent blocks of bytes.
func FindKeySize(cipher []byte, min, max int) (int, error) {
	numberOfBlocks := 7
	if len(cipher) < 4*max { //make sure cipher text is large enough to check for the max keysize guess.
		return 0, errors.New("Max keysize too long to score. ")
	}
	var key scoredKeySize
	for keysize := min; keysize <= max; keysize++ {
		blockedCipher := initCipherHolder(keysize, numberOfBlocks, cipher)
		tempKey := scoredKeySize{keysize, HammingScore(blockedCipher)}
		if tempKey.normScore > key.normScore {
			key = tempKey
		}
	}
	return key.size, nil
}

func initCipherHolder(keysize, length int, cipher []byte) [][]byte {
	out := make([][]byte, length)
	for i := 0; i < length; i++ {
		low := (keysize * i)
		high := ((keysize * i) + keysize)
		out[i] = cipher[low:high]
	}
	return out
}

//HammingScore returns the normalized hamming score of a slice of byte slices.
func HammingScore(blockedCipher [][]byte) float64 {
	var hammingCounter float64
	counter := 0
	keysize := len(blockedCipher[0])
	for i := 0; i < len(blockedCipher)-1; i++ {
		for j := i + 1; j < len(blockedCipher); j++ {
			hammingScore := Hamming(blockedCipher[i], blockedCipher[j])
			hammingCounter = hammingCounter + (float64(hammingScore) / float64(keysize))
			counter++
		}
	}
	return (hammingCounter / float64(counter))
}
