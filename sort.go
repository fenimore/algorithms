package main

import "fmt"

func main() {
	c := make([]int, 0)
	a := []int{1, 3, 4}
	b := []int{2, 5, 6}
	fmt.Println(merge(a, b, c))
}

// Merge, assume the two lists are sorted.
// Pass in an empty list three cause it's recursive
func merge(a, b, c []int) []int {
	switch {
	case a[0] > b[0]:
		c = append(c, b[0])
		b = b[1:]
	case a[0] < b[0]:
		c = append(c, a[0])
		a = a[1:]
	}
	//fmt.Println(a, b, c)
	if !(len(a) == 0 || len(b) == 0) {
		return merge(a, b, c)
	}
	if len(a) > len(b) {
		c = append(c, a...)
	} else if len(b) > len(a) {
		c = append(c, b...)
	}

	return c
}
