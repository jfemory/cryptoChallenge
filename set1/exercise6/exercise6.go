package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

//BlockArray is a struct containing all relevant information about a block
//encrypted cyphertext
type BlockArray struct {
	KeySize   int
	BlockText [][]byte
	HammScore float32
	EngScore  float32
	Key       []byte
	Pass      bool
}

var scoreThreshold float32 = 0.95
var asciiTable = buildASCIItable()

func main() {
	bigInput := openFile("6.txt")
	raw := flattenMe(bigInput)
	bigKeys := buildBlockArrayArray(raw, 2, 2) // min/max keysize
	for i := 0; i < len(bigKeys); i++ {
		bigKeys[i] = hammScoreIt(bigKeys[i], 4) // input, itterations
	}
	sort.Slice(bigKeys, func(i, j int) bool { return bigKeys[i].HammScore < bigKeys[j].HammScore })
	//Ok, we're good up through here.
	//fmt.Println(bigKeys)
	permuteSearch(bigKeys)
	//fmt.Println(raw)
	//fmt.Println(bigInput)
}

func permuteSearch(input []BlockArray) []BlockArray {
	for i := 0; i < len(input); i++ {
		tempScore := float32(0)
		n := 0
		for j := 0; j < len(input[i].BlockText[0]); j++ {
			permBlock := buildPermBlock(input[i].BlockText, j)
			scored := scorePermBlock(permBlock)
			input[i].Key[j] = scored.Key
			tempScore = tempScore + scored.Score
			n++
			//fmt.Println()
		}
		input[i].EngScore = tempScore / float32(n)
		//fmt.Println(input)
		fmt.Println()
	}
	return nil
}

//scorePermBlock takes a permuted block and scores it against english after XORing
//it with each ascii character.
func scorePermBlock(input []byte) scoredPerm {
	score := float32(0.0)
	key := byte(0)
	for i := 0; i < len(asciiTable); i++ {

		xordInput := sbyteXORit(asciiTable[i], input)
		scoreTmp := scoreIt(xordInput, asciiTable)
		if score < scoreTmp {
			key = asciiTable[i]
			score = score + scoreTmp
		}
		fmt.Println(string(key))
		fmt.Println(score)

	}
	scoreHolder := scoredPerm{key, score}
	fmt.Println(scoreHolder)
	return scoreHolder
}

func buildPermBlock(input [][]byte, number int) []byte {
	output := make([]byte, len(input))
	for i := 0; i < len(output); i++ {
		output[i] = input[i][number]
	}
	return output
}

//sbyteXORit takes a specific byte and a byte array. It returns the bytes
//array that is XOR'd with the input byte.
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

//scoreIt takes a byte array and a byte array of acceptable values and returns
//a score between 0 and 1.
func scoreIt(input []byte, table []byte) float32 {
	score := 0
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(input); j++ {
			if table[i] == input[j] {
				score = score + 1
			}
		}
	}
	output := float32(score) / float32(len(input))
	return output
}

//scoreBlockArray takes a [][]byte and returns a float64 which is the score of
//the normalized Hamming distance per keysize
func hammScoreIt(a BlockArray, itterations int) BlockArray {
	score := 0
	n := 0
	for i := 0; i < len(a.BlockText); i++ {
		for j := 0; j < len(a.BlockText); j++ {
			if n >= itterations {
				break
			}
			if j <= i {
			} else {
				score = score + hammIt(a.BlockText[j], a.BlockText[i])
				n++
			}
		}
	}
	a.HammScore = (float32(score) / float32(8*n*len(a.BlockText[1])))

	return a
}

//hammIt returns the Hamming distance of two []byte s as an int
func hammIt(a, b []byte) int {
	hammDiff := 0
	for i := 0; i < len(a); i++ {
		abyte := a[i]
		bbyte := b[i]
		for j := 0; j < 8; j++ {
			mask := byte(1 << uint(j))
			if (abyte & mask) != (bbyte & mask) {
				hammDiff++
			}
		}
	}
	return hammDiff
}

//buildBigKeyArray takes a []byte and returns [][][]byte of key lengths 2
//through 40
func buildBlockArrayArray(a []byte, min int, max int) []BlockArray {
	sizeRange := (max - min) + 1
	output := make([]BlockArray, sizeRange)
	//loop over keysize range, 2-40 bytes
	for i := 0; i < sizeRange; i++ {
		keysize := i + min
		middleLength := 0
		if len(a)%keysize < 1 {
			middleLength = (len(a) / keysize)
		} else {
			middleLength = (len(a) / keysize) + 1
		}
		blockArray := make([][]byte, middleLength)
		m := 0
		// build [][]byte out of test keylength
		for j := 0; j < middleLength; j++ {
			blockArray[j] = make([]byte, keysize)
			for k := 0; k < keysize; k++ {
				if m < len(a) {
					blockArray[j][k] = a[m]
				} else {
					blockArray[j][k] = byte(0)
				}
				m++
			}
		}
		output[i].BlockText = blockArray
		output[i].KeySize = keysize
		output[i].HammScore = 0.0
		output[i].Key = make([]byte, keysize)
		output[i].EngScore = 0.0
		output[i].Pass = true
	}
	return output
}

//flattenMe takes a [][]byte and flattens it into a []byte with extra padding
//of zeros at the end to make the output as if all input lines are same length
//Checked good
func flattenMe(a [][]byte) []byte {
	index := 0
	output := make([]byte, (((len(a) - 1) * len(a[1])) + len(a[len(a)-1])))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			output[index] = a[i][j]
			index++
		}
	}
	return output
}

//
// Some moving parts to make things play nice
//
//openFile opens a file into a [][]byte using lineCount
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

//buildASCIItable when called outputs a byte array of all valid english ascii characters with space
func buildASCIItable() []byte {
	output := make([]byte, 95)
	for j := 0; j < len(output); j++ {
		low := 32 + j
		output[j] = byte(low)
	}
	return output
}

//buildSBTable takes a byte and a length and returns an array of bytes of that length
func buildSBTable(a byte, length int) []byte {
	output := make([]byte, length)
	for i := 0; i < length; i++ {
		output[i] = a
	}
	return output
}
