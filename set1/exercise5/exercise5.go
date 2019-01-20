package main

import "fmt"
import "encoding/hex"

func main() {
	const input1 = "this is a test"
	const input2 = "wokka wokka!!!"

	modInput1 := hex.DecodeString(input1)
	modInput2 := hex.DecodeString(input2)
	fmt.Println(modInput1)
	fmt.Println(modInput2)
}
