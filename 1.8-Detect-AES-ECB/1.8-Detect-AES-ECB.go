package main

import (
	"fmt"

	"github.com/jfemory/cryptoChallenge/lib"
)

type scoredCipher struct {
	score  float64
	cipher [][]byte
}

func main() {
	cipherText := lib.ImportHex("8.txt")
	blockSize := 16
	fmt.Println(blockSize)
	fmt.Println(len(cipherText[0]) / 16)
	fmt.Println(len(cipherText))

	var tmpCipher scoredCipher

	for _, cipher := range cipherText {
		blocks, _ := splitIntoBlocks(cipher, blockSize)
		test := lib.HammingScore(blocks)
		if test > tmpCipher.score {
			tmpCipher.score = test
			tmpCipher.cipher = blocks
		}
		fmt.Println(test)
		for _, v := range blocks {
			fmt.Println(v)
		}
		fmt.Println("*************************")
	}
	fmt.Println(tmpCipher.score)
	for _, v := range tmpCipher.cipher {
		fmt.Println(v)
	}
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
func splitIntoBlocks(cipher []byte, size int) ([][]byte, error) {
	cipherSize := len(cipher) //length of cipher
	out := makeBlankBlocks(size)
	for i := 0; i < cipherSize; i++ {
		index := i % size
		out[index] = append(out[index], cipher[i])
	}
	return out, nil
}

func makeBlankBlocks(size int) [][]byte {
	out := make([][]byte, size)
	for i := 0; i < size; i++ {
		out[i] = make([]byte, 0)
	}
	return out
}
