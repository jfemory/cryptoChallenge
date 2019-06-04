package main

import "fmt"

func main() {
	input := []byte("YELLOW SUBMARINE")
	out := pad(input, byte(97), 20)
	fmt.Println(string(out))
}

func pad(a []byte, val byte, length int) []byte {
	padBy := length - (len(a) % length)
	for i := 0; i < padBy; i++ {
		a = append(a, val)
	}
	return a
}
