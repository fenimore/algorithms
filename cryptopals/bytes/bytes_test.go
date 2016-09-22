package bytes

import (
	"sort"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	given := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	result, err := StringHexToBase64(given)
	if err != nil {
		t.Error(err)
	}
	if result != expected {

	}
	var byteResult []byte
	byteResult, err = HexToBase64([]byte(given))
	if err != nil {
		t.Error(err)
	}
	if string(byteResult) != expected {
		t.Error("Unexpected %s", string(result))
	}
}

func TestFixedXORHex(t *testing.T) {
	given := "1c0111001f010100061a024b53535009181c"
	toXor := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"
	result, err := FixedXORHex([]byte(given), []byte(toXor))
	if err != nil {
		t.Error(err)
	}
	if string(result) != expected {
		t.Error("Unexpected %s", string(result))
	}
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

func TestSingleByteCipher(t *testing.T) {
	given := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	alphabet := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	results := make([]CodeWord, 0)
	for _, cipher := range alphabet {
		text, _ := SingleByteXORCipher([]byte(given), cipher)
		likelyWord := CheckFrequency(string(text))
		results = append(results, CodeWord{text: string(text),
			cipher: string(cipher), likely: likelyWord})
	}

	sort.Sort(CodeSorter(results))
	mostLikelyWord := results[len(results)-1]
	if mostLikelyWord.cipher != "X" {
		t.Fail()
	}
}
