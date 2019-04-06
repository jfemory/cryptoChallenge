package lib

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"
)

//StringOut is a standard output struct that holds different types of StringOut.
type StringOut struct {
	ASCII  ASCIIstr
	Hex    Hexstr
	Base64 Base64str
}

func RKXencrypt(str ASCIIstr, key ASCIIstr) Hexstr {
	//prepare data
	byteKeyShort := []byte(key)
	byteString := []byte(str)
	byteKeyLong := buildKey(byteKeyShort, byteString)
	a, err := XorBytes(byteString, byteKeyLong)
	out := hex.EncodeToString(a)
	if err != nil {
		log.Fatal("Panic", err)
	}
	return Hexstr(out)
}

func RKXdecrypt(input Hexstr, key ASCIIstr) ASCIIstr {
	byteString, _ := hex.DecodeString(string(input))
	byteKeyShort := []byte(key)
	byteKeyLong := buildKey(byteKeyShort, byteString)
	a, err := XorBytes(byteString, byteKeyLong)
	out := string(a)
	if err != nil {
		log.Fatal("Panic", err)
	}
	return ASCIIstr(out)
}

func HexXOR(str1, str2 Hexstr) Hexstr {
	a1, _ := hex.DecodeString(string(str1))
	a2, _ := hex.DecodeString(string(str2))
	a3, err := XorBytes(a1, a2)
	out := hex.EncodeToString(a3)
	if err != nil {
		log.Fatal("Panic", err)
	}
	return Hexstr(out)
}

func buildKey(key []byte, str []byte) []byte {
	keyLength := len(key)
	strLength := len(str)
	out := make([]byte, strLength)
	for i := 0; i < strLength; i++ {
		out[i] = key[i%keyLength]
	}
	return out
}

func ASCIIXORHex(str1, str2 ASCIIstr) ASCIIstr {
	a1 := []byte(str1)
	a2 := []byte(str2)
	a3, err := XorBytes(a1, a2)
	out := hex.EncodeToString(a3)
	if err != nil {
		log.Fatal("Panic", err)
	}
	return ASCIIstr(out)
}

func Base64XOR(str1, str2 Base64str) Base64str {
	a1, _ := base64.StdEncoding.DecodeString(string(str1))
	a2, _ := base64.StdEncoding.DecodeString(string(str2))
	a3, err := XorBytes(a1, a2)
	out := base64.StdEncoding.EncodeToString(a3)
	if err != nil {
		log.Fatal("Panic", err)
	}
	return Base64str(out)
}

func XorBytes(a1, a2 []byte) ([]byte, error) {
	if len(a1) != len(a2) {
		return make([]byte, 0), errors.New(" XOR Bytes: mismatched byte slice lengths. ")
	}
	out := make([]byte, len(a1))
	for i, v := range a1 {
		out[i] = v ^ a2[i]
	}
	return out, nil
}

type sbxorInterface interface {
	sbXOR(b byte) StringOut
}

func (str Hexstr) sbXOR(b byte) StringOut {
	input, err := hex.DecodeString(string(str))
	if err != nil {
		log.Fatalf("Problem converting from hex to byte.")
	}
	outBytes := make([]byte, len(input))
	for i, v := range input {
		outBytes[i] = v ^ b
	}
	out := StringOut{ASCIIstr(string(outBytes)), Hexstr(hex.EncodeToString(outBytes)), Base64str(base64.StdEncoding.EncodeToString(outBytes))}

	return out
}

func SingleByteXOR(sbX sbxorInterface, b byte) StringOut {
	return sbX.sbXOR(b)
}
