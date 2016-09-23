package tools

import (
	"sort"
	"strings"
)

var LETTER_FREQUENCY = map[string]float64{
	"e": 12.7,
	"t": 9.05,
	"a": 8.16,
	"o": 7.5,
	"i": 6.96,
	"n": 6.74,
	"s": 6.32,
	"h": 6.09,
	"r": 5.98,
	"d": 4.25,
	"l": 4.025,
	"c": 2.782,
	"u": 2.75,
	"m": 2.4,
	"w": 2.36,
	"f": 2.228,
	"g": 2.015,
	"y": 1.974,
	"p": 1.929,
	"b": 1.492,
	"v": 0.978,
	"k": 0.772,
	"j": 0.153,
	"x": 0.150,
	"q": 0.095,
	"z": 0.074,
}

type Word struct {
	Phrase string
	Cipher string
	Score  float64
}

type Words []Word

type WordSorter []Word

func (c WordSorter) Len() int           { return len(c) }
func (c WordSorter) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c WordSorter) Less(i, j int) bool { return c[i].Score < c[j].Score }

// CheckFrequency checks frequency of etaoin shrdlu.
// The higher the counter, the most like the phrase
// is an English phrase. Very unsophisticated
func EvaluatePhrase(phrase string) float64 {
	var score float64
	phrase = strings.ToLower(phrase)
	var letterCount = map[string]int{
		"e": 0,
		"t": 0,
		"a": 0,
		"o": 0,
		"i": 0,
		"n": 0,
		"s": 0,
		"h": 0,
		"r": 0,
		"d": 0,
		"l": 0,
		"c": 0,
		"u": 0,
		"m": 0,
		"w": 0,
		"f": 0,
		"g": 0,
		"y": 0,
		"p": 0,
		"b": 0,
		"v": 0,
		"k": 0,
		"j": 0,
		"x": 0,
		"q": 0,
		"z": 0,
	}

	for l := range letterCount {
		count := strings.Count(phrase, l)
		letterCount[string(l)] = count
	}
	for l := range letterCount {
		val := LETTER_FREQUENCY[string(l)] * float64(letterCount[string(l)])
		score += val
	}

	return score
}

func (w Words) MostFrequent() Word {
	sort.Sort(WordSorter(w))
	return w[len(w)-1]
}
