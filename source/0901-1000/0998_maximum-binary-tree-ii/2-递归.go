package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	newRoot := root
	if root.Val < val {
		newRoot = &TreeNode{
			Val:  val,
			Left: root,
		}
	} else if root.Right == nil {
		root.Right = &TreeNode{Val: val}
	} else {
		root.Right = insertIntoMaxTree(root.Right, val)
	}
	return newRoot
}
