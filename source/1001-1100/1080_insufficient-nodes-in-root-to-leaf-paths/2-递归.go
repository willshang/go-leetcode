package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	return dfs(root, limit, 0)
}

func dfs(root *TreeNode, limit, sum int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		if root.Val+sum < limit { // 需要删除
			return nil
		}
		return root
	}
	left := dfs(root.Left, limit, sum+root.Val)
	right := dfs(root.Right, limit, sum+root.Val)
	if left == nil && right == nil { // 都为nil直接返回
		return nil
	}
	root.Left = left
	root.Right = right
	return root
}
