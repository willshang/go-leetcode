package main

import "container/list"

func main() {

}

type Node struct {
	Val       int
	Neighbors []*Node
}

// leetcode133_克隆图
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, node)
	visited := make(map[*Node]*Node)
	visited[node] = &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]
		for i, v := range temp.Neighbors {
			if _, ok := visited[v]; !ok {
				queue = append(queue, v)
				visited[v] = &Node{
					Val:       v.Val,
					Neighbors: make([]*Node, len(v.Neighbors)),
				}
			}
			visited[temp].Neighbors[i] = visited[v]
		}
	}
	return visited[node]
}
