/*Set 1 Exercise 1: Convert hex to base64.
Expected Input: "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
Expected Output: "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
*/
package main

import (
	"fmt"
	"github.com/jfemory/crypto"
)

type Hexstr string

func main() {
	var input Hexstr
	input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Println(cryptoPals.Convert.ToBase64(input))
}
