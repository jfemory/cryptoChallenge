package main

import (
	"fmt"

	c "github.com/jfemory/cryptoChallenge/lib" // c for crypto
)

func main() {
	key := c.ASCIIstr("ICE")
	input := c.ASCIIstr("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	out := c.RKXencrypt(input, key)
	fmt.Println(out)
	decrypted := c.RKXdecrypt(out, key)
	fmt.Println(decrypted)
}
