package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func c1() {
	given := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	var data []byte
	data, _ = hex.DecodeString(given)
	result := base64.StdEncoding.EncodeToString(data)
	fmt.Printf("Challenge 1: %t", result == expected)
}

func main() {
	c1()
}
