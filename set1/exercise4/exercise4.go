package main

import "fmt"
import "os"
import "bufio"
import "encoding/hex"

type scoredString struct {
	plainText []byte
	score     int
}

func main() {
	file, err := os.Open("4.txt")
	if err != nil {
		fmt.Println("File Read Error")
	}
	scanner := bufio.NewScanner(file)
	i := 0
	linecount, _ := lineCount("4.txt")
	bigInput := make([][]byte, linecount)

	for scanner.Scan() {
		bigInput[i] = (hex2byte(scanner.Text()))
		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	//Ok, here, we have bigInput, an array of byte arrays corresponding to the file.
	final := bigTab(bigInput)
	fmt.Println(string(final.plainText))
}

func hex2byte(a string) []byte {
	out, _ := hex.DecodeString(a)
	return out
}

func bigTab(input [][]byte) scoredString {
	var output scoredString
	table := buildASCIItable()
	for i := 0; i < len(input); i++ {
		tabbedString := tabulateIt(input[i], table)
		if output.score < tabbedString.score {
			output = tabbedString
		}
	}
	return output
}

func tabulateIt(input []byte, table []byte) scoredString {
	var output scoredString
	for i := 0; i < 255; i++ {
		working := sbyteXORit(byte(i), input)
		scored := scoreIt(working, table)
		if output.score < scored.score {
			output = scored
		}
	}
	return output
}

//LineCount counts the lines of "filename"
// https://stackoverflow.com/questions/29559358/count-lines-via-bufio
// Thanks PeterCO
func lineCount(filename string) (int64, error) {
	lc := int64(0)
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		lc++
	}
	return lc, s.Err()
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
	output = make([]byte, 53)
	for j := 0; j <= 25; j++ {
		low := 65 + j
		output[j] = byte(low)
	}
	for k := 26; k <= 51; k++ {
		high := 71 + k
		output[k] = byte(high)
	}
	output[52] = byte(32)
	return output
}

func scoreIt(input []byte, table []byte) scoredString {
	score := 0
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(input); j++ {
			if table[i] == input[j] {
				score = score + 1
			}
		}
	}
	output := scoredString{input, score}
	return output
}

//flattenMe takes a [][]byte and flattends it into a []byte with extra padding
//of zeros at the end
func flattenMe(a [][]byte) []byte {
	index := 0
	output := make([]byte, len(a)*len(a[1]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			output[index] = a[i][j]
			index++
		}
	}
	return output
}
