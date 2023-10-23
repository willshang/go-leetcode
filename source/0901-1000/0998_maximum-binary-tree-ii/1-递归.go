package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode998_最大二叉树II
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	if val > root.Val {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}
