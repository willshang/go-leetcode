package main

func main() {

}

type TreeNode struct {
	Val   int
	Count int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) Insert(node *TreeNode) {
	if node.Val <= root.Val {
		root.Count++
		if root.Left == nil {
			root.Left = node
		} else {
			root.Left.Insert(node)
		}
		return
	}
	node.Count = node.Count + root.Count + 1
	if root.Right == nil {
		root.Right = node
	} else {
		root.Right.Insert(node)
	}
}
