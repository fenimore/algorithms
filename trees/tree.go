// Represent a tree
package main

import "fmt"

type Node struct {
	neighbors Nodes
	visited   bool
}

type Nodes []*Node

func (n *Node) String() string {
	return fmt.Sprintf("%t", n.visited)
}

func (n *Node) grow(branches int) Nodes {
	var children Nodes
	for i := 0; i < branches; i++ {
		newNode := new(Node)
		newNode.visited = false
		n.neighbors = append(n.neighbors, newNode)
		newNode.neighbors = append(newNode.neighbors, n)
		children = append(children, newNode)
	}
	return children
}

var root Node = Node{visited: false}

func main() {
	fmt.Println(root.visited)
	nodes := root.grow(2)
	//for _, neighbor := range root.neighbors {
	//	fmt.Print(neighbor.visited)
	//}
	_ = nodes[0].grow(3)
	_ = nodes[1].grow(1)

}
