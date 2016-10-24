package main

import "github.com/polypmer/algor/trees"
import "fmt"

func main() {
	n := trees.NewTree()
	fmt.Println(n[0].Visited)

	fmt.Println(n)
}
