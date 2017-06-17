package tree

import "fmt"

type Node struct {
	Name     string
	Path     string // Only leafs have a path
	ID       int
	ParentID int
	Nodes    []*Node
}

func (n *Node) NewRoot(rootName string) {
	if n.Name == "" {
		n.Name = rootName
	}
}

func Print(n *Node, depth int) {
	var indent []byte
	for i := 0; i < depth; i++ {
		indent = append(indent, []byte(".")...)
	}
	fmt.Println(string(indent) + n.Name)
	for _, subN := range n.Nodes {
		Print(subN, depth+1)
	}
}

func (n *Node) GetOrMakeChildWithName(nodeName string) *Node {
	for _, currentNode := range n.Nodes {
		if currentNode.Name == nodeName {
			return currentNode
		}
	}
	newNode := &Node{Name: nodeName}
	n.Nodes = append(n.Nodes, newNode)
	return newNode
}
