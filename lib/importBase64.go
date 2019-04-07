package lib

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"
)

func ImportHex(file string) [][]byte {
	var temp [][]byte //hold the ciphertext, here
	//open file of cipher text, load it into the variable, temp.
	openFile, err := os.Open(file)
	checkError("Failed to open 6.txt ", err)
	defer openFile.Close()

	reader := bufio.NewReader(openFile)

	for {
		line, isPrefix, error := reader.ReadLine()
		if error == io.EOF {
			break
		}
		if isPrefix != false {
			break
		}
		hexed, _ := hex.DecodeString(string(line))
		temp = append(temp, hexed)
	}
	return temp
}

func ImportBase64(file string) []byte {
	var temp []byte //hold the ciphertext, here
	//open file of cipher text, load it into the variable, temp.
	openFile, err := os.Open(file)
	checkError("Failed to open 6.txt ", err)
	defer openFile.Close()

	reader := bufio.NewReader(openFile)
	for {
		line, isPrefix, error := reader.ReadLine()
		if error == io.EOF {
			break
		}
		if isPrefix != false {
			break
		}
		temp = append(temp, line...)
	}
	//convert base64 encoded string source file to []byte for further processing.
	cipherString64 := string(temp)
	cipher, _ := base64.StdEncoding.DecodeString(cipherString64)
	return cipher
}
