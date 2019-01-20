package main
import "fmt"
import "encoding/hex"
import "encoding/base64"
//import "unicode/utf8"

const  input string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

func main() {
    fmt.Println(hex2Byte(input))
}

//DecodeHex does Decode a string of hex into a byte array
func hex2Byte(a string) string{
   middle, _ := hex.DecodeString(a)
   out := base64.StdEncoding.EncodeToString(middle)
   fmt.Println(middle)
   fmt.Println(byte2ascii(middle))
  return out

}

func byte2ascii(a []byte) string{
  out := string(a[:])
  return out
}
