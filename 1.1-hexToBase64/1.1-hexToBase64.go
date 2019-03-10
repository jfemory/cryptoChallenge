package main

import (
	"fmt"
	"github.com/jfemory/cryptoPals"
)

type Hexstr string

func main() {
	var input Hexstr
	input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Println(cryptoPals.Convert.ToBase64(input))
}
