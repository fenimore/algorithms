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

// SingleByteCipherOfHex checks a single char cipher against
// a hex.
func SingleByteCipherOfHex(h []byte, cipher byte) ([]byte, error) {
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

// SingleByteCipher checks a single char cipher against
// a hex.
func SingleByteCipher(data []byte, cipher byte) ([]byte, error) {
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

// FindRepeatingKey needed for challenge size
func FindRepeatingKey(crypto [][]byte) []byte {
	key := make([]byte, 0)
	for _, val := range crypto {
		msgs := make(Messages, 0)
		for i := 0; i < 256; i++ {
			result, err = SingleByteCipher(val, byte(i))
			if err != nil {
				fmt.Println(string(val))
			}
			msgs = append(msgs, Message{
				Phrase: result,
				Cipher: byte(i),
				Score:  EvaluatePhrase(string(result)),
			})
		}
		sort.Sort(MesSort(msgs))
		messages = append(messages, msgs[0])
		key = append(key, msgs[0].Cipher)
	}
	return key

}

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

// GetDistances returns a map of keysize to normalized distances
func GetDistances(buffer []byte) (map[int]int, error) {
	// Distances =>
	// Key is keysize
	// Val is distance normalized
	distances := make(map[int]int)

	for keysize := 2; keysize < 41; keysize++ {
		// Four blocks
		a := buffer[:keysize]
		b := buffer[keysize : keysize*2]
		c := buffer[keysize*2 : keysize*3]
		d := buffer[keysize*3 : keysize*4]
		n1, err := HammingDistance(a, b)
		if err != nil {
			return nil, err
		}

		n2, err := HammingDistance(b, c)
		if err != nil {
			return nil, err
		}

		n3, err := HammingDistance(c, d)
		if err != nil {
			return nil, err
		}

		// Normalize Hamming Distance
		n := (n1 + n2 + n3) / 3

		distances[keysize] = n / keysize
	}

	return distances, nil
}

func GetKeySizes(distances map[int]int) []int {
	var min = 1000
	keysizes := make([]int, 0)
	for keysize, distance := range distances {
		if distance <= min {
			min = distance
			keysizes = append(keysizes, keysize)
		}
	}
	return keysizes
}

// TransposeCipher transposes a cipher text into keysize blocks.
func TransposeCipher(buffer []byte, keysize int) [][]byte {
	ciphers := make([][]byte, 0)
	for i := 0; i < len(buffer); i += keysize {
		ciphers = append(ciphers, buffer[i:i+keysize])
	}

	blocks := make([][]byte, len(buffer)/len(ciphers))
	for idx := range blocks {
		blocks[idx] = make([]byte, len(buffer)/keysize)
	}

	for i := 0; i < keysize; i++ {
		for cipher := range ciphers {
			blocks[i][cipher] = ciphers[cipher][i]
		}
	}

	return blocks
}

func DecryptXOR(message, cipher []byte) []byte {
	data := make([]byte, len(message))
	for i := range data {
		data[i] = message[i] ^ cipher[0]
		cipher = CycleByte(cipher)
	}
	return data
}

/* Challenges Set One*/

var (
	input    string
	result   []byte
	expected string
	err      error
	messages Messages
	decoded  []byte
)

func main() {
	// Challenge 1 ################################################################
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

	// Challenge 2 ################################################################
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

	// Challenge 3 ################################################################
	input = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	for i := 0; i < 256; i++ {

		result, err = SingleByteCipherOfHex([]byte(input), byte(i))
		if err != nil {
			fmt.Println(err)
		}
		messages = append(messages, Message{
			Phrase: result,
			Cipher: byte(i),
			Score:  EvaluatePhrase(string(result)),
		})
	}
	sort.Sort(MesSort(messages))
	if "Cooking MC's like a pound of bacon" != string(messages[0].Phrase) {
		fmt.Println("Challenge Three Failed")
	} else {
		fmt.Println("Challenge Three Succeeded")
	}

	// Challenge 4 ################################################################
	messages = nil
	messages = make(Messages, 0)
	hashes := make([][]byte, 0)
	file, err := os.Open("inputs/challenge_4.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // Scan by line
		hashes = append(hashes, []byte(scanner.Text()))
	}
	file.Close()

	for _, val := range hashes {
		msgs := Messages{}
		for i := 0; i < 256; i++ {
			result, err = SingleByteCipherOfHex(val, byte(i))
			if err != nil {
				fmt.Println(err, "Result")
				fmt.Println(string(val))
				break
			}
			//messages = append(messages, Message{
			msgs = append(msgs, Message{
				Phrase: result,
				Cipher: byte(i),
				Score:  EvaluatePhrase(string(result)),
			})
		}
		sort.Sort(MesSort(msgs))
		messages = append(messages, msgs[0])
	}
	sort.Sort(MesSort(messages))
	if string(messages[0].Phrase) != "Now that the party is jumping\n" {
		fmt.Println(string(messages[0].Phrase))
		fmt.Println("Challenge Four Failure")
	} else {
		fmt.Println("Challenge Four Succeeds")
	}

	// Challenge 5 ################################################################
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

	// Challenge 6 ################################################################
	file, err = os.Open("inputs/challenge_6.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	crypto, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	result = make([]byte, base64.StdEncoding.DecodedLen(len(crypto)))
	_, err = base64.StdEncoding.Decode(result, crypto)
	//sDec, _ := base64.StdEncoding.DecodeString(string(crypto))
	// NOTE: decodedLen is longer than n
	if err != nil {
		fmt.Println(err)
	}
	distances, err := GetDistances(result) // possible keysizes
	if err != nil {
		fmt.Println(err)
	}
	possibleKeysize := GetKeySizes(distances)

	//blocks := TransposeCipher(result, keysize)
	//fmt.Println("Result")

	//key := FindRepeatingKey(blocks)
	//fmt.Println(string(key))
	//msg := DecryptXOR(result, key)
	//fmt.Println(string(msg))

}
