package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode865_具有所有最深节点的最小子树
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	res, _ := dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int) (*TreeNode, int) {
	if root == nil {
		return root, level
	}
	leftNode, left := dfs(root.Left, level+1)
	rightNode, right := dfs(root.Right, level+1)
	if left == right {
		return root, left + 1
	} else if left > right {
		return leftNode, left + 1
	}
	return rightNode, right + 1
}
