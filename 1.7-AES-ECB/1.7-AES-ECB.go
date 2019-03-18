package main

import (
	"crypto/aes"
	"errors"
	"fmt"

	"github.com/jfemory/cryptoChallenge/lib"
)

func main() {
	key := []byte("YELLOW SUBMARINE") //16 byte key selects AES-128.
	block, _ := aes.NewCipher(key)
	cipherText := lib.ImportBase64("7.txt")

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(cipherText) < aes.BlockSize {
		panic("ciphertext too short")
	}
	//iv := cipherText[:aes.BlockSize]
	mode := lib.NewECBDecrypter(block)
	mode.CryptBlocks(cipherText, cipherText)
	fmt.Println(string(cipherText))
}

func splitIntoBlocks(cipher []byte, size int) ([][]byte, error) {
	if len(cipher) < size {
		return nil, errors.New("Cipher text is shorter than keyblock size. ")
	} else if size == 0 {
		size++
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
