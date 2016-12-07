package main

import (
	"math"
	"strings"
)

var LETTERS = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var LETTER_FREQUENCY = map[string]float64{
	" ": .14,
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

type Message struct {
	Phrase []byte
	Cipher byte
	Score  float64 // Probability of English
}

type Messages []Message

type MesSort []Message

func (c MesSort) Len() int           { return len(c) }
func (c MesSort) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c MesSort) Less(i, j int) bool { return c[i].Score < c[j].Score }

// EvaluatePhrase takes a string and returns the
// score of probability of it being an English string.
func EvaluatePhrase(phrase string) float64 {
	var sum, score float64
	phrase = strings.ToUpper(phrase)
	frequencies := map[string]float64{
		" ": 0,
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

	// NOTE: This normalizes
	for letter := range frequencies {
		count := float64(strings.Count(phrase, letter))
		frequencies[letter] = count
		sum += count
	}

	// NOTE: ?
	for letter := range frequencies {
		frequencies[letter] /= sum
		score += math.Sqrt(frequencies[letter] * LETTER_FREQUENCY[letter])
	}

	return score
}
