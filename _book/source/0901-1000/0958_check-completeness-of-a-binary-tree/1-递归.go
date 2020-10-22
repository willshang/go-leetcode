package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var m map[int]bool

func isCompleteTree(root *TreeNode) bool {
	m = make(map[int]bool)
	dfs(root, 1)
	for i := 1; i <= len(m); i++ {
		if m[i] == false {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode, id int) {
	if root == nil {
		return
	}
	m[id] = true
	dfs(root.Left, id*2)
	dfs(root.Right, id*2+1)
}
