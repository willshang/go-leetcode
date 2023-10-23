package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode2236_判断根结点是否等于子结点之和
func checkTree(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}
