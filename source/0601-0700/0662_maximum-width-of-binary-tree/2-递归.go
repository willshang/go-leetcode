package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int
var m map[int]int

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res = 0
	m = make(map[int]int)
	dfs(root, 0, 1)
	return res
}

func dfs(root *TreeNode, level int, id int) {
	if root == nil {
		return
	}
	if _, ok := m[level]; !ok {
		m[level] = id
	}
	if id-m[level]+1 > res {
		res = id - m[level] + 1
	}
	dfs(root.Left, level+1, id*2)
	dfs(root.Right, level+1, id*2+1)
}
