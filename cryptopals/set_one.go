package main

import "fmt"
import "github.com/polypmer/algor/cryptopals/convert"

// Convert Hex to Base64
func c1() {
	given := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	byt, _ := convert.HexToBase64OfByte([]byte(given))

	fmt.Printf("Challenge 1: %t\n", string(byt) == expected)
}

// Fixed Xor
func c2() {
	given := "1c0111001f010100061a024b53535009181c"
	pair := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"
	fmt.Println(given, pair, expected)
	res, _ := convert.FixedXOR([]byte(given), []byte(pair))
	fmt.Println(string(res))
	byt, _ := convert.HexToBase64OfString(given)
	fmt.Println(string(byt))
}

func main() {
	c1()
	c2()
}
