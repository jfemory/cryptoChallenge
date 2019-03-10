package lib

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

/*
Lowercase 'to': Methods.
Capital 'To': Exported function.
Ends in 'str': exported string type.
Capitalized encoding type: interface that is satisficed by other types.
*/

//Hexstr is a string of bytewise encodded hex.
type Hexstr string

//Base64str is a string of encoded base64.
type Base64str string

//ASCIIstr is a byte string represented as ASCII.
type ASCIIstr string

//Base64 is satisfied by Hexstr, ...
type Base64 interface {
	toBase64() Base64str
}

//Hex is satisfied by Base64str
type Hex interface {
	toHex() Hexstr
}

//ToBase64 takes a Base64 interface and returns a Base64str
func ToBase64(str Base64) Base64str {
	return str.toBase64()
}

//ToHex takes a Hex interface and returns a Hexstr
func ToHex(str Hex) Hexstr {
	return str.toHex()
}

//toBase64 (Hexstr) coverts a Hexstr to Base64str
func (str Hexstr) toBase64() Base64str {
	out, err := hex.DecodeString(string(str))
	checkError("Failed to convert hex to base64. ", err)
	return Base64str(base64.StdEncoding.EncodeToString(out))
}

func (str Base64str) toHex() Hexstr {
	out, err := base64.StdEncoding.DecodeString(string(str))
	checkError("Failed to convert base64 to hex. ", err)
	return Hexstr(hex.EncodeToString(out))
}

//checkError returns fatal and an error message, given by the string.
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
