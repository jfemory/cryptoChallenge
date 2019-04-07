package main

import (
	"fmt"

	"github.com/jfemory/cryptoChallenge/lib"
)

type scoredBlocks struct {
	score  int
	cipher [][]byte
}

func main() {
	cipherText := lib.ImportHex("8.txt")
	blockSize := 16
	var scored scoredBlocks

	for _, cipher := range cipherText {
		var temp scoredBlocks
		blocks := splitIntoBlocks(cipher, blockSize)
		temp.cipher = blocks
		for i := 0; i < len(blocks)-1; i++ {
			for j := i + 1; j < len(blocks); j++ {
				for k, v := range blocks[i] {
					if v == blocks[j][k] {
						temp.score++
					}
				}
			}
		}
		if temp.score > scored.score {
			scored = temp
		}
	}
	for _, v := range scored.cipher {
		fmt.Println(v)
	}
}

//isEqual takes two byte slices and determines if they are equal, i.e. each byte in order has the same value.
func isEqual(a, b []byte) bool {
	testSlice, _ := lib.XorBytes(a, b)
	return checkZeroBytes(testSlice)
}

//checkZeroBytes returns true if all bytes are zero, else it returns false
func checkZeroBytes(checkMe []byte) bool {
	counter := 0
	zero := byte(0)
	for _, val := range checkMe {
		if val != zero {
			counter++
		}
	}
	if counter == 0 {
		return true
		fmt.Println(checkMe)
	}
	return false
}

//splitIntoBlocks takes a cipher/plain text worth of bytes and splits it into blocks of a given size, returning the [][]byte
func splitIntoBlocks(cipher []byte, size int) [][]byte {
	var out [][]byte
	cipherSize := len(cipher) //length of cipher

	outLength := cipherSize / size
	if cipherSize%size != 0 {
		panic("cipher length is non zero mod block size.")
	}
	for i := 0; i < outLength; i++ {
		out = append(out, cipher[0:size])
		cipher = cipher[size:]
	}
	return out
}

func makeBlankBlocks(size int) [][]byte {
	out := make([][]byte, size)
	for i := 0; i < size; i++ {
		out[i] = make([]byte, 0)
	}
	return out
}
