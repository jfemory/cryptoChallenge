/* */
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	c "github.com/jfemory/cryptoChallenge/lib" // c for crypto
)

type outputData struct {
	line  int
	Key   byte
	ASCII c.ASCIIstr
	Hex   c.Hexstr
	Score float64
}

func main() {
	openFile, err := os.Open("4.txt")
	checkError("Failed to open prime list file. ", err)
	defer openFile.Close()

	reader := bufio.NewReader(openFile)
	var out outputData
	for {
		line, isPrefix, error := reader.ReadLine()
		if error == io.EOF {
			break
		}
		if isPrefix != false {
			break
		}
		tempOut := evalString(c.Hexstr(string(line)))
		if tempOut.Score > out.Score {
			out = tempOut
		}
	}
	fmt.Println(out.ASCII)
}

//checkError returns fatal and an error message, given by the string.
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func evalString(str c.Hexstr) outputData {
	var out outputData
	out.Score = 0.0
	for i := 0; i < 256; i++ {
		xored := c.SingleByteXOR(str, byte(i))
		tempScore := c.ScoreString(xored.Hex)
		if out.Score < tempScore {
			out.Score = tempScore
			out.ASCII = xored.ASCII
			out.Hex = xored.Hex
			out.Key = byte(i)
		}
	}
	return out
}
