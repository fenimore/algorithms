package main

import "fmt"

func main() {

	fmt.Println(fib(7))
	fmt.Println(fib(1))
	fmt.Println(fib(3))
	fmt.Println(fib(5))
	fmt.Println(fib(100))
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	var c int
	a := 0
	b := 1
	for i := 0; i < n-1; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
