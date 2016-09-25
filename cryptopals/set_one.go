package main

import (
	"fmt"

	"github.com/polypmer/algor/cryptopals/bytes"
	"github.com/polypmer/algor/cryptopals/words"
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

// Single Byte Xor Cipher
func c3() {
	given := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	alphabet := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	expected := "X"
	results := make(words.Words, 0)
	for _, cipher := range alphabet {
		text, _ := bytes.SingleByteXORCipher([]byte(given), cipher)
		score := words.EvaluatePhrase(string(text))
		results = append(results, words.Word{Phrase: string(text),
			Cipher: string(cipher), Score: score})
	}
	highest := results.MostFrequent()
	fmt.Printf("Challenge 3: %t\n", highest.Cipher == expected)
}

// Failure
func c4() {
	//_ = bytes.DetectSingleCharacterXOR("inputs/challenge_4.txt")
	fmt.Println("failure - challenge 4")
}

func c5() {
	input := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	answer := bytes.IceEncrypt(input)
	expected := `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`
	fmt.Printf("Challenge 5: %t\n", string(answer) == expected)

}

func c6() {
	input := []byte("this is a test")
	from := []byte("wokka wokka!!!")
	data := bytes.ByteToHex(input)
	comp := bytes.ByteToHex(from)
	xor := make([]byte, len(data))
	for i := range data {
		xor[i] = data[i] ^ comp[i]
	}
	//data, _ = bytes.HexToBase64(data)
	fmt.Println(len(data), len(comp), len(xor))
	//fmt.Println(string(xor))
	//fmt.Println(string(data))
}

func main() {
	//c1()
	//c2()
	//c3()
	//c4() // Failure...
	//c5()
	c6() // the real deal
}
