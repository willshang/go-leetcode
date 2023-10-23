package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1080_根到叶路径上的不足节点
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		if root.Val < limit { // 需要删除
			return nil
		}
		return root
	}
	left := sufficientSubset(root.Left, limit-root.Val)
	right := sufficientSubset(root.Right, limit-root.Val)
	if left == nil && right == nil { // 都为nil直接返回
		return nil
	}
	root.Left = left
	root.Right = right
	return root
}
