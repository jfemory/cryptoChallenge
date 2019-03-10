/*Set 1 Exercise 1: Convert hex to base64.
Expected Input: "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
Expected Output: "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
*/
package main

import (
	"fmt"

	c "github.com/jfemory/cryptoChallenge/lib" // c for crypto
)

func main() {
	//var input Hexstr
	input := c.Hexstr("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	output := c.ToBase64(input)
	fmt.Println(output)
	fmt.Println(c.ToHex(output))
}
