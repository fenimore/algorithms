// Sort, Divide and conquer, merge sort. O(n log(n))?
package main

import "fmt"

func main() {
	//x := []int{3, 2, 6, 1, 8, 7, 4, 5}
	//x := []int{3, 2, 6, 1, 8, 7, 4, 5, 11, 45, 54, 65}
	//x := []int{3, 2, 6, 1, 8, 7, 4, 5, 11, 45, 54}
	x := []int{3, 2, 6, 11, 8, 7, 4, 11, 11, 45, 54, 65}
	fmt.Println(MergeSort(x))
}

// mergeSort sorts a list.
func MergeSort(unsorted []int) []int {
	fmt.Println(unsorted)
	c := make([]int, 0)
	a, b := Split(unsorted)
	if len(a) > 1 {
		c = Merge(MergeSort(a), MergeSort(b), c)
	} else {
		c = Merge(a, b, c)
	}
	return c
}

// Split returns two slice, halving the input slice.
func Split(x []int) ([]int, []int) {
	mid := len(x) / 2
	return x[:mid], x[mid:]
}

// Merge, assume the two lists are sorted.
// Pass in an empty list three cause it's recursive
func Merge(a, b, c []int) []int {
	switch {
	case a[0] > b[0]:
		c = append(c, b[0])
		b = b[1:]
	case a[0] < b[0]:
		c = append(c, a[0])
		a = a[1:]
	case a[0] == b[0]:
		c = append(c, a[0])
		a = a[1:]
	}
	//fmt.Println(a, b, c)
	if !(len(a) == 0 || len(b) == 0) {
		return Merge(a, b, c)
	}
	if len(a) > len(b) {
		c = append(c, a...)
	} else if len(b) > len(a) {
		c = append(c, b...)
	}

	return c
}
