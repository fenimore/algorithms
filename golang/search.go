package main

import "fmt"

func main() {
	x := []int{2, 3, 4, 6, 8, 7, 11, 11, 11, 45, 54, 65}
	fmt.Println(search(4, x))
}

func search(target int, sorted []int) int {
	// Find the indexes
	low := 0
	high := len(sorted) - 1
	mid := len(sorted) / 2
	if target > high {
		fmt.Println("wtf")
		return -1
	}
	// Find the value
	middle := sorted[len(sorted)/2]

	if target == middle {
		return mid
	} else if target > middle {
		// should this be [middle+1:]??
		newList := sorted[middle:high]
		return search(target, newList)
	} else if target < middle {
		newList := sorted[low:middle]
		fmt.Println(newList)
		return search(target, newList)
	}
	return mid
}
