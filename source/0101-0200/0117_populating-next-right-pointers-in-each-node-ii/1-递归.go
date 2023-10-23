package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	if root.Left != nil {
		root.Left.Next = root.Right
	}
	prev := root.Right
	if prev == nil {
		prev = root.Left
	}
	nextRoot := root.Next
	for nextRoot != nil && (nextRoot.Left == nil && nextRoot.Right == nil) {
		nextRoot = nextRoot.Next
	}
	if nextRoot != nil {
		if nextRoot.Left != nil {
			prev.Next = nextRoot.Left
		} else {
			prev.Next = nextRoot.Right
		}
	}
	connect(root.Right)
	connect(root.Left)
	return root
}
