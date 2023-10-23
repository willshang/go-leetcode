package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var targetNode *TreeNode

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	dfsTarget(root, x)
	// 统计根节点、目标节点左子树、目标节点右子树
	return dfs(root) > n/2 || dfs(targetNode.Left) > n/2 || dfs(targetNode.Right) > n/2
}

func dfsTarget(root *TreeNode, target int) {
	if root != nil {
		if root.Val == target {
			targetNode = root
			return
		}
		dfsTarget(root.Left, target)
		dfsTarget(root.Right, target)
	}
}

func dfs(root *TreeNode) int {
	if root == nil || root == targetNode {
		return 0
	}
	return 1 + dfs(root.Left) + dfs(root.Right)
}
