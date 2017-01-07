// levenshtein edit difference between words
// use for fuzzy search
// adapted from https://en.wikibooks.org/w/index.php?title=Algorithm_Implementation/Strings/Levenshtein_distance&stable=0
// c implementatation
// TODO: pass in max distance
package main

import (
	"fmt"
)

func Leven(a, b string) int {
	if len(a) == 0 {
		return len(b)
	} else if len(b) == 0 {
		return len(a)
	}

	column := make([]int, len(a)+1)

	for i := 1; i <= len(a); i++ {
		column[i] = i
	}
	for i := 1; i <= len(b); i++ {
		column[0] = i
		lastDiag := i - 1
		for j := 1; j <= len(a); j++ {
			oldDiag := column[j]
			column[j] = min(column[j]+1,
				column[j-1]+1, lastDiag+score(a[j-1], b[i-1]))
			lastDiag = oldDiag
		}
	}

	return column[len(a)]
}

func score(a, b byte) int {
	if a == b {
		return 0
	}
	return 1
}

///#define MIN3(a, b, c)
// ((a) < (b) ? ((a) < (c) ? (a) : (c))
// : ((b) < (c) ? (b) : (c)))

// min returns the lowest of three numbers
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	} else if b < c {
		return b
	}

	return c
}

func main() {
	i := Leven("help", "hello")
	fmt.Println(i)
}
