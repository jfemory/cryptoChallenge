/*Set 1 Exercise 2: Single-Byte XOR.
Expected Input: "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
Expected Output: "Cooking MC's like a pound of bacon"
*/

package main

import (
	"fmt"

	c "github.com/jfemory/cryptoChallenge/lib" // c for crypto
)

func main() {
	input := c.Hexstr("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	var out c.ASCIIstr
	score := 0.0
	for i := 0; i < 256; i++ {
		xored := c.SingleByteXOR(input, byte(i))
		fmt.Println(xored)
		tempScore := c.ScoreString(xored.Hex)
		if score < tempScore {
			score = tempScore
			out = xored.ASCII
		}
	}
	fmt.Println(out)
}
