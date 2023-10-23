package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode863_二叉树中所有距离为K的结点
var m map[int]*TreeNode // 存储值对应的父节点
var res []int

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	m = make(map[int]*TreeNode)
	res = make([]int, 0)
	dfs(root)                  // 生成值对应的父节点
	findDfs(K, target, nil, 0) // 遍历
	return res
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		m[root.Left.Val] = root
		dfs(root.Left)
	}
	if root.Right != nil {
		m[root.Right.Val] = root
		dfs(root.Right)
	}
}

func findDfs(K int, node *TreeNode, prev *TreeNode, dis int) {
	if node == nil {
		return
	}
	if dis == K {
		res = append(res, node.Val)
		return
	}
	if node.Left != prev { // 防止重复
		findDfs(K, node.Left, node, dis+1)
	}
	if node.Right != prev { // 防止重复
		findDfs(K, node.Right, node, dis+1)
	}
	if m[node.Val] != prev { // 防止重复：搜索父节点
		findDfs(K, m[node.Val], node, dis+1)
	}
}
