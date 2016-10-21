package main

import "fmt"

func main() {
}

func search(target int, sorted []int) int {
	// Find the indexes
	L := 0
	H := len(sorted) - 1
	M := len(sorted) / 2
	if target > H {
		fmt.Println("wtf")
		return -1
	}
	// Find the value
	middle := sorted[len(sorted)/2]
	if target == middle {
		return M
	} else if target > middle {

	} else if target < middle {
	}
}
