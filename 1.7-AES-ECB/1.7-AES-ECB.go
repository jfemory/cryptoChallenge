package main

import (
	"crypto/aes"
	"errors"
	"fmt"

	"github.com/jfemory/cryptoChallenge/lib"
)

func main() {
	key := []byte("YELLOW SUBMARINE")
	block, _ := aes.NewCipher(key)
	cipher := lib.ImportBase64("7.txt")
	splitCipher, _ := splitIntoBlocks(cipher, len(key))
	iv := splitCipher[0]
	splitCipher = splitCipher[1:]
	for i, split := range splitCipher {
		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptoBlocks(split, split)
		fmt.Println(block)
	}

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
