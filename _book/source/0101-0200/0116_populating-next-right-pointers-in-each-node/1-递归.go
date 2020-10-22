package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// leetcode116_填充每个节点下一个右侧节点指针
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	left := root.Left
	right := root.Right
	// 从上往下，连接最中间的
	for left != nil {
		left.Next = right
		left = left.Right
		right = right.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root
}
