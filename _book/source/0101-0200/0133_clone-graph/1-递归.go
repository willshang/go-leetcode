package main

func main() {

}

type Node struct {
	Val       int
	Neighbors []*Node
}

var visited map[*Node]*Node

func cloneGraph(node *Node) *Node {
	visited = make(map[*Node]*Node)
	return clone(node)
}

func clone(node *Node) *Node {
	if node == nil {
		return node
	}
	if v, ok := visited[node]; ok {
		return v
	}
	newNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	visited[node] = newNode
	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors[i] = clone(node.Neighbors[i])
	}
	return newNode
}
