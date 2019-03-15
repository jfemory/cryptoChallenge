package main

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"math/bits"
	"os"
	// c for crypto
)

func main() {
	var temp []byte
	openFile, err := os.Open("6.txt")
	checkError("Failed to open 6.txt ", err)
	defer openFile.Close()

	reader := bufio.NewReader(openFile)
	//var out outputData
	for {
		line, isPrefix, error := reader.ReadLine()
		if error == io.EOF {
			break
		}
		if isPrefix != false {
			break
		}
		//do logic here
		temp = append(temp, line...)
	}
	//convert base64 encoded string source file to []byte for further processing.
	cipherString64 := string(temp)
	cipher, _ := base64.StdEncoding.DecodeString(cipherString64)
	keysize, _ := findKeySize(cipher, 2, 40)
	splitBlocks, _ := splitIntoBlocks(cipher, keysize)
	key := scoreBlocks(splitBlocks, keysize)
	fmt.Println(key)
	fmt.Println(string(key))
}

//FindKeySize guess the most likely keysize in a range for a presumed vigernere cipher bsed on the hamming distance of adjacent blocks of bytes.
func findKeySize(cipher []byte, min, max int) (int, error) {
	numberOfBlocks := 7
	if len(cipher) < 4*max { //make sure cipher text is large enough to check for the max keysize guess.
		return 0, errors.New("Max keysize too long to score. ")
	}
	var score float64
	var keySize int

	for keysize := min; keysize <= max; keysize++ {
		blockedCipher := initCipherHolder(keysize, numberOfBlocks, cipher)
		tempKeySize := keysize
		tempScore := hammingHelper(blockedCipher)
		if tempScore > score {
			keySize = tempKeySize
			score = tempScore
		}
	}
	return keySize, nil
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

func splitIntoBlocks(cipher []byte, size int) ([][]byte, error) {
	if len(cipher) < size {
		return nil, errors.New("Cipher text is shorter than keyblock size. ")
	}
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

//scoreBlocks builds the prospective key from sucessive single byte XORs on the splitBlocks.
func scoreBlocks(splitBlocks [][]byte, keysize int) []byte {
	var outBytes []byte
	for _, block := range splitBlocks {
		blockLength := len(block)
		var score float64
		var outKey byte
		for key := 0; key < 256; key++ {
			xored, _ := xorBytes(block, buildKey(byte(key), blockLength))
			tempScore := scoreBytes(xored)
			if tempScore > score {
				score = tempScore
				outKey = byte(key)
			}
		}
		outBytes = append(outBytes, outKey)
	}
	return outBytes
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

func buildKey(key byte, length int) []byte {
	out := make([]byte, length)
	for i := 0; i < length; i++ {
		out[i] = key
	}
	return out
}

//

//checkError returns fatal and an error message, given by the string.
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func hamming(a, b []byte) int {
	xored, _ := xorBytes(a, b)
	var out int
	for _, v := range xored {
		out = out + bits.OnesCount(uint(v))
	}
	return out
}

func hammingHelper(blockedCipher [][]byte) float64 {
	var hammingCounter float64
	counter := 0
	keysize := len(blockedCipher[0])
	for i := 0; i < len(blockedCipher)-1; i++ {
		for j := i + 1; j < len(blockedCipher); j++ {
			hammingScore := hamming(blockedCipher[i], blockedCipher[j])
			hammingCounter = hammingCounter + (float64(hammingScore) / float64(keysize))
			counter++
		}
	}
	return (hammingCounter / float64(counter))
}

func xorBytes(a1, a2 []byte) ([]byte, error) {
	if len(a1) != len(a2) {
		return make([]byte, 0), errors.New(" XOR Bytes: mismatched byte slice lengths. ")
	}
	out := make([]byte, len(a1))
	for i, v := range a1 {
		out[i] = v ^ a2[i]
	}
	return out, nil
}
