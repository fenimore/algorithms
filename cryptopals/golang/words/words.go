package words

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

var LETTER_FREQUENCY = map[string]float64{
	"E": .1202,
	"T": .0910,
	"A": .0812,
	"O": .0768,
	"I": .0731,
	"N": .0695,
	"S": .0628,
	"R": .0602,
	"H": .0592,
	"D": .0432,
	"L": .0398,
	"U": .0288,
	"C": .0271,
	"M": .0261,
	"F": .0230,
	"Y": .0211,
	"W": .0209,
	"G": .0203,
	"P": .0182,
	"B": .0149,
	"V": .0111,
	"K": .0069,
	"X": .0017,
	"Q": .0011,
	"J": .0010,
	"Z": .0007,
}

type Word struct {
	Phrase string
	Cipher string
	Score  float64
}

func (w Word) String() string {
	result := fmt.Sprintf("%s, %f, %s", w.Cipher, w.Score, w.Phrase)
	return result
}

type Words []Word

type WordSorter []Word

func (c WordSorter) Len() int           { return len(c) }
func (c WordSorter) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c WordSorter) Less(i, j int) bool { return c[i].Score < c[j].Score }

// Find word with most typical English frequencies.
func (w Words) MostFrequent() Word {
	sort.Sort(WordSorter(w))
	return w[len(w)-1]
}

// CheckFrequency checks frequency of etaoin shrdlu.
// The higher the counter, the most like the phrase
// is an English phrase.
func EvaluatePhrase(phrase string) float64 {
	var sum, score float64
	phrase = strings.ToUpper(phrase)
	frequencies := map[string]float64{
		"E": 0,
		"T": 0,
		"A": 0,
		"O": 0,
		"I": 0,
		"N": 0,
		"S": 0,
		"R": 0,
		"H": 0,
		"D": 0,
		"L": 0,
		"U": 0,
		"C": 0,
		"M": 0,
		"F": 0,
		"Y": 0,
		"W": 0,
		"G": 0,
		"P": 0,
		"B": 0,
		"V": 0,
		"K": 0,
		"X": 0,
		"Q": 0,
		"J": 0,
		"Z": 0,
	}

	// This normalizes? Or something.
	// I'm not totally sure what I'm doing here.
	for letter := range frequencies {
		count := float64(strings.Count(phrase, letter))
		frequencies[letter] = count
		sum += count
	}

	for letter := range frequencies {
		frequencies[letter] /= sum
		score += math.Sqrt(frequencies[letter] * LETTER_FREQUENCY[letter])
	}

	return score
}
