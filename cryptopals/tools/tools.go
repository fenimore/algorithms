package tools

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/polypmer/algor/cryptopals/words"
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

func ByteToHex(input []byte) []byte {
	data := make([]byte, hex.EncodedLen(len(input)))
	_ = hex.Encode(data, input)
	return data
}

// StringHexToBase64 wrapper around HexToBase64 for using
// strings instead of byte buffers.
func StringHexToBase64(input string) (string, error) {
	result, err := HexToBase64([]byte(input))
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// FixedXORHex takes two hex of equal length
// and returns the hex result of an XOR operation.
func FixedXORHex(a, b []byte) ([]byte, error) {
	first := make([]byte, hex.DecodedLen(len(a)))
	second := make([]byte, hex.DecodedLen(len(b)))
	count1, _ := hex.Decode(first, a)
	count2, _ := hex.Decode(second, b)
	if count1 != count2 {
		return nil, errors.New("Buffers not the same size")
	} else if count1 < 1 {
		return nil, errors.New("Buffer zero length")
	}
	data := make([]byte, count1)
	for i := range data {
		data[i] = first[i] ^ second[i]
	}
	result := make([]byte, hex.EncodedLen(len(data)))
	_ = hex.Encode(result, data)

	return result, nil
}

// Check a hex against a XOR cipher of a single character.
func SingleByteXORCipher(h []byte, cipher byte) ([]byte, error) {
	data := make([]byte, hex.DecodedLen(len(h)))
	_, err := hex.Decode(data, h)
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(cipher), string(data))
	result := make([]byte, len(data))
	for i := range data {
		result[i] = data[i] ^ cipher
	}

	return result, nil
}

// SOOOO
// When you XOR a space with a byte it'll give you the byte.
// And the space is the most frequent letter in the english language.
// SOOOOO, the most frequent byte in the cipher text will be the byte.
func AssumedByteXORCipher(h []byte) ([]byte, byte, error) {
	// Decode Hex
	decod := make([]byte, hex.DecodedLen(len([]byte(h))))
	_, err := hex.Decode(decod, []byte(h))
	if err != nil {
		return nil, 'x', err
	}

	// Find Cipher, most common byte
	// Seeing as how space is the most common letter?
	var cnt int
	var cipher byte
	ciphers := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for _, b := range ciphers {
		// The count of b in decoded hex
		if cnt < bytes.Count(decod, []byte{b}) {
			cnt = bytes.Count(decod, []byte{b})
			cipher = b
		}
	}
	// Maybe it ought to be uppercase what?
	//cipher = []byte(bytes.ToUpper([]byte{cipher}[0 : 0+1]))[0]
	// Find Result
	result := make([]byte, len(decod))
	for i := range decod {
		result[i] = decod[i] ^ cipher
	}
	return result, cipher, nil
}

// CheckFrequency checks frequency of etaoin shrdlu.
// The higher the counter, the most like the phrase
// is an English phrase. Very unsophisticated
func CheckFrequency(data string) int {
	mostFrequent := "etaoin shrdlu"
	var counter int
	data = strings.ToLower(data)
	for _, l := range data {
	Loop:
		for _, f := range mostFrequent {
			if l == f {
				counter++
				break Loop
			}
		}
	}

	return counter
}

func CheckIfAllLetters(data string) bool {
	var IS_LETTER = map[string]bool{
		" ": true,
		"e": true,
		"t": true,
		"a": true,
		"o": true,
		"i": true,
		"n": true,
		"s": true,
		"h": true,
		"r": true,
		"d": true,
		"l": true,
		"c": true,
		"u": true,
		"m": true,
		"w": true,
		"f": true,
		"g": true,
		"y": true,
		"p": true,
		"b": true,
		"v": true,
		"k": true,
		"j": true,
		"x": true,
		"q": true,
		"z": true,
	}

	data = strings.ToLower(data)
	for _, x := range data {
		if !IS_LETTER[string(x)] {
			return false
		}
	}
	return true

}

func DetectSingleCharacterXOR(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var messages []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}
	w := make(words.Words, 0)

	for _, msg := range messages {
		text, c, _ := AssumedByteXORCipher([]byte(msg))
		// fmt.Println(string(c), string(text))
		s := words.EvaluatePhrase(string(text))
		w = append(w,
			words.Word{Phrase: string(text), Score: s,
				Cipher: string(c)})
	}
	sort.Sort(words.WordSorter(w))
	//fmt.Println(w[0].Score, w[len(w)-2].Score)
	for i := 1; i < len(w); i++ {
		fmt.Println(w[len(w)-i])
	}

	return nil
}

func IceEncrypt(input string) []byte {
	//result := make([]byte)
	cipher := []byte{'I', 'C', 'E'}

	data := []byte(input)

	for i := range data {
		data[i] = data[i] ^ cipher[0]
		cipher = CycleByte(cipher)
	}

	result := make([]byte, hex.EncodedLen(len(data)))
	_ = hex.Encode(result, data)

	return result
}

// CycleByte Shifts the slice then Pushes it.
func CycleByte(cipher []byte) []byte {
	first, cipher := cipher[0], cipher[1:]
	cipher = append(cipher, first)
	return cipher
}

// HammingDistance returns the minimun number fo
// substituions required to change one string into the other.
// From wikipedia:
// For binary strings a and b the Hamming distance is equal to the number of ones (hamming weight) in a XOR b.
func HammingDistance(a, b []byte) int {
	return 0
}

// HammingWeight gets us distance
func HammingWeight(a []byte) int {
	return 0
}

// PopulationCount returns number of ones in binary rep.
func PopulationCount(a []byte) int {
	return 0
}
