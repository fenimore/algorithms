package bytes

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
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
	result := make([]byte, len(data))
	for i := range data {
		result[i] = data[i] ^ cipher
	}

	return result, nil
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
	alphabet := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
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
	results := make(words.Words, 0)

	for _, msg := range messages {
		for _, cipher := range alphabet {
			text, _ := SingleByteXORCipher([]byte(msg), cipher)
			score := words.EvaluatePhrase(string(text))
			results = append(results,
				words.Word{Phrase: string(text),
					Cipher: string(cipher), Score: score})
		}
	}
	//count := len(results)
	//sort.Sort(words.WordSorter(results))
	//for i := 1; i < 30; i++ {
	//res := results[count-1]
	//	fmt.Println(res.Score, res.Cipher, res.Phrase)
	//}
	for idx, res := range results {
		//fmt.Println(idx, res.Cipher, res.Phrase)
		//if idx%10 == 0 {
		//	time.Sleep(5000 * time.Millisecond)
		//}
		//if CheckIfAllLetters(res.Phrase) {
		//fmt.Println(idx, res.Phrase)
		//}
		fmt.Println(idx, res.Cipher, res.Phrase)
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
