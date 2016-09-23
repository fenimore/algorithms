package main

import (
	"fmt"
	"sort"

	"github.com/polypmer/algor/cryptopals/bytes"
)

// Convert Hex to Base64
func c1() {
	given := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	byt, _ := bytes.HexToBase64([]byte(given))

	fmt.Printf("Challenge 1: %t\n", string(byt) == expected)
}

// Fixed Xor of two equal length Hex
func c2() {
	given := "1c0111001f010100061a024b53535009181c"
	toXor := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"
	res, _ := bytes.FixedXORHex([]byte(given), []byte(toXor))
	fmt.Printf("Challenge 2: %t\n", string(res) == expected)
}

type CodeWord struct {
	text   string
	cipher string
	likely int
}

type CodeSorter []CodeWord

func (c CodeSorter) Len() int           { return len(c) }
func (c CodeSorter) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c CodeSorter) Less(i, j int) bool { return c[i].likely < c[j].likely }

// Single Byte Xor Cipher
func c3() {
	given := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	alphabet := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	expected := "X"
	results := make([]CodeWord, 0)
	for _, cipher := range alphabet {
		text, _ := bytes.SingleByteXORCipher([]byte(given), cipher)
		likelyWord := bytes.CheckFrequency(string(text))
		results = append(results, CodeWord{text: string(text),
			cipher: string(cipher), likely: likelyWord})
		//fmt.Println(string(cipher)+": ", string(text), likelyWord)
	}

	sort.Sort(CodeSorter(results))
	mostLikelyWord := results[len(results)-1]
	//fmt.Println(mostLikelyWord.text, "| Cipher:", mostLikelyWord.cipher)
	fmt.Printf("Challenge 3: %t\n", mostLikelyWord.cipher == expected)
}

func c4() {
	_ = bytes.DetectSingleCharacterXOR("inputs/challenge_4.txt")

}

func main() {
	//c1()
	//c2()
	//c3()
	c4()
}
