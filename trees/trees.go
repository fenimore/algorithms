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

// pass
func (n *Node) Grow(branches int) Nodes {
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

func (n *Node) Leaf(count int) (*Node, int) {
	newNode := new(Node)
	newNode.Visited = false
	newNode.Index = count + 1
	newNode.Neighbors = append(newNode.Neighbors, n)
	n.Neighbors = append(n.Neighbors, newNode)
	return newNode, newNode.Index
}

func NewTree() Nodes {
	var results Nodes
	var root Node = Node{Visited: false, Index: 0}
	results = append(results, &root)

	node, cnt := root.Leaf(root.Index)
	results = append(results, node)
	node, cnt = root.Leaf(cnt)
	results = append(results, node)
	// Third layer
	n, cnt := node.Leaf(cnt)
	results = append(results, n)
	n, cnt = node.Leaf(cnt)
	results = append(results, n)
	// Fourth layer
	child, cnt := n.Leaf(cnt)
	results = append(results, child)
	child, cnt = n.Leaf(cnt)
	results = append(results, child)
	child, cnt = n.Leaf(cnt)
	results = append(results, child)
	return results
}

func OldTree() Nodes {
	var results Nodes
	var root Node = Node{Visited: false, Index: 0}
	results = append(results, &root)
	nodes := root.Grow(2)
	results = append(results, nodes...)
	children := nodes[0].Grow(3)
	results = append(results, children...)
	branches := nodes[1].Grow(1)
	results = append(results, branches...)
	return results
}
