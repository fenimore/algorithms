package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// HexToBase64 takes in a hex and converts it to base64.
func HexToBase64(input []byte) ([]byte, error) {
	data := make([]byte, hex.DecodedLen(len(input)))

	_, err := hex.Decode(data, input)
	if err != nil {
		return nil, err
	}

	result := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(result, data)

	return result, nil
}

// ByteToHex takes []byte and returns a hex
func ByteToHex(input []byte) []byte {
	data := make([]byte, hex.EncodedLen(len(input)))
	_ = hex.Encode(data, input)

	return data
}

// XOR takes two hex of equal length
// and returns the hex result of an XOR operation.
func XORHex(a, b []byte) ([]byte, error) {
	bufA := make([]byte, hex.DecodedLen(len(a)))
	bufB := make([]byte, hex.DecodedLen(len(b)))
	// Decode []bytes into hex
	cntA, err := hex.Decode(bufA, a)
	if err != nil {
		return nil, err
	}

	cntB, err := hex.Decode(bufB, b)
	if err != nil {
		return nil, err
	}
	// NOTE: do I just take len(a) and len(b)
	if cntA != cntB {
		return nil, errors.New("Hex not same size")
	}

	if cntA < 1 {
		return nil, errors.New("Buffer length zero")
	}

	data := make([]byte, cntA)
	for i := range data {
		data[i] = bufA[i] ^ bufB[i]
	}

	result := make([]byte, hex.EncodedLen(len(data)))
	_ = hex.Encode(result, data)

	return result, nil
}

// SingleByteCipher checks a single char cipher against
// a hex.
func SingleByteCipher(h []byte, cipher byte) ([]byte, error) {
	data := make([]byte, hex.DecodedLen(len(h)))
	_, err := hex.Decode(data, h)
	if err != nil {
		return nil, err
	}

	result := make([]byte, len(data))
	for i := range data {
		result[i] = data[i] ^ cipher
	}

	return result, nil
}

// EncryptXOR encrypts message with repeating-key.
func EncryptXOR(message, cipher []byte) []byte {
	data := make([]byte, len(message))
	for i := range data {
		data[i] = message[i] ^ cipher[0]
		cipher = CycleByte(cipher)
	}

	result := make([]byte, hex.EncodedLen(len(message)))
	_ = hex.Encode(result, data)

	return result
}

// CycleByte cycles a repeating cipher key.
func CycleByte(cipher []byte) []byte {
	first, cipher := cipher[0], cipher[1:]
	cipher = append(cipher, first)
	return cipher
}

var KEYSIZE = 2 // to forty

// HammingDistance the number of differing bits
// in two equal length byte slices.
func HammingDistance(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Slices are not the same length")
	}
	var count int
	for i := range a {
		for j := 0; j < 8; j++ {
			differing := (((a[i] >> uint(j)) & 1) == ((b[i] >> uint(j)) & 1))
			if !differing {
				count++
			}
		}
	}

	return count, nil
}

/* Challenges Set One*/

var (
	input    string
	result   []byte
	expected string
	err      error
	messages Messages
)

func main() {
	// Challenge One
	input = "49276d206b696c6c696e6720796f757220627261696e206c"
	input += "696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	result, err = HexToBase64([]byte(input))
	if err != nil {
		fmt.Println(err)
	}
	if string(result) != expected {
		fmt.Println("Challenge one  failed")
	} else {
		fmt.Println("Challenge one Succeeds")
	}

	// Challenge Two
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	expected = "746865206b696420646f6e277420706c6179"

	result, err = XORHex([]byte(input1), []byte(input2))
	if err != nil {
		fmt.Println(err)
	}

	if string(result) != expected {
		fmt.Println("Challenge Two Failed")
	} else {
		fmt.Println("Challenge Two Succeeded")
	}

	// Challenge Three
	input = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	for _, val := range LETTERS {
		result, err = SingleByteCipher([]byte(input), val)
		if err != nil {
			fmt.Println(err)
		}
		messages = append(messages, Message{
			Phrase: result,
			Cipher: val,
			Score:  EvaluatePhrase(string(result)),
		})
	}
	sort.Sort(MesSort(messages))
	if "Cooking MC's like a pound of bacon" != string(messages[len(messages)-1].Phrase) {
		fmt.Println("Challenge Three Failed")
	} else {
		fmt.Println("Challenge Three Succeeded")
	}

	// Challenge Four
	messages = nil
	messages = make(Messages, 0)
	var hashes []string
	file, err := os.Open("inputs/challenge_4.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // Scan by line
		hashes = append(hashes, scanner.Text())
	}
	file.Close()

	for _, val := range hashes {
		for _, letter := range LETTERS {
			result, err = SingleByteCipher([]byte(val), letter)
			if err != nil {
				fmt.Println(err)
			}
			messages = append(messages, Message{
				Phrase: result,
				Cipher: letter,
				Score:  EvaluatePhrase(string(result)),
			})
		}
	}

	fmt.Println("Challenge Four Failure")

	// Challenge 5
	input = `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	result = EncryptXOR([]byte(input), []byte("ICE"))
	expected = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272"
	expected += "a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	if string(result) != expected {
		fmt.Println(string(result), "\n", expected)
		fmt.Println("Challenge Five Failure")
	} else {
		fmt.Println("Challenge Five Success")
	}

	// Challenge 6
	file, err = os.Open("inputs/challenge_6.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	crypto, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// Decode from the base64 encoding NOTE: ?
	result = make([]byte, base64.StdEncoding.DecodedLen(len(crypto)))
	_, _ = base64.StdEncoding.Decode(result, crypto)
	fmt.Println(len(result))

}
