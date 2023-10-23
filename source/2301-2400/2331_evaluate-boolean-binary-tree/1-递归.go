package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode2331_计算布尔二叉树的值
func evaluateTree(root *TreeNode) bool {
	if root.Left == nil || root.Right == nil { // 叶子节点
		return root.Val == 1 //
	}
	if root.Val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	}
	if root.Val == 3 {
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	}
	return true
}
