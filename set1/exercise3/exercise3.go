package main

import "fmt"
import "encoding/hex"

//import "encoding/base64"
//import "unicode/utf8"

const input1 string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
const n = byte(23)

type scoredString struct{
	plainText []byte
	score int
}

func main() {
workingSlice := hex2byte(input1)
test := buildASCIItable()
finalOutput := tabulateIt(workingSlice, test)
fmt.Println(string(finalOutput.plainText))
fmt.Println("The score of this string is", finalOutput.score)
}

func tabulateIt(input []byte, table []byte) scoredString{
	var output scoredString
	for i := 0; i < 255; i++{
		working := sbyteXORit(byte(i), input)
		scored := scoreIt(working, table)
		if output.score < scored.score {
			output = scored
		}
	}
	return output
}

func scoreIt(input []byte, table []byte) scoredString {
	score := 0
	for i := 0; i < len(table); i++{
		for j := 0; j < len(input); j++{
			if table[i] == input[j] {score = score + 1}
		}
	}
	output := scoredString{input, score}
	return output
}


func hex2byte(a string) []byte {
	out, _ := hex.DecodeString(a)
	return out
}

func sbyteXORit(a byte, b []byte) []byte {
	length := len(b)
	sbSlice := make([]byte, length)
	for i := 0; i < length; i++ {
		sbSlice[i] = a
	}
	out := make([]byte, length)
	for j := 0; j < length; j++ {
		out[j] = b[j] ^ sbSlice[j]
	}
	return out
}

func buildASCIItable() []byte {
	var output []byte
	output = make([]byte, 52)
	for j:=0; j <= 25; j++{
		low := 65 + j
		output[j] = byte(low)
		}
	for k := 26; k <= 51; k++{
		high := 71 + k
		output[k] = byte(high)
	}
	return output
}
