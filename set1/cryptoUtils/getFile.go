package getFile

import (
	"bufio"
	"fmt"
	"os"
)

func openFile(a string) [][]byte {
	file, err := os.Open(a)
	if err != nil {
		fmt.Println("File Read Error")
	}
	scanner := bufio.NewScanner(file)
	i := 0
	linecount, _ := lineCount(a)
	bigInput := make([][]byte, linecount)

	for scanner.Scan() {
		bigInput[i] = []byte(scanner.Text())
		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	//Ok, here, we have bigInput, an array of byte arrays corresponding to the file.
	return bigInput
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
