package lib

import "errors"

func SplitIntoBlocks(cipher []byte, size int) ([][]byte, error) {
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
