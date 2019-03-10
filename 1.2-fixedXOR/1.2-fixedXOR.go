/*Set 1 Exercise 1: Convert hex to base64.
Expected Input: "1c0111001f010100061a024b53535009181c"
Expected Input: "686974207468652062756c6c277320657965"
Expected Output: "746865206b696420646f6e277420706c6179"
*/
package main

import(
c "github.com/jfemory/cryptoChallenge/lib" // c for crypto
)

func main() {
	in1 := "1c0111001f010100061a024b53535009181c"
	in2 := "686974207468652062756c6c277320657965"
	out := c.xorHex(in1, in2)
}
