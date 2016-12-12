package main

import (
	"strings"
)

var LETTERS = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var LETTER_FREQUENCY = map[string]float64{
	" ": 20,
	"E": 12.02,
	"T": 9.10,
	"A": 8.12,
	"O": 7.68,
	"I": 7.31,
	"N": 6.95,
	"S": 6.28,
	"R": 6.02,
	"H": 5.92,
	"D": 4.32,
	"L": 3.98,
	"U": 2.88,
	"C": 2.71,
	"M": 2.61,
	"F": 2.30,
	"Y": 2.11,
	"W": 2.09,
	"G": 2.03,
	"P": 1.82,
	"B": 1.49,
	"V": 1.11,
	"K": 0.69,
	"X": 0.17,
	"Q": 0.11,
	"J": 0.10,
	"Z": 0.07,
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
func (c MesSort) Less(i, j int) bool { return c[i].Score > c[j].Score }

// EvaluatePhrase takes a string and returns the
// score of probability of it being an English string.
func EvaluatePhrase(phrase string) float64 {
	var score float64
	phrase = strings.ToUpper(phrase)

	for _, char := range phrase {
		if val, ok := LETTER_FREQUENCY[string(char)]; ok {
			//frequencies[string(char)] += val
			score += val
		}
	}

	return score
}

func GetFrequencies(phrase string) float64 {
	frequencies := make(map[string]float64)
	var score float64
	phrase = strings.ToUpper(phrase)u

	for _, char := range phrase {
		if val, ok := frequencies[string(char)]; ok {
			frequencies += 1
		} else {
			frequencies[string(char)] = 1
		}
	}
	return frequencies
}

func EvaluateFrequencies(phrase string) {
	// freq := GetFrequencies(phrase)
	// var count int
	// var num int
	// for _, char range := LETTER_FREQUENCY {
	//	if
	// }
}
