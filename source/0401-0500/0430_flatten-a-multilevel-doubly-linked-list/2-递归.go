package main

func main() {

}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// leetcode430_扁平化多级双向链表
var arr []*Node

func flatten(root *Node) *Node {
	arr = make([]*Node, 0)
	dfs(root)
	for i := 0; i < len(arr); i++ {
		if i+1 < len(arr) {
			arr[i].Next = arr[i+1]
		}
		if i > 0 {
			arr[i].Prev = arr[i-1]
		}
		arr[i].Child = nil
	}
	return root
}

func dfs(root *Node) {
	if root == nil {
		return
	}
	arr = append(arr, root)
	dfs(root.Child)
	dfs(root.Next)
}
