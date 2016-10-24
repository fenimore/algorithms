// Represent a tree
package trees

import "fmt"

type Node struct {
	Index     int
	Neighbors Nodes
	Visited   bool
}

type Nodes []*Node

func (n *Node) String() string {
	return fmt.Sprintf("%d %t ", n.Index, n.Visited)
}

func (n *Node) grow(branches int) Nodes {
	//parentIndex := n.Index + 1
	var children Nodes
	for i := 0; i < branches; i++ {
		newNode := new(Node)
		newNode.Visited = false
		//newNode.Index = parentIndex + i
		n.Neighbors = append(n.Neighbors, newNode)
		newNode.Neighbors = append(newNode.Neighbors, n)
		children = append(children, newNode)
	}
	return children
}

func NewTree() Nodes {
	var results Nodes
	var root Node = Node{Visited: false, Index: 0}
	results = append(results, &root)
	nodes := root.grow(2)
	results = append(results, nodes...)
	children := nodes[0].grow(3)
	results = append(results, children...)
	branches := nodes[1].grow(1)
	results = append(results, branches...)
	return results
}
