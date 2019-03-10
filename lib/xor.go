package lib

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"
)

func HexXOR(str1, str2 Hexstr) Hexstr {
	a1, _ := hex.DecodeString(string(str1))
	a2, _ := hex.DecodeString(string(str2))
	a3, err := xorBytes(a1, a2)
	out := hex.EncodeToString(a3)
	if err != nil {
		log.Fatal("Panic", err)
	}
	return Hexstr(out)
}

func Base64XOR(str1, str2 Base64str) Base64str {
	a1, _ := base64.StdEncoding.DecodeString(string(str1))
	a2, _ := base64.StdEncoding.DecodeString(string(str2))
	a3, err := xorBytes(a1, a2)
	out := base64.StdEncoding.EncodeToString(a3)
	if err != nil {
		log.Fatal("Panic", err)
	}
	return Base64str(out)
}

func xorBytes(a1, a2 []byte) ([]byte, error) {
	if len(a1) != len(a2) {
		return make([]byte, 0), errors.New("xorBytes: mismatched byte slice lengths. ")
	}
	out := make([]byte, len(a1))
	for i, v := range a1 {
		out[i] = v ^ a2[i]
	}
	return out, nil
}
