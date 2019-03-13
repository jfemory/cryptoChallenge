package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	openFile, err := os.Open("6.txt")
	checkError("Failed to open 6.txt ", err)
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
		//do logic here
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

func findKeysize() []int {

}

func 