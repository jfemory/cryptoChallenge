package main

import "fmt"
import "encoding/hex"
//import "encoding/base64"

//import "unicode/utf8"

const input1 string = "1c0111001f010100061a024b53535009181c"
const input2 string = "686974207468652062756c6c277320657965"

func main() {
	byte1 := hex2byte(input1)
	byte2 := hex2byte(input2)
  result := XORit(byte1, byte2)
  fmt.Println(result)
  fmt.Println(hex.EncodeToString(result))
  fmt.Println(byte2ascii(result))
}

//DecodeHex does Decode a string of hex into a byte array
func hex2byte(a string) []byte {
	out, _ := hex.DecodeString(a)
	return out

}

func byte2ascii(a []byte) string {
	out := string(a[:])
	return out
}

//XORit xors two byte slices and returns a byte slice
func XORit(a []byte, b []byte) []byte {
  l := len(a)
  out := make([]byte, l)
	for i:=0; i < len(a); i++ {
    out[i] = a[i] ^ b[i]
	}
  return out
}
